// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=batch.azure.com,resources=batchaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.azure.com,resources={batchaccounts/status,batchaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210101.BatchAccount
// Generator information:
// - Generated from: /batch/resource-manager/Microsoft.Batch/stable/2021-01-01/BatchManagement.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}
type BatchAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccount_Spec   `json:"spec,omitempty"`
	Status            BatchAccount_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &BatchAccount{}

// GetConditions returns the conditions of the resource
func (account *BatchAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *BatchAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &BatchAccount{}

// AzureName returns the Azure name of the resource
func (account *BatchAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (account BatchAccount) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (account *BatchAccount) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (account *BatchAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *BatchAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Batch/batchAccounts"
func (account *BatchAccount) GetType() string {
	return "Microsoft.Batch/batchAccounts"
}

// NewEmptyStatus returns a new empty (blank) status
func (account *BatchAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BatchAccount_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (account *BatchAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *BatchAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BatchAccount_STATUS); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st BatchAccount_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this BatchAccount is the hub type for conversion
func (account *BatchAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *BatchAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "BatchAccount",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210101.BatchAccount
// Generator information:
// - Generated from: /batch/resource-manager/Microsoft.Batch/stable/2021-01-01/BatchManagement.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}
type BatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchAccount `json:"items"`
}

// Storage version of v1api20210101.APIVersion
// +kubebuilder:validation:Enum={"2021-01-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-01-01")

// Storage version of v1api20210101.BatchAccount_Spec
type BatchAccount_Spec struct {
	AutoStorage *AutoStorageBaseProperties `json:"autoStorage,omitempty"`

	// +kubebuilder:validation:MaxLength=24
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern="^[a-z0-9]+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string                `json:"azureName,omitempty"`
	Encryption        *EncryptionProperties `json:"encryption,omitempty"`
	Identity          *BatchAccountIdentity `json:"identity,omitempty"`
	KeyVaultReference *KeyVaultReference    `json:"keyVaultReference,omitempty"`
	Location          *string               `json:"location,omitempty"`
	OriginalVersion   string                `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PoolAllocationMode  *string                            `json:"poolAllocationMode,omitempty"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &BatchAccount_Spec{}

// ConvertSpecFrom populates our BatchAccount_Spec from the provided source
func (account *BatchAccount_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(account)
}

// ConvertSpecTo populates the provided destination from our BatchAccount_Spec
func (account *BatchAccount_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(account)
}

// Storage version of v1api20210101.BatchAccount_STATUS
// Contains information about an Azure Batch account.
type BatchAccount_STATUS struct {
	AccountEndpoint                       *string                                `json:"accountEndpoint,omitempty"`
	ActiveJobAndJobScheduleQuota          *int                                   `json:"activeJobAndJobScheduleQuota,omitempty"`
	AutoStorage                           *AutoStorageProperties_STATUS          `json:"autoStorage,omitempty"`
	Conditions                            []conditions.Condition                 `json:"conditions,omitempty"`
	DedicatedCoreQuota                    *int                                   `json:"dedicatedCoreQuota,omitempty"`
	DedicatedCoreQuotaPerVMFamily         []VirtualMachineFamilyCoreQuota_STATUS `json:"dedicatedCoreQuotaPerVMFamily,omitempty"`
	DedicatedCoreQuotaPerVMFamilyEnforced *bool                                  `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`
	Encryption                            *EncryptionProperties_STATUS           `json:"encryption,omitempty"`
	Id                                    *string                                `json:"id,omitempty"`
	Identity                              *BatchAccountIdentity_STATUS           `json:"identity,omitempty"`
	KeyVaultReference                     *KeyVaultReference_STATUS              `json:"keyVaultReference,omitempty"`
	Location                              *string                                `json:"location,omitempty"`
	LowPriorityCoreQuota                  *int                                   `json:"lowPriorityCoreQuota,omitempty"`
	Name                                  *string                                `json:"name,omitempty"`
	PoolAllocationMode                    *string                                `json:"poolAllocationMode,omitempty"`
	PoolQuota                             *int                                   `json:"poolQuota,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_STATUS     `json:"privateEndpointConnections,omitempty"`
	PropertyBag                           genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ProvisioningState                     *string                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess                   *string                                `json:"publicNetworkAccess,omitempty"`
	Tags                                  map[string]string                      `json:"tags,omitempty"`
	Type                                  *string                                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BatchAccount_STATUS{}

// ConvertStatusFrom populates our BatchAccount_STATUS from the provided source
func (account *BatchAccount_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our BatchAccount_STATUS
func (account *BatchAccount_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

// Storage version of v1api20210101.AutoStorageBaseProperties
// The properties related to the auto-storage account.
type AutoStorageBaseProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// StorageAccountReference: The resource ID of the storage account to be used for auto-storage account.
	StorageAccountReference *genruntime.ResourceReference `armReference:"StorageAccountId" json:"storageAccountReference,omitempty"`
}

// Storage version of v1api20210101.AutoStorageProperties_STATUS
// Contains information about the auto-storage account associated with a Batch account.
type AutoStorageProperties_STATUS struct {
	LastKeySync      *string                `json:"lastKeySync,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountId *string                `json:"storageAccountId,omitempty"`
}

// Storage version of v1api20210101.BatchAccountIdentity
// The identity of the Batch account, if configured. This is only used when the user specifies 'Microsoft.KeyVault' as
// their Batch account encryption configuration.
type BatchAccountIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210101.BatchAccountIdentity_STATUS
// The identity of the Batch account, if configured. This is only used when the user specifies 'Microsoft.KeyVault' as
// their Batch account encryption configuration.
type BatchAccountIdentity_STATUS struct {
	PrincipalId            *string                                                       `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	TenantId               *string                                                       `json:"tenantId,omitempty"`
	Type                   *string                                                       `json:"type,omitempty"`
	UserAssignedIdentities map[string]BatchAccountIdentity_UserAssignedIdentities_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210101.EncryptionProperties
// Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft
// managed key. For additional control, a customer-managed key can be used instead.
type EncryptionProperties struct {
	KeySource          *string                `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties    `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.EncryptionProperties_STATUS
// Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft
// managed key. For additional control, a customer-managed key can be used instead.
type EncryptionProperties_STATUS struct {
	KeySource          *string                    `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.KeyVaultReference
// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: The resource ID of the Azure key vault associated with the Batch account.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	Url       *string                       `json:"url,omitempty"`
}

// Storage version of v1api20210101.KeyVaultReference_STATUS
// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReference_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20210101.PrivateEndpointConnection_STATUS
// Contains information about a private link resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.VirtualMachineFamilyCoreQuota_STATUS
// A VM Family and its associated core quota for the Batch account.
type VirtualMachineFamilyCoreQuota_STATUS struct {
	CoreQuota   *int                   `json:"coreQuota,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.BatchAccountIdentity_UserAssignedIdentities_STATUS
type BatchAccountIdentity_UserAssignedIdentities_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.KeyVaultProperties
// KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.
type KeyVaultProperties struct {
	KeyIdentifier *string                `json:"keyIdentifier,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.KeyVaultProperties_STATUS
// KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.
type KeyVaultProperties_STATUS struct {
	KeyIdentifier *string                `json:"keyIdentifier,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210101.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BatchAccount{}, &BatchAccountList{})
}
