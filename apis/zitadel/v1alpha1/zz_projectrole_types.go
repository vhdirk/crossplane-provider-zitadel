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

type ProjectRoleInitParameters struct {

	// (String) Name used for project role
	// Name used for project role
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Group used for project role
	// Group used for project role
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

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

	// (String) ID of the project
	// ID of the project
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-zitadel/apis/zitadel/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (String) Key used for project role
	// Key used for project role
	RoleKey *string `json:"roleKey,omitempty" tf:"role_key,omitempty"`
}

type ProjectRoleObservation struct {

	// (String) Name used for project role
	// Name used for project role
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Group used for project role
	// Group used for project role
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) ID of the project
	// ID of the project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) Key used for project role
	// Key used for project role
	RoleKey *string `json:"roleKey,omitempty" tf:"role_key,omitempty"`
}

type ProjectRoleParameters struct {

	// (String) Name used for project role
	// Name used for project role
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Group used for project role
	// Group used for project role
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

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

	// (String) ID of the project
	// ID of the project
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-zitadel/apis/zitadel/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in zitadel to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (String) Key used for project role
	// Key used for project role
	// +kubebuilder:validation:Optional
	RoleKey *string `json:"roleKey,omitempty" tf:"role_key,omitempty"`
}

// ProjectRoleSpec defines the desired state of ProjectRole
type ProjectRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectRoleParameters `json:"forProvider"`
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
	InitProvider ProjectRoleInitParameters `json:"initProvider,omitempty"`
}

// ProjectRoleStatus defines the observed state of ProjectRole.
type ProjectRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectRole is the Schema for the ProjectRoles API. Resource representing the project roles, which can be given as authorizations to users.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type ProjectRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleKey) || (has(self.initProvider) && has(self.initProvider.roleKey))",message="spec.forProvider.roleKey is a required parameter"
	Spec   ProjectRoleSpec   `json:"spec"`
	Status ProjectRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectRoleList contains a list of ProjectRoles
type ProjectRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRole `json:"items"`
}

// Repository type metadata.
var (
	ProjectRole_Kind             = "ProjectRole"
	ProjectRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectRole_Kind}.String()
	ProjectRole_KindAPIVersion   = ProjectRole_Kind + "." + CRDGroupVersion.String()
	ProjectRole_GroupVersionKind = CRDGroupVersion.WithKind(ProjectRole_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectRole{}, &ProjectRoleList{})
}
