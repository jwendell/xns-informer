// Code generated by xns-informer-gen. DO NOT EDIT.
package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/core/v1"
)

type Interface interface {
	ConfigMaps() informers.ConfigMapInformer
	Endpoints() informers.EndpointsInformer
	Namespaces() informers.NamespaceInformer
	Pods() informers.PodInformer
	Services() informers.ServiceInformer
}

type version struct {
	factory xnsinformers.InformerFactory
}

func New(factory xnsinformers.InformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) ConfigMaps() informers.ConfigMapInformer {
	return &configMapInformer{factory: v.factory}
}
func (v *version) Endpoints() informers.EndpointsInformer {
	return &endpointsInformer{factory: v.factory}
}
func (v *version) Namespaces() informers.NamespaceInformer {
	return &namespaceInformer{factory: v.factory}
}
func (v *version) Pods() informers.PodInformer {
	return &podInformer{factory: v.factory}
}
func (v *version) Services() informers.ServiceInformer {
	return &serviceInformer{factory: v.factory}
}
