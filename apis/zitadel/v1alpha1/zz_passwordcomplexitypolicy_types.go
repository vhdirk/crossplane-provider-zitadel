/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PasswordComplexityPolicyInitParameters struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-zitadel/apis/zitadel/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Organization in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`
}

type PasswordComplexityPolicyObservation struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type PasswordComplexityPolicyParameters struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	// +kubebuilder:validation:Optional
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	// +kubebuilder:validation:Optional
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	// +kubebuilder:validation:Optional
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	// +kubebuilder:validation:Optional
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	// +kubebuilder:validation:Optional
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-zitadel/apis/zitadel/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Organization in zitadel to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`
}

// PasswordComplexityPolicySpec defines the desired state of PasswordComplexityPolicy
type PasswordComplexityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PasswordComplexityPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PasswordComplexityPolicyInitParameters `json:"initProvider,omitempty"`
}

// PasswordComplexityPolicyStatus defines the observed state of PasswordComplexityPolicy.
type PasswordComplexityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PasswordComplexityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PasswordComplexityPolicy is the Schema for the PasswordComplexityPolicys API. Resource representing the custom password complexity policy of an organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type PasswordComplexityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasLowercase) || (has(self.initProvider) && has(self.initProvider.hasLowercase))",message="spec.forProvider.hasLowercase is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasNumber) || (has(self.initProvider) && has(self.initProvider.hasNumber))",message="spec.forProvider.hasNumber is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasSymbol) || (has(self.initProvider) && has(self.initProvider.hasSymbol))",message="spec.forProvider.hasSymbol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasUppercase) || (has(self.initProvider) && has(self.initProvider.hasUppercase))",message="spec.forProvider.hasUppercase is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minLength) || (has(self.initProvider) && has(self.initProvider.minLength))",message="spec.forProvider.minLength is a required parameter"
	Spec   PasswordComplexityPolicySpec   `json:"spec"`
	Status PasswordComplexityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordComplexityPolicyList contains a list of PasswordComplexityPolicys
type PasswordComplexityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PasswordComplexityPolicy `json:"items"`
}

// Repository type metadata.
var (
	PasswordComplexityPolicy_Kind             = "PasswordComplexityPolicy"
	PasswordComplexityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PasswordComplexityPolicy_Kind}.String()
	PasswordComplexityPolicy_KindAPIVersion   = PasswordComplexityPolicy_Kind + "." + CRDGroupVersion.String()
	PasswordComplexityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(PasswordComplexityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&PasswordComplexityPolicy{}, &PasswordComplexityPolicyList{})
}
