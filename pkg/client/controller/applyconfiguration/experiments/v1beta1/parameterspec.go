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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	experimentsv1beta1 "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1beta1"
)

// ParameterSpecApplyConfiguration represents a declarative configuration of the ParameterSpec type for use
// with apply.
type ParameterSpecApplyConfiguration struct {
	Name          *string                           `json:"name,omitempty"`
	ParameterType *experimentsv1beta1.ParameterType `json:"parameterType,omitempty"`
	FeasibleSpace *FeasibleSpaceApplyConfiguration  `json:"feasibleSpace,omitempty"`
}

// ParameterSpecApplyConfiguration constructs a declarative configuration of the ParameterSpec type for use with
// apply.
func ParameterSpec() *ParameterSpecApplyConfiguration {
	return &ParameterSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ParameterSpecApplyConfiguration) WithName(value string) *ParameterSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithParameterType sets the ParameterType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParameterType field is set to the value of the last call.
func (b *ParameterSpecApplyConfiguration) WithParameterType(value experimentsv1beta1.ParameterType) *ParameterSpecApplyConfiguration {
	b.ParameterType = &value
	return b
}

// WithFeasibleSpace sets the FeasibleSpace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FeasibleSpace field is set to the value of the last call.
func (b *ParameterSpecApplyConfiguration) WithFeasibleSpace(value *FeasibleSpaceApplyConfiguration) *ParameterSpecApplyConfiguration {
	b.FeasibleSpace = value
	return b
}
