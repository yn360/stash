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

package v1alpha1

// ParameterApplyConfiguration represents an declarative configuration of the Parameter type for use
// with apply.
type ParameterApplyConfiguration struct {
	Key       *string `json:"key,omitempty"`
	Value     *string `json:"value,omitempty"`
	ValuePath *string `json:"valuePath,omitempty"`
}

// ParameterApplyConfiguration constructs an declarative configuration of the Parameter type for use with
// apply.
func Parameter() *ParameterApplyConfiguration {
	return &ParameterApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *ParameterApplyConfiguration) WithKey(value string) *ParameterApplyConfiguration {
	b.Key = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *ParameterApplyConfiguration) WithValue(value string) *ParameterApplyConfiguration {
	b.Value = &value
	return b
}

// WithValuePath sets the ValuePath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ValuePath field is set to the value of the last call.
func (b *ParameterApplyConfiguration) WithValuePath(value string) *ParameterApplyConfiguration {
	b.ValuePath = &value
	return b
}