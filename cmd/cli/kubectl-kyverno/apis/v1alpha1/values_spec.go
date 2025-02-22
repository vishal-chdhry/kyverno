package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// ValuesSpec declares values to be loaded by the Kyverno CLI
type ValuesSpec struct {
	// GlobalValues are the global values
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	GlobalValues map[string]interface{} `json:"globalValues,omitempty"`

	// Policies are the policy values
	Policies []Policy `json:"policies,omitempty"`

	// NamespaceSelectors are the namespace labels
	NamespaceSelectors []NamespaceSelector `json:"namespaceSelector,omitempty"`

	// Namespaces are the namespaces
	Namespaces []corev1.Namespace `json:"namespaces,omitempty"`

	// Subresources are the subresource/parent resource mappings
	Subresources []Subresource `json:"subresources,omitempty"`
}
