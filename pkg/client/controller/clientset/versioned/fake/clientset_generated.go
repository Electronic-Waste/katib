/*
Copyright 2022 The Kubeflow Authors.

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
	applyconfiguration "github.com/kubeflow/katib/pkg/client/controller/applyconfiguration"
	clientset "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned"
	experimentv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/experiments/v1beta1"
	fakeexperimentv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/experiments/v1beta1/fake"
	suggestionv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/suggestions/v1beta1"
	fakesuggestionv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/suggestions/v1beta1/fake"
	trialv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/trials/v1beta1"
	faketrialv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/trials/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any field management, validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
//
// DEPRECATED: NewClientset replaces this with support for field management, which significantly improves
// server side apply testing. NewClientset is only available when apply configurations are generated (e.g.
// via --with-applyconfig).
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
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
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

// NewClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewFieldManagedObjectTracker(
		scheme,
		codecs.UniversalDecoder(),
		applyconfiguration.NewTypeConverter(scheme),
	)
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
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

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ExperimentV1beta1 retrieves the ExperimentV1beta1Client
func (c *Clientset) ExperimentV1beta1() experimentv1beta1.ExperimentV1beta1Interface {
	return &fakeexperimentv1beta1.FakeExperimentV1beta1{Fake: &c.Fake}
}

// SuggestionV1beta1 retrieves the SuggestionV1beta1Client
func (c *Clientset) SuggestionV1beta1() suggestionv1beta1.SuggestionV1beta1Interface {
	return &fakesuggestionv1beta1.FakeSuggestionV1beta1{Fake: &c.Fake}
}

// TrialV1beta1 retrieves the TrialV1beta1Client
func (c *Clientset) TrialV1beta1() trialv1beta1.TrialV1beta1Interface {
	return &faketrialv1beta1.FakeTrialV1beta1{Fake: &c.Fake}
}
