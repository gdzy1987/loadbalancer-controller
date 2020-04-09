/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	"time"

	scheme "github.com/caicloud/clientset/customclient/scheme"
	v1alpha3 "github.com/caicloud/clientset/pkg/apis/clever/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MLNeuronsGetter has a method to return a MLNeuronInterface.
// A group's client should implement this interface.
type MLNeuronsGetter interface {
	MLNeurons(namespace string) MLNeuronInterface
}

// MLNeuronInterface has methods to work with MLNeuron resources.
type MLNeuronInterface interface {
	Create(*v1alpha3.MLNeuron) (*v1alpha3.MLNeuron, error)
	Update(*v1alpha3.MLNeuron) (*v1alpha3.MLNeuron, error)
	UpdateStatus(*v1alpha3.MLNeuron) (*v1alpha3.MLNeuron, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha3.MLNeuron, error)
	List(opts v1.ListOptions) (*v1alpha3.MLNeuronList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.MLNeuron, err error)
	MLNeuronExpansion
}

// mLNeurons implements MLNeuronInterface
type mLNeurons struct {
	client rest.Interface
	ns     string
}

// newMLNeurons returns a MLNeurons
func newMLNeurons(c *CleverV1alpha3Client, namespace string) *mLNeurons {
	return &mLNeurons{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mLNeuron, and returns the corresponding mLNeuron object, and an error if there is any.
func (c *mLNeurons) Get(name string, options v1.GetOptions) (result *v1alpha3.MLNeuron, err error) {
	result = &v1alpha3.MLNeuron{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlneurons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MLNeurons that match those selectors.
func (c *mLNeurons) List(opts v1.ListOptions) (result *v1alpha3.MLNeuronList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.MLNeuronList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlneurons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mLNeurons.
func (c *mLNeurons) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mlneurons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mLNeuron and creates it.  Returns the server's representation of the mLNeuron, and an error, if there is any.
func (c *mLNeurons) Create(mLNeuron *v1alpha3.MLNeuron) (result *v1alpha3.MLNeuron, err error) {
	result = &v1alpha3.MLNeuron{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mlneurons").
		Body(mLNeuron).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mLNeuron and updates it. Returns the server's representation of the mLNeuron, and an error, if there is any.
func (c *mLNeurons) Update(mLNeuron *v1alpha3.MLNeuron) (result *v1alpha3.MLNeuron, err error) {
	result = &v1alpha3.MLNeuron{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mlneurons").
		Name(mLNeuron.Name).
		Body(mLNeuron).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mLNeurons) UpdateStatus(mLNeuron *v1alpha3.MLNeuron) (result *v1alpha3.MLNeuron, err error) {
	result = &v1alpha3.MLNeuron{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mlneurons").
		Name(mLNeuron.Name).
		SubResource("status").
		Body(mLNeuron).
		Do().
		Into(result)
	return
}

// Delete takes name of the mLNeuron and deletes it. Returns an error if one occurs.
func (c *mLNeurons) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlneurons").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mLNeurons) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlneurons").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mLNeuron.
func (c *mLNeurons) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.MLNeuron, err error) {
	result = &v1alpha3.MLNeuron{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mlneurons").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}