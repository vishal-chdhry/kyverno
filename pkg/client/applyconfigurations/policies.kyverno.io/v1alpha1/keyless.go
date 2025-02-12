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

package v1alpha1

// KeylessApplyConfiguration represents an declarative configuration of the Keyless type for use
// with apply.
type KeylessApplyConfiguration struct {
	URL        *string                      `json:"url,omitempty"`
	Identities []IdentityApplyConfiguration `json:"identities,omitempty"`
	CACert     *KeyApplyConfiguration       `json:"ca-cert,omitempty"`
}

// KeylessApplyConfiguration constructs an declarative configuration of the Keyless type for use with
// apply.
func Keyless() *KeylessApplyConfiguration {
	return &KeylessApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *KeylessApplyConfiguration) WithURL(value string) *KeylessApplyConfiguration {
	b.URL = &value
	return b
}

// WithIdentities adds the given value to the Identities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Identities field.
func (b *KeylessApplyConfiguration) WithIdentities(values ...*IdentityApplyConfiguration) *KeylessApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIdentities")
		}
		b.Identities = append(b.Identities, *values[i])
	}
	return b
}

// WithCACert sets the CACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CACert field is set to the value of the last call.
func (b *KeylessApplyConfiguration) WithCACert(value *KeyApplyConfiguration) *KeylessApplyConfiguration {
	b.CACert = value
	return b
}
