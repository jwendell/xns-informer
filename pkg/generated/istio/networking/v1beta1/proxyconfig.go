/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	internalinterfaces "github.com/maistra/xns-informer/pkg/generated/istio/internalinterfaces"
	informers "github.com/maistra/xns-informer/pkg/informers"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	v1beta1 "istio.io/client-go/pkg/listers/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProxyConfigInformer provides access to a shared informer and lister for
// ProxyConfigs.
type ProxyConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ProxyConfigLister
}

type proxyConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespaces       informers.NamespaceSet
}

// NewProxyConfigInformer constructs a new informer for ProxyConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProxyConfigInformer(client versioned.Interface, namespaces informers.NamespaceSet, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProxyConfigInformer(client, namespaces, resyncPeriod, indexers, nil)
}

// NewFilteredProxyConfigInformer constructs a new informer for ProxyConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProxyConfigInformer(client versioned.Interface, namespaces informers.NamespaceSet, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	newInformer := func(namespace string) cache.SharedIndexInformer {
		return cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.NetworkingV1beta1().ProxyConfigs(namespace).List(context.TODO(), options)
				},
				WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.NetworkingV1beta1().ProxyConfigs(namespace).Watch(context.TODO(), options)
				},
			},
			&networkingv1beta1.ProxyConfig{},
			resyncPeriod,
			indexers,
		)
	}

	return informers.NewMultiNamespaceInformer(namespaces, resyncPeriod, newInformer)
}

func (f *proxyConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProxyConfigInformer(client, f.namespaces, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *proxyConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1beta1.ProxyConfig{}, f.defaultInformer)
}

func (f *proxyConfigInformer) Lister() v1beta1.ProxyConfigLister {
	return v1beta1.NewProxyConfigLister(f.Informer().GetIndexer())
}
