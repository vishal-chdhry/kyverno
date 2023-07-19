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

package v1

// ImageExtractorConfigApplyConfiguration represents an declarative configuration of the ImageExtractorConfig type for use
// with apply.
type ImageExtractorConfigApplyConfiguration struct {
	Path     *string `json:"path,omitempty"`
	Value    *string `json:"value,omitempty"`
	Name     *string `json:"name,omitempty"`
	Key      *string `json:"key,omitempty"`
	JMESPath *string `json:"jmesPath,omitempty"`
}

// ImageExtractorConfigApplyConfiguration constructs an declarative configuration of the ImageExtractorConfig type for use with
// apply.
func ImageExtractorConfig() *ImageExtractorConfigApplyConfiguration {
	return &ImageExtractorConfigApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ImageExtractorConfigApplyConfiguration) WithPath(value string) *ImageExtractorConfigApplyConfiguration {
	b.Path = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *ImageExtractorConfigApplyConfiguration) WithValue(value string) *ImageExtractorConfigApplyConfiguration {
	b.Value = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ImageExtractorConfigApplyConfiguration) WithName(value string) *ImageExtractorConfigApplyConfiguration {
	b.Name = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *ImageExtractorConfigApplyConfiguration) WithKey(value string) *ImageExtractorConfigApplyConfiguration {
	b.Key = &value
	return b
}

// WithJMESPath sets the JMESPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JMESPath field is set to the value of the last call.
func (b *ImageExtractorConfigApplyConfiguration) WithJMESPath(value string) *ImageExtractorConfigApplyConfiguration {
	b.JMESPath = &value
	return b
}