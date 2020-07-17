/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/iam/v1alpha2"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// GlobalRolesGetter has a method to return a GlobalRoleInterface.
// A group's client should implement this interface.
type GlobalRolesGetter interface {
	GlobalRoles() GlobalRoleInterface
}

// GlobalRoleInterface has methods to work with GlobalRole resources.
type GlobalRoleInterface interface {
	Create(*v1alpha2.GlobalRole) (*v1alpha2.GlobalRole, error)
	Update(*v1alpha2.GlobalRole) (*v1alpha2.GlobalRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.GlobalRole, error)
	List(opts v1.ListOptions) (*v1alpha2.GlobalRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.GlobalRole, err error)
	GlobalRoleExpansion
}

// globalRoles implements GlobalRoleInterface
type globalRoles struct {
	client rest.Interface
}

// newGlobalRoles returns a GlobalRoles
func newGlobalRoles(c *IamV1alpha2Client) *globalRoles {
	return &globalRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalRole, and returns the corresponding globalRole object, and an error if there is any.
func (c *globalRoles) Get(name string, options v1.GetOptions) (result *v1alpha2.GlobalRole, err error) {
	result = &v1alpha2.GlobalRole{}
	err = c.client.Get().
		Resource("globalroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalRoles that match those selectors.
func (c *globalRoles) List(opts v1.ListOptions) (result *v1alpha2.GlobalRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.GlobalRoleList{}
	err = c.client.Get().
		Resource("globalroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalRoles.
func (c *globalRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a globalRole and creates it.  Returns the server's representation of the globalRole, and an error, if there is any.
func (c *globalRoles) Create(globalRole *v1alpha2.GlobalRole) (result *v1alpha2.GlobalRole, err error) {
	result = &v1alpha2.GlobalRole{}
	err = c.client.Post().
		Resource("globalroles").
		Body(globalRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a globalRole and updates it. Returns the server's representation of the globalRole, and an error, if there is any.
func (c *globalRoles) Update(globalRole *v1alpha2.GlobalRole) (result *v1alpha2.GlobalRole, err error) {
	result = &v1alpha2.GlobalRole{}
	err = c.client.Put().
		Resource("globalroles").
		Name(globalRole.Name).
		Body(globalRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the globalRole and deletes it. Returns an error if one occurs.
func (c *globalRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched globalRole.
func (c *globalRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.GlobalRole, err error) {
	result = &v1alpha2.GlobalRole{}
	err = c.client.Patch(pt).
		Resource("globalroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
