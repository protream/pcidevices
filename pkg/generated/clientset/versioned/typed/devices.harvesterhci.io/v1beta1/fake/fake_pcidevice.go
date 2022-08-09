/*
Copyright 2022 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/harvester/pcidevices/pkg/apis/devices.harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePCIDevices implements PCIDeviceInterface
type FakePCIDevices struct {
	Fake *FakeDevicesV1beta1
}

var pcidevicesResource = schema.GroupVersionResource{Group: "devices.harvesterhci.io", Version: "v1beta1", Resource: "pcidevices"}

var pcidevicesKind = schema.GroupVersionKind{Group: "devices.harvesterhci.io", Version: "v1beta1", Kind: "PCIDevice"}

// Get takes name of the pCIDevice, and returns the corresponding pCIDevice object, and an error if there is any.
func (c *FakePCIDevices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PCIDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pcidevicesResource, name), &v1beta1.PCIDevice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PCIDevice), err
}

// List takes label and field selectors, and returns the list of PCIDevices that match those selectors.
func (c *FakePCIDevices) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PCIDeviceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pcidevicesResource, pcidevicesKind, opts), &v1beta1.PCIDeviceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PCIDeviceList{ListMeta: obj.(*v1beta1.PCIDeviceList).ListMeta}
	for _, item := range obj.(*v1beta1.PCIDeviceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pCIDevices.
func (c *FakePCIDevices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pcidevicesResource, opts))
}

// Create takes the representation of a pCIDevice and creates it.  Returns the server's representation of the pCIDevice, and an error, if there is any.
func (c *FakePCIDevices) Create(ctx context.Context, pCIDevice *v1beta1.PCIDevice, opts v1.CreateOptions) (result *v1beta1.PCIDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pcidevicesResource, pCIDevice), &v1beta1.PCIDevice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PCIDevice), err
}

// Update takes the representation of a pCIDevice and updates it. Returns the server's representation of the pCIDevice, and an error, if there is any.
func (c *FakePCIDevices) Update(ctx context.Context, pCIDevice *v1beta1.PCIDevice, opts v1.UpdateOptions) (result *v1beta1.PCIDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pcidevicesResource, pCIDevice), &v1beta1.PCIDevice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PCIDevice), err
}

// Delete takes name of the pCIDevice and deletes it. Returns an error if one occurs.
func (c *FakePCIDevices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(pcidevicesResource, name, opts), &v1beta1.PCIDevice{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePCIDevices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pcidevicesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PCIDeviceList{})
	return err
}

// Patch applies the patch and returns the patched pCIDevice.
func (c *FakePCIDevices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PCIDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pcidevicesResource, name, pt, data, subresources...), &v1beta1.PCIDevice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PCIDevice), err
}