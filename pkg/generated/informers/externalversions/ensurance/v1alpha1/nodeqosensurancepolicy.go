// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	ensurancev1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	versioned "github.com/gocrane/api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gocrane/api/pkg/generated/listers/ensurance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeQOSEnsurancePolicyInformer provides access to a shared informer and lister for
// NodeQOSEnsurancePolicies.
type NodeQOSEnsurancePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeQOSEnsurancePolicyLister
}

type nodeQOSEnsurancePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeQOSEnsurancePolicyInformer constructs a new informer for NodeQOSEnsurancePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeQOSEnsurancePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeQOSEnsurancePolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeQOSEnsurancePolicyInformer constructs a new informer for NodeQOSEnsurancePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeQOSEnsurancePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnsuranceV1alpha1().NodeQOSEnsurancePolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnsuranceV1alpha1().NodeQOSEnsurancePolicies().Watch(context.TODO(), options)
			},
		},
		&ensurancev1alpha1.NodeQOSEnsurancePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeQOSEnsurancePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeQOSEnsurancePolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeQOSEnsurancePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ensurancev1alpha1.NodeQOSEnsurancePolicy{}, f.defaultInformer)
}

func (f *nodeQOSEnsurancePolicyInformer) Lister() v1alpha1.NodeQOSEnsurancePolicyLister {
	return v1alpha1.NewNodeQOSEnsurancePolicyLister(f.Informer().GetIndexer())
}
