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

// SidecarDefinitionLister helps list SidecarDefinitions.
// All objects returned here must be treated as read-only.
type SidecarDefinitionLister interface {
	// List lists all SidecarDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SidecarDefinition, err error)
	// Get retrieves the SidecarDefinition from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SidecarDefinition, error)
	SidecarDefinitionListerExpansion
}

// sidecarDefinitionLister implements the SidecarDefinitionLister interface.
type sidecarDefinitionLister struct {
	indexer cache.Indexer
}

// NewSidecarDefinitionLister returns a new SidecarDefinitionLister.
func NewSidecarDefinitionLister(indexer cache.Indexer) SidecarDefinitionLister {
	return &sidecarDefinitionLister{indexer: indexer}
}

// List lists all SidecarDefinitions in the indexer.
func (s *sidecarDefinitionLister) List(selector labels.Selector) (ret []*v1.SidecarDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SidecarDefinition))
	})
	return ret, err
}

// Get retrieves the SidecarDefinition from the index for a given name.
func (s *sidecarDefinitionLister) Get(name string) (*v1.SidecarDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("sidecardefinition"), name)
	}
	return obj.(*v1.SidecarDefinition), nil
}
