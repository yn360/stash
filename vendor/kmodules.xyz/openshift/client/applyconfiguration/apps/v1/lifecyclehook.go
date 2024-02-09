/*
Copyright AppsCode Inc. and Contributors

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

package v1

import (
	v1 "kmodules.xyz/openshift/apis/apps/v1"
)

// LifecycleHookApplyConfiguration represents an declarative configuration of the LifecycleHook type for use
// with apply.
type LifecycleHookApplyConfiguration struct {
	FailurePolicy *v1.LifecycleHookFailurePolicy    `json:"failurePolicy,omitempty"`
	ExecNewPod    *ExecNewPodHookApplyConfiguration `json:"execNewPod,omitempty"`
	TagImages     []TagImageHookApplyConfiguration  `json:"tagImages,omitempty"`
}

// LifecycleHookApplyConfiguration constructs an declarative configuration of the LifecycleHook type for use with
// apply.
func LifecycleHook() *LifecycleHookApplyConfiguration {
	return &LifecycleHookApplyConfiguration{}
}

// WithFailurePolicy sets the FailurePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailurePolicy field is set to the value of the last call.
func (b *LifecycleHookApplyConfiguration) WithFailurePolicy(value v1.LifecycleHookFailurePolicy) *LifecycleHookApplyConfiguration {
	b.FailurePolicy = &value
	return b
}

// WithExecNewPod sets the ExecNewPod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExecNewPod field is set to the value of the last call.
func (b *LifecycleHookApplyConfiguration) WithExecNewPod(value *ExecNewPodHookApplyConfiguration) *LifecycleHookApplyConfiguration {
	b.ExecNewPod = value
	return b
}

// WithTagImages adds the given value to the TagImages field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TagImages field.
func (b *LifecycleHookApplyConfiguration) WithTagImages(values ...*TagImageHookApplyConfiguration) *LifecycleHookApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTagImages")
		}
		b.TagImages = append(b.TagImages, *values[i])
	}
	return b
}