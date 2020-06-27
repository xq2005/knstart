/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "knstart/pkg/apis/operator/v1"
	scheme "knstart/pkg/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KNLearningsGetter has a method to return a KNLearningInterface.
// A group's client should implement this interface.
type KNLearningsGetter interface {
	KNLearnings(namespace string) KNLearningInterface
}

// KNLearningInterface has methods to work with KNLearning resources.
type KNLearningInterface interface {
	Create(*v1.KNLearning) (*v1.KNLearning, error)
	Update(*v1.KNLearning) (*v1.KNLearning, error)
	UpdateStatus(*v1.KNLearning) (*v1.KNLearning, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KNLearning, error)
	List(opts metav1.ListOptions) (*v1.KNLearningList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KNLearning, err error)
	KNLearningExpansion
}

// kNLearnings implements KNLearningInterface
type kNLearnings struct {
	client rest.Interface
	ns     string
}

// newKNLearnings returns a KNLearnings
func newKNLearnings(c *Xq2005V1Client, namespace string) *kNLearnings {
	return &kNLearnings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kNLearning, and returns the corresponding kNLearning object, and an error if there is any.
func (c *kNLearnings) Get(name string, options metav1.GetOptions) (result *v1.KNLearning, err error) {
	result = &v1.KNLearning{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knlearnings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KNLearnings that match those selectors.
func (c *kNLearnings) List(opts metav1.ListOptions) (result *v1.KNLearningList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KNLearningList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knlearnings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kNLearnings.
func (c *kNLearnings) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("knlearnings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kNLearning and creates it.  Returns the server's representation of the kNLearning, and an error, if there is any.
func (c *kNLearnings) Create(kNLearning *v1.KNLearning) (result *v1.KNLearning, err error) {
	result = &v1.KNLearning{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("knlearnings").
		Body(kNLearning).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kNLearning and updates it. Returns the server's representation of the kNLearning, and an error, if there is any.
func (c *kNLearnings) Update(kNLearning *v1.KNLearning) (result *v1.KNLearning, err error) {
	result = &v1.KNLearning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knlearnings").
		Name(kNLearning.Name).
		Body(kNLearning).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kNLearnings) UpdateStatus(kNLearning *v1.KNLearning) (result *v1.KNLearning, err error) {
	result = &v1.KNLearning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knlearnings").
		Name(kNLearning.Name).
		SubResource("status").
		Body(kNLearning).
		Do().
		Into(result)
	return
}

// Delete takes name of the kNLearning and deletes it. Returns an error if one occurs.
func (c *kNLearnings) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knlearnings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kNLearnings) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knlearnings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kNLearning.
func (c *kNLearnings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KNLearning, err error) {
	result = &v1.KNLearning{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("knlearnings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
