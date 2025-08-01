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

type OrganizationInitParameters struct {

	// (Boolean) True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) Name of the org
	// Name of the org
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Primary domain of the org
	// Primary domain of the org
	PrimaryDomain *string `json:"primaryDomain,omitempty" tf:"primary_domain,omitempty"`
}

type OrganizationObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) Name of the org
	// Name of the org
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Primary domain of the org
	// Primary domain of the org
	PrimaryDomain *string `json:"primaryDomain,omitempty" tf:"primary_domain,omitempty"`

	// (String) State of the org
	// State of the org
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type OrganizationParameters struct {

	// (Boolean) True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	// True sets the org as default org for the instance. Only one org can be default org. Nothing happens if you set it to false until you set another org as default org.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) Name of the org
	// Name of the org
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Primary domain of the org
	// Primary domain of the org
	// +kubebuilder:validation:Optional
	PrimaryDomain *string `json:"primaryDomain,omitempty" tf:"primary_domain,omitempty"`
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
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
	InitProvider OrganizationInitParameters `json:"initProvider,omitempty"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Organization is the Schema for the Organizations API. Resource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OrganizationSpec   `json:"spec"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
