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

package v1alpha1

import (
	v1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupRepoLister helps list BackupRepos.
// All objects returned here must be treated as read-only.
type BackupRepoLister interface {
	// List lists all BackupRepos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackupRepo, err error)
	// Get retrieves the BackupRepo from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BackupRepo, error)
	BackupRepoListerExpansion
}

// backupRepoLister implements the BackupRepoLister interface.
type backupRepoLister struct {
	indexer cache.Indexer
}

// NewBackupRepoLister returns a new BackupRepoLister.
func NewBackupRepoLister(indexer cache.Indexer) BackupRepoLister {
	return &backupRepoLister{indexer: indexer}
}

// List lists all BackupRepos in the indexer.
func (s *backupRepoLister) List(selector labels.Selector) (ret []*v1alpha1.BackupRepo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupRepo))
	})
	return ret, err
}

// Get retrieves the BackupRepo from the index for a given name.
func (s *backupRepoLister) Get(name string) (*v1alpha1.BackupRepo, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backuprepo"), name)
	}
	return obj.(*v1alpha1.BackupRepo), nil
}
