/*
Copyright The Kubernetes Authors.

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
	"context"
	"time"

	v1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S3sGetter has a method to return a S3Interface.
// A group's client should implement this interface.
type S3sGetter interface {
	S3s(namespace string) S3Interface
}

// S3Interface has methods to work with S3 resources.
type S3Interface interface {
	Create(ctx context.Context, s3 *v1.S3, opts metav1.CreateOptions) (*v1.S3, error)
	Update(ctx context.Context, s3 *v1.S3, opts metav1.UpdateOptions) (*v1.S3, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.S3, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.S3List, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.S3, err error)
	S3Expansion
}

// s3s implements S3Interface
type s3s struct {
	client rest.Interface
	ns     string
}

// newS3s returns a S3s
func newS3s(c *EdgefsV1Client, namespace string) *s3s {
	return &s3s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3, and returns the corresponding s3 object, and an error if there is any.
func (c *s3s) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.S3, err error) {
	result = &v1.S3{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3s that match those selectors.
func (c *s3s) List(ctx context.Context, opts metav1.ListOptions) (result *v1.S3List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.S3List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3s.
func (c *s3s) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a s3 and creates it.  Returns the server's representation of the s3, and an error, if there is any.
func (c *s3s) Create(ctx context.Context, s3 *v1.S3, opts metav1.CreateOptions) (result *v1.S3, err error) {
	result = &v1.S3{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a s3 and updates it. Returns the server's representation of the s3, and an error, if there is any.
func (c *s3s) Update(ctx context.Context, s3 *v1.S3, opts metav1.UpdateOptions) (result *v1.S3, err error) {
	result = &v1.S3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3s").
		Name(s3.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the s3 and deletes it. Returns an error if one occurs.
func (c *s3s) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3s) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched s3.
func (c *s3s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.S3, err error) {
	result = &v1.S3{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
