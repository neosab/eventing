/*
Copyright 2018 The Knative Authors

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

	feeds_v1alpha1 "github.com/knative/eventing/pkg/apis/feeds/v1alpha1"
	versioned "github.com/knative/eventing/pkg/client/clientset/versioned"
	internalinterfaces "github.com/knative/eventing/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/knative/eventing/pkg/client/listers/feeds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterEventSourceInformer provides access to a shared informer and lister for
// ClusterEventSources.
type ClusterEventSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterEventSourceLister
}

type clusterEventSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterEventSourceInformer constructs a new informer for ClusterEventSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterEventSourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterEventSourceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterEventSourceInformer constructs a new informer for ClusterEventSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterEventSourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FeedsV1alpha1().ClusterEventSources().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FeedsV1alpha1().ClusterEventSources().Watch(options)
			},
		},
		&feeds_v1alpha1.ClusterEventSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterEventSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterEventSourceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterEventSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&feeds_v1alpha1.ClusterEventSource{}, f.defaultInformer)
}

func (f *clusterEventSourceInformer) Lister() v1alpha1.ClusterEventSourceLister {
	return v1alpha1.NewClusterEventSourceLister(f.Informer().GetIndexer())
}
