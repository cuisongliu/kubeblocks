/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

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

package v1

import (
	v1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComponentDefinitionLister helps list ComponentDefinitions.
// All objects returned here must be treated as read-only.
type ComponentDefinitionLister interface {
	// List lists all ComponentDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ComponentDefinition, err error)
	// Get retrieves the ComponentDefinition from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ComponentDefinition, error)
	ComponentDefinitionListerExpansion
}

// componentDefinitionLister implements the ComponentDefinitionLister interface.
type componentDefinitionLister struct {
	indexer cache.Indexer
}

// NewComponentDefinitionLister returns a new ComponentDefinitionLister.
func NewComponentDefinitionLister(indexer cache.Indexer) ComponentDefinitionLister {
	return &componentDefinitionLister{indexer: indexer}
}

// List lists all ComponentDefinitions in the indexer.
func (s *componentDefinitionLister) List(selector labels.Selector) (ret []*v1.ComponentDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ComponentDefinition))
	})
	return ret, err
}

// Get retrieves the ComponentDefinition from the index for a given name.
func (s *componentDefinitionLister) Get(name string) (*v1.ComponentDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("componentdefinition"), name)
	}
	return obj.(*v1.ComponentDefinition), nil
}
