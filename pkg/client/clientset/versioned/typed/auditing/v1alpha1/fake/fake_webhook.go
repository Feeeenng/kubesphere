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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/auditing/v1alpha1"
)

// FakeWebhooks implements WebhookInterface
type FakeWebhooks struct {
	Fake *FakeAuditingV1alpha1
}

var webhooksResource = schema.GroupVersionResource{Group: "auditing.kubesphere.io", Version: "v1alpha1", Resource: "webhooks"}

var webhooksKind = schema.GroupVersionKind{Group: "auditing.kubesphere.io", Version: "v1alpha1", Kind: "Webhook"}

// Get takes name of the webhook, and returns the corresponding webhook object, and an error if there is any.
func (c *FakeWebhooks) Get(name string, options v1.GetOptions) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(webhooksResource, name), &v1alpha1.Webhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// List takes label and field selectors, and returns the list of Webhooks that match those selectors.
func (c *FakeWebhooks) List(opts v1.ListOptions) (result *v1alpha1.WebhookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(webhooksResource, webhooksKind, opts), &v1alpha1.WebhookList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WebhookList{ListMeta: obj.(*v1alpha1.WebhookList).ListMeta}
	for _, item := range obj.(*v1alpha1.WebhookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested webhooks.
func (c *FakeWebhooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(webhooksResource, opts))
}

// Create takes the representation of a webhook and creates it.  Returns the server's representation of the webhook, and an error, if there is any.
func (c *FakeWebhooks) Create(webhook *v1alpha1.Webhook) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(webhooksResource, webhook), &v1alpha1.Webhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// Update takes the representation of a webhook and updates it. Returns the server's representation of the webhook, and an error, if there is any.
func (c *FakeWebhooks) Update(webhook *v1alpha1.Webhook) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(webhooksResource, webhook), &v1alpha1.Webhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// Delete takes name of the webhook and deletes it. Returns an error if one occurs.
func (c *FakeWebhooks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(webhooksResource, name), &v1alpha1.Webhook{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWebhooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(webhooksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WebhookList{})
	return err
}

// Patch applies the patch and returns the patched webhook.
func (c *FakeWebhooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(webhooksResource, name, pt, data, subresources...), &v1alpha1.Webhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}
