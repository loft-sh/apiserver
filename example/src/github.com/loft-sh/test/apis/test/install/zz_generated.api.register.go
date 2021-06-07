// Code generated by main. DO NOT EDIT.

package install

import (
	"github.com/loft-sh/apiserver/pkg/builders"
	"github.com/loft-sh/test/apis/test"
	v1 "github.com/loft-sh/test/apis/test/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func init() {
	Install(builders.Scheme)
}

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(test.AddToScheme(scheme))
	utilruntime.Must(addKnownTypes(scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(test.SchemeGroupVersion,
		&test.ClusterRole{},
		&test.ClusterRoleList{},
	)
	return nil
}
