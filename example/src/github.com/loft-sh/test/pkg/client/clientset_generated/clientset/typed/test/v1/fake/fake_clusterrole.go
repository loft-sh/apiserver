// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	testv1 "github.com/loft-sh/test/apis/test/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoles implements ClusterRoleInterface
type FakeClusterRoles struct {
	Fake *FakeTestV1
}

var clusterrolesResource = schema.GroupVersionResource{Group: "test.loft.sh", Version: "v1", Resource: "clusterroles"}

var clusterrolesKind = schema.GroupVersionKind{Group: "test.loft.sh", Version: "v1", Kind: "ClusterRole"}

// Get takes name of the clusterRole, and returns the corresponding clusterRole object, and an error if there is any.
func (c *FakeClusterRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *testv1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterrolesResource, name), &testv1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*testv1.ClusterRole), err
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors.
func (c *FakeClusterRoles) List(ctx context.Context, opts v1.ListOptions) (result *testv1.ClusterRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterrolesResource, clusterrolesKind, opts), &testv1.ClusterRoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &testv1.ClusterRoleList{ListMeta: obj.(*testv1.ClusterRoleList).ListMeta}
	for _, item := range obj.(*testv1.ClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoles.
func (c *FakeClusterRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterrolesResource, opts))
}

// Create takes the representation of a clusterRole and creates it.  Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *FakeClusterRoles) Create(ctx context.Context, clusterRole *testv1.ClusterRole, opts v1.CreateOptions) (result *testv1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterrolesResource, clusterRole), &testv1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*testv1.ClusterRole), err
}

// Update takes the representation of a clusterRole and updates it. Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *FakeClusterRoles) Update(ctx context.Context, clusterRole *testv1.ClusterRole, opts v1.UpdateOptions) (result *testv1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterrolesResource, clusterRole), &testv1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*testv1.ClusterRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRoles) UpdateStatus(ctx context.Context, clusterRole *testv1.ClusterRole, opts v1.UpdateOptions) (*testv1.ClusterRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterrolesResource, "status", clusterRole), &testv1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*testv1.ClusterRole), err
}

// Delete takes name of the clusterRole and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterrolesResource, name), &testv1.ClusterRole{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterrolesResource, listOpts)

	_, err := c.Fake.Invokes(action, &testv1.ClusterRoleList{})
	return err
}

// Patch applies the patch and returns the patched clusterRole.
func (c *FakeClusterRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *testv1.ClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrolesResource, name, pt, data, subresources...), &testv1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*testv1.ClusterRole), err
}