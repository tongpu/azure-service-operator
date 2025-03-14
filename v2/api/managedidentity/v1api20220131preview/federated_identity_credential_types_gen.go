// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220131preview

import (
	"fmt"
	v1api20220131ps "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
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

var _ conversion.Convertible = &FederatedIdentityCredential{}

// ConvertFrom populates our FederatedIdentityCredential from the provided hub FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220131ps.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20220131previewstorage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_From_FederatedIdentityCredential(source)
}

// ConvertTo populates the provided hub FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220131ps.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20220131previewstorage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_To_FederatedIdentityCredential(destination)
}

// +kubebuilder:webhook:path=/mutate-managedidentity-azure-com-v1api20220131preview-federatedidentitycredential,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=federatedidentitycredentials,verbs=create;update,versions=v1api20220131preview,name=default.v1api20220131preview.federatedidentitycredentials.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &FederatedIdentityCredential{}

// Default applies defaults to the FederatedIdentityCredential resource
func (credential *FederatedIdentityCredential) Default() {
	credential.defaultImpl()
	var temp any = credential
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (credential *FederatedIdentityCredential) defaultAzureName() {
	if credential.Spec.AzureName == "" {
		credential.Spec.AzureName = credential.Name
	}
}

// defaultImpl applies the code generated defaults to the FederatedIdentityCredential resource
func (credential *FederatedIdentityCredential) defaultImpl() { credential.defaultAzureName() }

var _ genruntime.ImportableResource = &FederatedIdentityCredential{}

// InitializeSpec initializes the spec for this resource from the given status
func (credential *FederatedIdentityCredential) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*UserAssignedIdentities_FederatedIdentityCredential_STATUS); ok {
		return credential.Spec.Initialize_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(s)
	}

	return fmt.Errorf("expected Status of type UserAssignedIdentities_FederatedIdentityCredential_STATUS but received %T instead", status)
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

// +kubebuilder:webhook:path=/validate-managedidentity-azure-com-v1api20220131preview-federatedidentitycredential,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=federatedidentitycredentials,verbs=create;update,versions=v1api20220131preview,name=validate.v1api20220131preview.federatedidentitycredentials.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &FederatedIdentityCredential{}

// ValidateCreate validates the creation of the resource
func (credential *FederatedIdentityCredential) ValidateCreate() (admission.Warnings, error) {
	validations := credential.createValidations()
	var temp any = credential
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (credential *FederatedIdentityCredential) ValidateDelete() (admission.Warnings, error) {
	validations := credential.deleteValidations()
	var temp any = credential
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (credential *FederatedIdentityCredential) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := credential.updateValidations()
	var temp any = credential
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (credential *FederatedIdentityCredential) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){credential.validateResourceReferences, credential.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (credential *FederatedIdentityCredential) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (credential *FederatedIdentityCredential) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return credential.validateResourceReferences()
		},
		credential.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return credential.validateOptionalConfigMapReferences()
		},
	}
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (credential *FederatedIdentityCredential) validateOptionalConfigMapReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&credential.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateOptionalConfigMapReferences(refs)
}

// validateResourceReferences validates all resource references
func (credential *FederatedIdentityCredential) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&credential.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (credential *FederatedIdentityCredential) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*FederatedIdentityCredential)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, credential)
}

// AssignProperties_From_FederatedIdentityCredential populates our FederatedIdentityCredential from the provided source FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_From_FederatedIdentityCredential(source *v1api20220131ps.FederatedIdentityCredential) error {

	// ObjectMeta
	credential.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec UserAssignedIdentities_FederatedIdentityCredential_Spec
	err := spec.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	credential.Spec = spec

	// Status
	var status UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err = status.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	credential.Status = status

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredential populates the provided destination FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_To_FederatedIdentityCredential(destination *v1api20220131ps.FederatedIdentityCredential) error {

	// ObjectMeta
	destination.ObjectMeta = *credential.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec
	err := credential.Spec.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err = credential.Status.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (credential *FederatedIdentityCredential) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: credential.Spec.OriginalVersion(),
		Kind:    "FederatedIdentityCredential",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// +kubebuilder:validation:Enum={"2022-01-31-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-31-preview")

type UserAssignedIdentities_FederatedIdentityCredential_Spec struct {
	// +kubebuilder:validation:Required
	// Audiences: The list of audiences that can appear in the issued token.
	Audiences []string `json:"audiences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Issuer: The URL of the issuer to be trusted.
	Issuer *string `json:"issuer,omitempty" optionalConfigMapPair:"Issuer"`

	// IssuerFromConfig: The URL of the issuer to be trusted.
	IssuerFromConfig *genruntime.ConfigMapReference `json:"issuerFromConfig,omitempty" optionalConfigMapPair:"Issuer"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a managedidentity.azure.com/UserAssignedIdentity resource
	Owner *genruntime.KnownResourceReference `group:"managedidentity.azure.com" json:"owner,omitempty" kind:"UserAssignedIdentity"`

	// Subject: The identifier of the external identity.
	Subject *string `json:"subject,omitempty" optionalConfigMapPair:"Subject"`

	// SubjectFromConfig: The identifier of the external identity.
	SubjectFromConfig *genruntime.ConfigMapReference `json:"subjectFromConfig,omitempty" optionalConfigMapPair:"Subject"`
}

var _ genruntime.ARMTransformer = &UserAssignedIdentities_FederatedIdentityCredential_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if credential == nil {
		return nil, nil
	}
	result := &UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if credential.Audiences != nil ||
		credential.Issuer != nil ||
		credential.IssuerFromConfig != nil ||
		credential.Subject != nil ||
		credential.SubjectFromConfig != nil {
		result.Properties = &FederatedIdentityCredentialProperties_ARM{}
	}
	for _, item := range credential.Audiences {
		result.Properties.Audiences = append(result.Properties.Audiences, item)
	}
	if credential.Issuer != nil {
		issuer := *credential.Issuer
		result.Properties.Issuer = &issuer
	}
	if credential.IssuerFromConfig != nil {
		issuerValue, err := resolved.ResolvedConfigMaps.Lookup(*credential.IssuerFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property Issuer")
		}
		issuer := issuerValue
		result.Properties.Issuer = &issuer
	}
	if credential.Subject != nil {
		subject := *credential.Subject
		result.Properties.Subject = &subject
	}
	if credential.SubjectFromConfig != nil {
		subjectValue, err := resolved.ResolvedConfigMaps.Lookup(*credential.SubjectFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property Subject")
		}
		subject := subjectValue
		result.Properties.Subject = &subject
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM, got %T", armInput)
	}

	// Set property ‘Audiences’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Audiences {
			credential.Audiences = append(credential.Audiences, item)
		}
	}

	// Set property ‘AzureName’:
	credential.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Issuer’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Issuer != nil {
			issuer := *typedInput.Properties.Issuer
			credential.Issuer = &issuer
		}
	}

	// no assignment for property ‘IssuerFromConfig’

	// Set property ‘Owner’:
	credential.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Subject’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Subject != nil {
			subject := *typedInput.Properties.Subject
			credential.Subject = &subject
		}
	}

	// no assignment for property ‘SubjectFromConfig’

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_FederatedIdentityCredential_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec{}
	err := credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(source *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error {

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// AzureName
	credential.AzureName = source.AzureName

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// IssuerFromConfig
	if source.IssuerFromConfig != nil {
		issuerFromConfig := source.IssuerFromConfig.Copy()
		credential.IssuerFromConfig = &issuerFromConfig
	} else {
		credential.IssuerFromConfig = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		credential.Owner = &owner
	} else {
		credential.Owner = nil
	}

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// SubjectFromConfig
	if source.SubjectFromConfig != nil {
		subjectFromConfig := source.SubjectFromConfig.Copy()
		credential.SubjectFromConfig = &subjectFromConfig
	} else {
		credential.SubjectFromConfig = nil
	}

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec populates the provided destination UserAssignedIdentities_FederatedIdentityCredential_Spec from our UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(destination *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// AzureName
	destination.AzureName = credential.AzureName

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// IssuerFromConfig
	if credential.IssuerFromConfig != nil {
		issuerFromConfig := credential.IssuerFromConfig.Copy()
		destination.IssuerFromConfig = &issuerFromConfig
	} else {
		destination.IssuerFromConfig = nil
	}

	// OriginalVersion
	destination.OriginalVersion = credential.OriginalVersion()

	// Owner
	if credential.Owner != nil {
		owner := credential.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// SubjectFromConfig
	if credential.SubjectFromConfig != nil {
		subjectFromConfig := credential.SubjectFromConfig.Copy()
		destination.SubjectFromConfig = &subjectFromConfig
	} else {
		destination.SubjectFromConfig = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) Initialize_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(source *UserAssignedIdentities_FederatedIdentityCredential_STATUS) error {

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) SetAzureName(azureName string) {
	credential.AzureName = azureName
}

type UserAssignedIdentities_FederatedIdentityCredential_STATUS struct {
	// Audiences: The list of audiences that can appear in the issued token.
	Audiences []string `json:"audiences,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Issuer: The URL of the issuer to be trusted.
	Issuer *string `json:"issuer,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Subject: The identifier of the external identity.
	Subject *string `json:"subject,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}

// ConvertStatusFrom populates our UserAssignedIdentities_FederatedIdentityCredential_STATUS from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
	err := credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM, got %T", armInput)
	}

	// Set property ‘Audiences’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Audiences {
			credential.Audiences = append(credential.Audiences, item)
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		credential.Id = &id
	}

	// Set property ‘Issuer’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Issuer != nil {
			issuer := *typedInput.Properties.Issuer
			credential.Issuer = &issuer
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		credential.Name = &name
	}

	// Set property ‘Subject’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Subject != nil {
			subject := *typedInput.Properties.Subject
			credential.Subject = &subject
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		credential.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS populates our UserAssignedIdentities_FederatedIdentityCredential_STATUS from the provided source UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(source *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error {

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// Conditions
	credential.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	credential.Id = genruntime.ClonePointerToString(source.Id)

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// Name
	credential.Name = genruntime.ClonePointerToString(source.Name)

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// Type
	credential.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS populates the provided destination UserAssignedIdentities_FederatedIdentityCredential_STATUS from our UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(destination *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(credential.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(credential.Id)

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// Name
	destination.Name = genruntime.ClonePointerToString(credential.Name)

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// Type
	destination.Type = genruntime.ClonePointerToString(credential.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
