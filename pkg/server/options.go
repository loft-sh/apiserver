package apiserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/loft-sh/apiserver/pkg/admission"
	"github.com/loft-sh/apiserver/pkg/apiserver"
	"github.com/loft-sh/apiserver/pkg/builders"
	oteltrace "go.opentelemetry.io/otel/trace"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericfilters "k8s.io/apiserver/pkg/server/filters"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/apiserver/pkg/util/webhook"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	versionutil "k8s.io/component-base/version"
	"k8s.io/klog/v2"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	openapi "k8s.io/kube-openapi/pkg/common"
	ctrl "sigs.k8s.io/controller-runtime"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
	APIBuilders        []*builders.APIGroupBuilder

	GetOpenAPIDefinitions openapi.GetOpenAPIDefinitions
	DisableWebhooks       bool
}

func (o *ServerOptions) GenericConfig(tweakConfig func(config *genericapiserver.RecommendedConfig) error) (*genericapiserver.RecommendedConfig, error) {
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, nil); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(builders.Codecs)
	loopbackKubeConfig, kubeInformerFactory, err := o.buildLoopback()
	if err != nil {
		klog.Warningf("attempting to instantiate loopback client but failed: %v", err)
	} else {
		serverConfig.LoopbackClientConfig = loopbackKubeConfig
		serverConfig.SharedInformerFactory = kubeInformerFactory
	}
	kubeClient, err := kubernetes.NewForConfig(serverConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}

	_ = o.RecommendedOptions.Authorization.ApplyTo(&serverConfig.Authorization)

	// admission webhooks
	if !o.DisableWebhooks && serverConfig.LoopbackClientConfig != nil {
		proxyTransport := createNodeDialer()
		admissionConfig := &admission.Config{
			ExternalInformers:    kubeInformerFactory,
			LoopbackClientConfig: serverConfig.LoopbackClientConfig,
		}

		serviceResolver := buildServiceResolver(serverConfig.LoopbackClientConfig.Host, kubeInformerFactory)
		tp := oteltrace.NewNoopTracerProvider()
		pluginInitializers, admissionPostStartHook, err := admissionConfig.New(proxyTransport, serverConfig.EgressSelector, serviceResolver, &tp)
		if err != nil {
			return nil, fmt.Errorf("failed to create admission plugin initializer: %v", err)
		}
		if err := serverConfig.AddPostStartHook("start-kube-apiserver-admission-initializer", admissionPostStartHook); err != nil {
			return nil, err
		}

		dynamicClient, err := dynamic.NewForConfig(serverConfig.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}

		err = o.RecommendedOptions.Admission.ApplyTo(
			&serverConfig.Config,
			kubeInformerFactory,
			kubeClient,
			dynamicClient,
			utilfeature.DefaultFeatureGate,
			pluginInitializers...)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize admission: %v", err)
		}
	}

	err = applyOptions(
		&serverConfig.Config,
		// o.RecommendedOptions.Etcd.ApplyTo,
		func(cfg *genericapiserver.Config) error {
			return o.RecommendedOptions.SecureServing.ApplyTo(&cfg.SecureServing, &cfg.LoopbackClientConfig)
		},
		func(cfg *genericapiserver.Config) error {
			return o.RecommendedOptions.Audit.ApplyTo(
				&serverConfig.Config,
			)
		},
		func(c *genericapiserver.Config) error {
			return o.RecommendedOptions.Features.ApplyTo(c, kubeClient, kubeInformerFactory)
		},
	)
	if err != nil {
		return nil, err
	}

	_ = o.RecommendedOptions.Authentication.ApplyTo(&serverConfig.Authentication, serverConfig.Config.SecureServing, serverConfig.Config.OpenAPIConfig)
	if tweakConfig != nil {
		if err := tweakConfig(serverConfig); err != nil {
			return nil, err
		}
	}

	return serverConfig, nil
}

func (o *ServerOptions) RunServer(APIServerVersion *version.Info, stopCh <-chan struct{}, authorizer authorizer.Authorizer, tweakServerConfig func(config *genericapiserver.RecommendedConfig) error) error {
	aggregatedAPIServerConfig, err := o.GenericConfig(tweakServerConfig)
	if err != nil {
		return err
	}

	// aggregatedAPIServerConfig.EffectiveVersion = utilapiserverversion.DefaultComponentGlobalsRegistry.EffectiveVersionFor("loft-apiserver")
	aggregatedAPIServerConfig.EffectiveVersion = versionutil.NewEffectiveVersion(APIServerVersion.String())

	// set the basics
	genericConfig := &aggregatedAPIServerConfig.Config
	genericConfig.EffectiveVersion = aggregatedAPIServerConfig.EffectiveVersion
	genericConfig.Authorization.Authorizer = authorizer

	// set open api
	genericConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(o.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(builders.Scheme))
	genericConfig.OpenAPIConfig.Info.Title = "Api"
	genericConfig.OpenAPIConfig.Info.Version = "v0"
	if genericConfig.LongRunningFunc == nil {
		genericConfig.LongRunningFunc = genericfilters.BasicLongRunningRequestCheck(
			sets.NewString("watch", "proxy"),
			sets.NewString("attach", "exec", "proxy", "log", "portforward"),
		)
	}

	// create a new server
	genericServer, err := apiserver.NewServer(aggregatedAPIServerConfig, o.APIBuilders)
	if err != nil {
		return err
	}

	s := genericServer.GenericAPIServer.PrepareRun()
	return s.Run(stopCh)
}

func (o *ServerOptions) buildLoopback() (*rest.Config, informers.SharedInformerFactory, error) {
	var loopbackConfig *rest.Config
	var err error

	loopbackConfig, err = ctrl.GetConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get config: %w", err)
	}

	loopbackClient, err := kubernetes.NewForConfig(loopbackConfig)
	if err != nil {
		return nil, nil, err
	}
	kubeInformerFactory := informers.NewSharedInformerFactory(loopbackClient, 0)
	return loopbackConfig, kubeInformerFactory, nil
}

func createNodeDialer() *http.Transport {
	// Setup nodeTunneler if needed
	var proxyDialerFn utilnet.DialFunc

	// Proxying to pods and services is IP-based... don't expect to be able to verify the hostname
	proxyTLSClientConfig := &tls.Config{InsecureSkipVerify: true}
	proxyTransport := utilnet.SetTransportDefaults(&http.Transport{
		DialContext:     proxyDialerFn,
		TLSClientConfig: proxyTLSClientConfig,
	})
	return proxyTransport
}

func buildServiceResolver(hostname string, informer informers.SharedInformerFactory) webhook.ServiceResolver {
	var serviceResolver webhook.ServiceResolver
	serviceResolver = aggregatorapiserver.NewClusterIPServiceResolver(
		informer.Core().V1().Services().Lister(),
	)

	// resolve kubernetes.default.svc locally
	if localHost, err := url.Parse(hostname); err == nil {
		serviceResolver = aggregatorapiserver.NewLoopbackServiceResolver(serviceResolver, localHost)
	}
	return serviceResolver
}
