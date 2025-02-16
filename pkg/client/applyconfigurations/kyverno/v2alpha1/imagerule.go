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

// ImageRuleApplyConfiguration represents an declarative configuration of the ImageRule type for use
// with apply.
type ImageRuleApplyConfiguration struct {
	Glob          *string `json:"glob,omitempty"`
	CELExpression *string `json:"cel,omitempty"`
}

// ImageRuleApplyConfiguration constructs an declarative configuration of the ImageRule type for use with
// apply.
func ImageRule() *ImageRuleApplyConfiguration {
	return &ImageRuleApplyConfiguration{}
}

// WithGlob sets the Glob field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Glob field is set to the value of the last call.
func (b *ImageRuleApplyConfiguration) WithGlob(value string) *ImageRuleApplyConfiguration {
	b.Glob = &value
	return b
}

// WithCELExpression sets the CELExpression field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CELExpression field is set to the value of the last call.
func (b *ImageRuleApplyConfiguration) WithCELExpression(value string) *ImageRuleApplyConfiguration {
	b.CELExpression = &value
	return b
}
