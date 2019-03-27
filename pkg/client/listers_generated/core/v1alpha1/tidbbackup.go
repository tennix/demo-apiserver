/*
Copyright 2017 The Kubernetes Authors.

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
	v1alpha1 "github.com/pingcap/apiserver/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TidbBackupLister helps list TidbBackups.
type TidbBackupLister interface {
	// List lists all TidbBackups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TidbBackup, err error)
	// TidbBackups returns an object that can list and get TidbBackups.
	TidbBackups(namespace string) TidbBackupNamespaceLister
	TidbBackupListerExpansion
}

// tidbBackupLister implements the TidbBackupLister interface.
type tidbBackupLister struct {
	indexer cache.Indexer
}

// NewTidbBackupLister returns a new TidbBackupLister.
func NewTidbBackupLister(indexer cache.Indexer) TidbBackupLister {
	return &tidbBackupLister{indexer: indexer}
}

// List lists all TidbBackups in the indexer.
func (s *tidbBackupLister) List(selector labels.Selector) (ret []*v1alpha1.TidbBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbBackup))
	})
	return ret, err
}

// TidbBackups returns an object that can list and get TidbBackups.
func (s *tidbBackupLister) TidbBackups(namespace string) TidbBackupNamespaceLister {
	return tidbBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TidbBackupNamespaceLister helps list and get TidbBackups.
type TidbBackupNamespaceLister interface {
	// List lists all TidbBackups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TidbBackup, err error)
	// Get retrieves the TidbBackup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TidbBackup, error)
	TidbBackupNamespaceListerExpansion
}

// tidbBackupNamespaceLister implements the TidbBackupNamespaceLister
// interface.
type tidbBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TidbBackups in the indexer for a given namespace.
func (s tidbBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TidbBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbBackup))
	})
	return ret, err
}

// Get retrieves the TidbBackup from the indexer for a given namespace and name.
func (s tidbBackupNamespaceLister) Get(name string) (*v1alpha1.TidbBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tidbbackup"), name)
	}
	return obj.(*v1alpha1.TidbBackup), nil
}