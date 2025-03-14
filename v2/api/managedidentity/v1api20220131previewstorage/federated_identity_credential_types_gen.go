// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220131previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=managedidentity.azure.com,resources=federatedidentitycredentials,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=managedidentity.azure.com,resources={federatedidentitycredentials/status,federatedidentitycredentials/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220131preview.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentities_FederatedIdentityCredential_Spec   `json:"spec,omitempty"`
	Status            UserAssignedIdentities_FederatedIdentityCredential_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FederatedIdentityCredential{}

// GetConditions returns the conditions of the resource
func (credential *FederatedIdentityCredential) GetConditions() conditions.Conditions {
	return credential.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (credential *FederatedIdentityCredential) SetConditions(conditions conditions.Conditions) {
	credential.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FederatedIdentityCredential{}

// AzureName returns the Azure name of the resource
func (credential *FederatedIdentityCredential) AzureName() string {
	return credential.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-31-preview"
func (credential FederatedIdentityCredential) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (credential *FederatedIdentityCredential) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (credential *FederatedIdentityCredential) GetSpec() genruntime.ConvertibleSpec {
	return &credential.Spec
}

// GetStatus returns the status of this resource
func (credential *FederatedIdentityCredential) GetStatus() genruntime.ConvertibleStatus {
	return &credential.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
func (credential *FederatedIdentityCredential) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}

// NewEmptyStatus returns a new empty (blank) status
func (credential *FederatedIdentityCredential) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (credential *FederatedIdentityCredential) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(credential.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  credential.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (credential *FederatedIdentityCredential) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*UserAssignedIdentities_FederatedIdentityCredential_STATUS); ok {
		credential.Status = *st
		return nil
	}

	// Convert status to required version
	var st UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	credential.Status = st
	return nil
}

// Hub marks that this FederatedIdentityCredential is the hub type for conversion
func (credential *FederatedIdentityCredential) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (credential *FederatedIdentityCredential) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: credential.Spec.OriginalVersion,
		Kind:    "FederatedIdentityCredential",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220131preview.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Storage version of v1api20220131preview.APIVersion
// +kubebuilder:validation:Enum={"2022-01-31-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-31-preview")

// Storage version of v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec
type UserAssignedIdentities_FederatedIdentityCredential_Spec struct {
	Audiences []string `json:"audiences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                         `json:"azureName,omitempty"`
	Issuer           *string                        `json:"issuer,omitempty" optionalConfigMapPair:"Issuer"`
	IssuerFromConfig *genruntime.ConfigMapReference `json:"issuerFromConfig,omitempty" optionalConfigMapPair:"Issuer"`
	OriginalVersion  string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a managedidentity.azure.com/UserAssignedIdentity resource
	Owner             *genruntime.KnownResourceReference `group:"managedidentity.azure.com" json:"owner,omitempty" kind:"UserAssignedIdentity"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subject           *string                            `json:"subject,omitempty" optionalConfigMapPair:"Subject"`
	SubjectFromConfig *genruntime.ConfigMapReference     `json:"subjectFromConfig,omitempty" optionalConfigMapPair:"Subject"`
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_FederatedIdentityCredential_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == credential {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(credential)
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == credential {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(credential)
}

// Storage version of v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS
type UserAssignedIdentities_FederatedIdentityCredential_STATUS struct {
	Audiences   []string               `json:"audiences,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Issuer      *string                `json:"issuer,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject     *string                `json:"subject,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}

// ConvertStatusFrom populates our UserAssignedIdentities_FederatedIdentityCredential_STATUS from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == credential {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(credential)
}

// ConvertStatusTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == credential {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(credential)
}

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
