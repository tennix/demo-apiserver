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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	corev1alpha1 "github.com/pingcap/apiserver/pkg/apis/core/v1alpha1"
	clientset "github.com/pingcap/apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/pingcap/apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/pingcap/apiserver/pkg/client/listers_generated/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TidbBackupInformer provides access to a shared informer and lister for
// TidbBackups.
type TidbBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TidbBackupLister
}

type tidbBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTidbBackupInformer constructs a new informer for TidbBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTidbBackupInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTidbBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTidbBackupInformer constructs a new informer for TidbBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTidbBackupInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().TidbBackups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().TidbBackups(namespace).Watch(options)
			},
		},
		&corev1alpha1.TidbBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *tidbBackupInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTidbBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tidbBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.TidbBackup{}, f.defaultInformer)
}

func (f *tidbBackupInformer) Lister() v1alpha1.TidbBackupLister {
	return v1alpha1.NewTidbBackupLister(f.Informer().GetIndexer())
}