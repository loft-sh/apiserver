// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/test/apis/test/v1"
	"github.com/loft-sh/test/pkg/client/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type TestV1Interface interface {
	RESTClient() rest.Interface
	ClusterRolesGetter
}

// TestV1Client is used to interact with features provided by the test.loft.sh group.
type TestV1Client struct {
	restClient rest.Interface
}

func (c *TestV1Client) ClusterRoles() ClusterRoleInterface {
	return newClusterRoles(c)
}

// NewForConfig creates a new TestV1Client for the given config.
func NewForConfig(c *rest.Config) (*TestV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TestV1Client{client}, nil
}

// NewForConfigOrDie creates a new TestV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TestV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TestV1Client for the given RESTClient.
func New(c rest.Interface) *TestV1Client {
	return &TestV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TestV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
