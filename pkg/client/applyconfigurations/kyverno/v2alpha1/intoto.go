/*
Copyright The Kubernetes Authors.

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

package v2alpha1

// InTotoApplyConfiguration represents an declarative configuration of the InToto type for use
// with apply.
type InTotoApplyConfiguration struct {
	Type *string `json:"type,omitempty"`
}

// InTotoApplyConfiguration constructs an declarative configuration of the InToto type for use with
// apply.
func InToto() *InTotoApplyConfiguration {
	return &InTotoApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *InTotoApplyConfiguration) WithType(value string) *InTotoApplyConfiguration {
	b.Type = &value
	return b
}
