// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta2

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta2 "k8s.io/api/apps/v1beta2"
	informers "k8s.io/client-go/informers/apps/v1beta2"
	listers "k8s.io/client-go/listers/apps/v1beta2"
	"k8s.io/client-go/tools/cache"
)

type replicaSetInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ReplicaSetInformer = &replicaSetInformer{}

func NewReplicaSetInformer(f xnsinformers.SharedInformerFactory) informers.ReplicaSetInformer {
	resource := v1beta2.SchemeGroupVersion.WithResource("replicasets")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta2.ReplicaSet{},
		&v1beta2.ReplicaSetList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &replicaSetInformer{informer: informer.Informer()}
}

func (i *replicaSetInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *replicaSetInformer) Lister() listers.ReplicaSetLister {
	return listers.NewReplicaSetLister(i.informer.GetIndexer())
}
