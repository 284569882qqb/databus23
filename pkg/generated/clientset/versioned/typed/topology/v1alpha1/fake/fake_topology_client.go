// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/topology/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTopologyV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTopologyV1alpha1) NodeResourceTopologies() v1alpha1.NodeResourceTopologyInterface {
	return &FakeNodeResourceTopologies{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTopologyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
