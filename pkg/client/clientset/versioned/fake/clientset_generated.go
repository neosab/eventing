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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/knative/eventing/pkg/client/clientset/versioned"
	channelsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/channels/v1alpha1"
	fakechannelsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/channels/v1alpha1/fake"
	feedsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/feeds/v1alpha1"
	fakefeedsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/feeds/v1alpha1/fake"
	flowsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/flows/v1alpha1"
	fakeflowsv1alpha1 "github.com/knative/eventing/pkg/client/clientset/versioned/typed/flows/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ChannelsV1alpha1 retrieves the ChannelsV1alpha1Client
func (c *Clientset) ChannelsV1alpha1() channelsv1alpha1.ChannelsV1alpha1Interface {
	return &fakechannelsv1alpha1.FakeChannelsV1alpha1{Fake: &c.Fake}
}

// Channels retrieves the ChannelsV1alpha1Client
func (c *Clientset) Channels() channelsv1alpha1.ChannelsV1alpha1Interface {
	return &fakechannelsv1alpha1.FakeChannelsV1alpha1{Fake: &c.Fake}
}

// FeedsV1alpha1 retrieves the FeedsV1alpha1Client
func (c *Clientset) FeedsV1alpha1() feedsv1alpha1.FeedsV1alpha1Interface {
	return &fakefeedsv1alpha1.FakeFeedsV1alpha1{Fake: &c.Fake}
}

// Feeds retrieves the FeedsV1alpha1Client
func (c *Clientset) Feeds() feedsv1alpha1.FeedsV1alpha1Interface {
	return &fakefeedsv1alpha1.FakeFeedsV1alpha1{Fake: &c.Fake}
}

// FlowsV1alpha1 retrieves the FlowsV1alpha1Client
func (c *Clientset) FlowsV1alpha1() flowsv1alpha1.FlowsV1alpha1Interface {
	return &fakeflowsv1alpha1.FakeFlowsV1alpha1{Fake: &c.Fake}
}

// Flows retrieves the FlowsV1alpha1Client
func (c *Clientset) Flows() flowsv1alpha1.FlowsV1alpha1Interface {
	return &fakeflowsv1alpha1.FakeFlowsV1alpha1{Fake: &c.Fake}
}
