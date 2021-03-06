// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePacketMachineClasses implements PacketMachineClassInterface
type FakePacketMachineClasses struct {
	Fake *FakeMachineV1alpha1
	ns   string
}

var packetmachineclassesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "v1alpha1", Resource: "packetmachineclasses"}

var packetmachineclassesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "v1alpha1", Kind: "PacketMachineClass"}

// Get takes name of the packetMachineClass, and returns the corresponding packetMachineClass object, and an error if there is any.
func (c *FakePacketMachineClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.PacketMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packetmachineclassesResource, c.ns, name), &v1alpha1.PacketMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketMachineClass), err
}

// List takes label and field selectors, and returns the list of PacketMachineClasses that match those selectors.
func (c *FakePacketMachineClasses) List(opts v1.ListOptions) (result *v1alpha1.PacketMachineClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packetmachineclassesResource, packetmachineclassesKind, c.ns, opts), &v1alpha1.PacketMachineClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PacketMachineClassList{ListMeta: obj.(*v1alpha1.PacketMachineClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.PacketMachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packetMachineClasses.
func (c *FakePacketMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packetmachineclassesResource, c.ns, opts))

}

// Create takes the representation of a packetMachineClass and creates it.  Returns the server's representation of the packetMachineClass, and an error, if there is any.
func (c *FakePacketMachineClasses) Create(packetMachineClass *v1alpha1.PacketMachineClass) (result *v1alpha1.PacketMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packetmachineclassesResource, c.ns, packetMachineClass), &v1alpha1.PacketMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketMachineClass), err
}

// Update takes the representation of a packetMachineClass and updates it. Returns the server's representation of the packetMachineClass, and an error, if there is any.
func (c *FakePacketMachineClasses) Update(packetMachineClass *v1alpha1.PacketMachineClass) (result *v1alpha1.PacketMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packetmachineclassesResource, c.ns, packetMachineClass), &v1alpha1.PacketMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketMachineClass), err
}

// Delete takes name of the packetMachineClass and deletes it. Returns an error if one occurs.
func (c *FakePacketMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(packetmachineclassesResource, c.ns, name), &v1alpha1.PacketMachineClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePacketMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packetmachineclassesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PacketMachineClassList{})
	return err
}

// Patch applies the patch and returns the patched packetMachineClass.
func (c *FakePacketMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packetmachineclassesResource, c.ns, name, data, subresources...), &v1alpha1.PacketMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketMachineClass), err
}
