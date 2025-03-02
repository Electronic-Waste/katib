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
	v1beta1 "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1beta1"
	experimentsv1beta1 "github.com/kubeflow/katib/pkg/client/controller/applyconfiguration/experiments/v1beta1"
	typedexperimentsv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/experiments/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeExperiments implements ExperimentInterface
type fakeExperiments struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.Experiment, *v1beta1.ExperimentList, *experimentsv1beta1.ExperimentApplyConfiguration]
	Fake *FakeExperimentV1beta1
}

func newFakeExperiments(fake *FakeExperimentV1beta1, namespace string) typedexperimentsv1beta1.ExperimentInterface {
	return &fakeExperiments{
		gentype.NewFakeClientWithListAndApply[*v1beta1.Experiment, *v1beta1.ExperimentList, *experimentsv1beta1.ExperimentApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("experiments"),
			v1beta1.SchemeGroupVersion.WithKind("Experiment"),
			func() *v1beta1.Experiment { return &v1beta1.Experiment{} },
			func() *v1beta1.ExperimentList { return &v1beta1.ExperimentList{} },
			func(dst, src *v1beta1.ExperimentList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.ExperimentList) []*v1beta1.Experiment { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.ExperimentList, items []*v1beta1.Experiment) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
