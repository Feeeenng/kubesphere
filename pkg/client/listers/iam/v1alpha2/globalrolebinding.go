/*
Copyright 2019 The KubeSphere Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/iam/v1alpha2"
)

// GlobalRoleBindingLister helps list GlobalRoleBindings.
type GlobalRoleBindingLister interface {
	// List lists all GlobalRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.GlobalRoleBinding, err error)
	// Get retrieves the GlobalRoleBinding from the index for a given name.
	Get(name string) (*v1alpha2.GlobalRoleBinding, error)
	GlobalRoleBindingListerExpansion
}

// globalRoleBindingLister implements the GlobalRoleBindingLister interface.
type globalRoleBindingLister struct {
	indexer cache.Indexer
}

// NewGlobalRoleBindingLister returns a new GlobalRoleBindingLister.
func NewGlobalRoleBindingLister(indexer cache.Indexer) GlobalRoleBindingLister {
	return &globalRoleBindingLister{indexer: indexer}
}

// List lists all GlobalRoleBindings in the indexer.
func (s *globalRoleBindingLister) List(selector labels.Selector) (ret []*v1alpha2.GlobalRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.GlobalRoleBinding))
	})
	return ret, err
}

// Get retrieves the GlobalRoleBinding from the index for a given name.
func (s *globalRoleBindingLister) Get(name string) (*v1alpha2.GlobalRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("globalrolebinding"), name)
	}
	return obj.(*v1alpha2.GlobalRoleBinding), nil
}
