// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Analyticses returns a AnalyticsInformer.
	Analyticses() AnalyticsInformer
	// Recommendations returns a RecommendationInformer.
	Recommendations() RecommendationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Analyticses returns a AnalyticsInformer.
func (v *version) Analyticses() AnalyticsInformer {
	return &analyticsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Recommendations returns a RecommendationInformer.
func (v *version) Recommendations() RecommendationInformer {
	return &recommendationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
