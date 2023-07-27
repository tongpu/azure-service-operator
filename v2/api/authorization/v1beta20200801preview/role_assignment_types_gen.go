// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200801preview

import (
	"fmt"
	v20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801previewstorage"
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
// Deprecated version of RoleAssignment. Use v1api20200801preview.RoleAssignment instead
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignment_Spec   `json:"spec,omitempty"`
	Status            RoleAssignment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *RoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *RoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ conversion.Convertible = &RoleAssignment{}

// ConvertFrom populates our RoleAssignment from the provided hub RoleAssignment
func (assignment *RoleAssignment) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20200801ps.RoleAssignment

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = assignment.AssignProperties_From_RoleAssignment(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to assignment")
	}

	return nil
}

// ConvertTo populates the provided hub RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20200801ps.RoleAssignment
	err := assignment.AssignProperties_To_RoleAssignment(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from assignment")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-authorization-azure-com-v1beta20200801preview-roleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1beta20200801preview,name=default.v1beta20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RoleAssignment{}

// Default applies defaults to the RoleAssignment resource
func (assignment *RoleAssignment) Default() {
	assignment.defaultImpl()
	var temp any = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the RoleAssignment resource
func (assignment *RoleAssignment) defaultImpl() {}

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01-preview"
func (assignment RoleAssignment) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (assignment *RoleAssignment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (assignment *RoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *RoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleAssignments"
func (assignment *RoleAssignment) GetType() string {
	return "Microsoft.Authorization/roleAssignments"
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *RoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RoleAssignment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (assignment *RoleAssignment) Owner() *genruntime.ResourceReference {
	return &genruntime.ResourceReference{
		Group: assignment.Spec.Owner.Group,
		Kind:  assignment.Spec.Owner.Kind,
		Name:  assignment.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (assignment *RoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RoleAssignment_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleAssignment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-authorization-azure-com-v1beta20200801preview-roleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1beta20200801preview,name=validate.v1beta20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *RoleAssignment) ValidateCreate() (admission.Warnings, error) {
	validations := assignment.createValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (assignment *RoleAssignment) ValidateDelete() (admission.Warnings, error) {
	validations := assignment.deleteValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (assignment *RoleAssignment) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := assignment.updateValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (assignment *RoleAssignment) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){assignment.validateResourceReferences, assignment.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *RoleAssignment) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *RoleAssignment) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateResourceReferences()
		},
		assignment.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateOptionalConfigMapReferences()
		},
	}
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (assignment *RoleAssignment) validateOptionalConfigMapReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&assignment.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateOptionalConfigMapReferences(refs)
}

// validateResourceReferences validates all resource references
func (assignment *RoleAssignment) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&assignment.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (assignment *RoleAssignment) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RoleAssignment)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, assignment)
}

// AssignProperties_From_RoleAssignment populates our RoleAssignment from the provided source RoleAssignment
func (assignment *RoleAssignment) AssignProperties_From_RoleAssignment(source *v20200801ps.RoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RoleAssignment_Spec
	err := spec.AssignProperties_From_RoleAssignment_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RoleAssignment_Spec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status RoleAssignment_STATUS
	err = status.AssignProperties_From_RoleAssignment_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RoleAssignment_STATUS() to populate field Status")
	}
	assignment.Status = status

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment populates the provided destination RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) AssignProperties_To_RoleAssignment(destination *v20200801ps.RoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200801ps.RoleAssignment_Spec
	err := assignment.Spec.AssignProperties_To_RoleAssignment_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RoleAssignment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200801ps.RoleAssignment_STATUS
	err = assignment.Status.AssignProperties_To_RoleAssignment_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RoleAssignment_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *RoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion(),
		Kind:    "RoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of RoleAssignment. Use v1api20200801preview.RoleAssignment instead
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Deprecated version of APIVersion. Use v1api20200801preview.APIVersion instead
// +kubebuilder:validation:Enum={"2020-08-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-08-01-preview")

type RoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                          string  `json:"azureName,omitempty"`
	Condition                          *string `json:"condition,omitempty"`
	ConditionVersion                   *string `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner                 *genruntime.ArbitraryOwnerReference     `json:"owner,omitempty"`
	PrincipalId           *string                                 `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalIdFromConfig *genruntime.ConfigMapReference          `json:"principalIdFromConfig,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalType         *RoleAssignmentProperties_PrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Required
	RoleDefinitionReference *genruntime.ResourceReference `armReference:"RoleDefinitionId" json:"roleDefinitionReference,omitempty"`
}

var _ genruntime.ARMTransformer = &RoleAssignment_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (assignment *RoleAssignment_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if assignment == nil {
		return nil, nil
	}
	result := &RoleAssignment_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if assignment.Condition != nil ||
		assignment.ConditionVersion != nil ||
		assignment.DelegatedManagedIdentityResourceId != nil ||
		assignment.Description != nil ||
		assignment.PrincipalId != nil ||
		assignment.PrincipalIdFromConfig != nil ||
		assignment.PrincipalType != nil ||
		assignment.RoleDefinitionReference != nil {
		result.Properties = &RoleAssignmentProperties_ARM{}
	}
	if assignment.Condition != nil {
		condition := *assignment.Condition
		result.Properties.Condition = &condition
	}
	if assignment.ConditionVersion != nil {
		conditionVersion := *assignment.ConditionVersion
		result.Properties.ConditionVersion = &conditionVersion
	}
	if assignment.DelegatedManagedIdentityResourceId != nil {
		delegatedManagedIdentityResourceId := *assignment.DelegatedManagedIdentityResourceId
		result.Properties.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
	}
	if assignment.Description != nil {
		description := *assignment.Description
		result.Properties.Description = &description
	}
	if assignment.PrincipalId != nil {
		principalId := *assignment.PrincipalId
		result.Properties.PrincipalId = &principalId
	}
	if assignment.PrincipalIdFromConfig != nil {
		principalIdValue, err := resolved.ResolvedConfigMaps.Lookup(*assignment.PrincipalIdFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property PrincipalId")
		}
		principalId := principalIdValue
		result.Properties.PrincipalId = &principalId
	}
	if assignment.PrincipalType != nil {
		principalType := *assignment.PrincipalType
		result.Properties.PrincipalType = &principalType
	}
	if assignment.RoleDefinitionReference != nil {
		roleDefinitionIdARMID, err := resolved.ResolvedReferences.Lookup(*assignment.RoleDefinitionReference)
		if err != nil {
			return nil, err
		}
		roleDefinitionId := roleDefinitionIdARMID
		result.Properties.RoleDefinitionId = &roleDefinitionId
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	assignment.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Condition":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property "ConditionVersion":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// Set property "DelegatedManagedIdentityResourceId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property "Owner":
	assignment.Owner = &owner

	// Set property "PrincipalId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// no assignment for property "PrincipalIdFromConfig"

	// Set property "PrincipalType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// no assignment for property "RoleDefinitionReference"

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RoleAssignment_Spec{}

// ConvertSpecFrom populates our RoleAssignment_Spec from the provided source
func (assignment *RoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200801ps.RoleAssignment_Spec)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_RoleAssignment_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20200801ps.RoleAssignment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_RoleAssignment_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200801ps.RoleAssignment_Spec)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_RoleAssignment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200801ps.RoleAssignment_Spec{}
	err := assignment.AssignProperties_To_RoleAssignment_Spec(dst)
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

// AssignProperties_From_RoleAssignment_Spec populates our RoleAssignment_Spec from the provided source RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignProperties_From_RoleAssignment_Spec(source *v20200801ps.RoleAssignment_Spec) error {

	// AzureName
	assignment.AzureName = source.AzureName

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		assignment.Owner = &owner
	} else {
		assignment.Owner = nil
	}

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalIdFromConfig
	if source.PrincipalIdFromConfig != nil {
		principalIdFromConfig := source.PrincipalIdFromConfig.Copy()
		assignment.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		assignment.PrincipalIdFromConfig = nil
	}

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType(*source.PrincipalType)
		assignment.PrincipalType = &principalType
	} else {
		assignment.PrincipalType = nil
	}

	// RoleDefinitionReference
	if source.RoleDefinitionReference != nil {
		roleDefinitionReference := source.RoleDefinitionReference.Copy()
		assignment.RoleDefinitionReference = &roleDefinitionReference
	} else {
		assignment.RoleDefinitionReference = nil
	}

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment_Spec populates the provided destination RoleAssignment_Spec from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignProperties_To_RoleAssignment_Spec(destination *v20200801ps.RoleAssignment_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = assignment.AzureName

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// OriginalVersion
	destination.OriginalVersion = assignment.OriginalVersion()

	// Owner
	if assignment.Owner != nil {
		owner := assignment.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalIdFromConfig
	if assignment.PrincipalIdFromConfig != nil {
		principalIdFromConfig := assignment.PrincipalIdFromConfig.Copy()
		destination.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		destination.PrincipalIdFromConfig = nil
	}

	// PrincipalType
	if assignment.PrincipalType != nil {
		principalType := string(*assignment.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionReference
	if assignment.RoleDefinitionReference != nil {
		roleDefinitionReference := assignment.RoleDefinitionReference.Copy()
		destination.RoleDefinitionReference = &roleDefinitionReference
	} else {
		destination.RoleDefinitionReference = nil
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

// OriginalVersion returns the original API version used to create the resource.
func (assignment *RoleAssignment_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (assignment *RoleAssignment_Spec) SetAzureName(azureName string) {
	assignment.AzureName = azureName
}

// Deprecated version of RoleAssignment_STATUS. Use v1api20200801preview.RoleAssignment_STATUS instead
type RoleAssignment_STATUS struct {
	Condition        *string `json:"condition,omitempty"`
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	// Conditions: The observed state of the resource
	Conditions                         []conditions.Condition                         `json:"conditions,omitempty"`
	CreatedBy                          *string                                        `json:"createdBy,omitempty"`
	CreatedOn                          *string                                        `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string                                        `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                                        `json:"description,omitempty"`
	Id                                 *string                                        `json:"id,omitempty"`
	Name                               *string                                        `json:"name,omitempty"`
	PrincipalId                        *string                                        `json:"principalId,omitempty"`
	PrincipalType                      *RoleAssignmentProperties_PrincipalType_STATUS `json:"principalType,omitempty"`
	RoleDefinitionId                   *string                                        `json:"roleDefinitionId,omitempty"`
	Scope                              *string                                        `json:"scope,omitempty"`
	Type                               *string                                        `json:"type,omitempty"`
	UpdatedBy                          *string                                        `json:"updatedBy,omitempty"`
	UpdatedOn                          *string                                        `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleAssignment_STATUS{}

// ConvertStatusFrom populates our RoleAssignment_STATUS from the provided source
func (assignment *RoleAssignment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200801ps.RoleAssignment_STATUS)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_RoleAssignment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20200801ps.RoleAssignment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_RoleAssignment_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200801ps.RoleAssignment_STATUS)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_RoleAssignment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20200801ps.RoleAssignment_STATUS{}
	err := assignment.AssignProperties_To_RoleAssignment_STATUS(dst)
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

var _ genruntime.FromARMConverter = &RoleAssignment_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_STATUS_ARM, got %T", armInput)
	}

	// Set property "Condition":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property "ConditionVersion":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// no assignment for property "Conditions"

	// Set property "CreatedBy":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedBy != nil {
			createdBy := *typedInput.Properties.CreatedBy
			assignment.CreatedBy = &createdBy
		}
	}

	// Set property "CreatedOn":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedOn != nil {
			createdOn := *typedInput.Properties.CreatedOn
			assignment.CreatedOn = &createdOn
		}
	}

	// Set property "DelegatedManagedIdentityResourceId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		assignment.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		assignment.Name = &name
	}

	// Set property "PrincipalId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// Set property "PrincipalType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// Set property "RoleDefinitionId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RoleDefinitionId != nil {
			roleDefinitionId := *typedInput.Properties.RoleDefinitionId
			assignment.RoleDefinitionId = &roleDefinitionId
		}
	}

	// Set property "Scope":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		assignment.Type = &typeVar
	}

	// Set property "UpdatedBy":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedBy != nil {
			updatedBy := *typedInput.Properties.UpdatedBy
			assignment.UpdatedBy = &updatedBy
		}
	}

	// Set property "UpdatedOn":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedOn != nil {
			updatedOn := *typedInput.Properties.UpdatedOn
			assignment.UpdatedOn = &updatedOn
		}
	}

	// No error
	return nil
}

// AssignProperties_From_RoleAssignment_STATUS populates our RoleAssignment_STATUS from the provided source RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignProperties_From_RoleAssignment_STATUS(source *v20200801ps.RoleAssignment_STATUS) error {

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedBy
	assignment.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedOn
	assignment.CreatedOn = genruntime.ClonePointerToString(source.CreatedOn)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// Id
	assignment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	assignment.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType_STATUS(*source.PrincipalType)
		assignment.PrincipalType = &principalType
	} else {
		assignment.PrincipalType = nil
	}

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// Type
	assignment.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedBy
	assignment.UpdatedBy = genruntime.ClonePointerToString(source.UpdatedBy)

	// UpdatedOn
	assignment.UpdatedOn = genruntime.ClonePointerToString(source.UpdatedOn)

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment_STATUS populates the provided destination RoleAssignment_STATUS from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignProperties_To_RoleAssignment_STATUS(destination *v20200801ps.RoleAssignment_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(assignment.CreatedBy)

	// CreatedOn
	destination.CreatedOn = genruntime.ClonePointerToString(assignment.CreatedOn)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// Id
	destination.Id = genruntime.ClonePointerToString(assignment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(assignment.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalType
	if assignment.PrincipalType != nil {
		principalType := string(*assignment.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Type
	destination.Type = genruntime.ClonePointerToString(assignment.Type)

	// UpdatedBy
	destination.UpdatedBy = genruntime.ClonePointerToString(assignment.UpdatedBy)

	// UpdatedOn
	destination.UpdatedOn = genruntime.ClonePointerToString(assignment.UpdatedOn)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of RoleAssignmentProperties_PrincipalType. Use
// v1api20200801preview.RoleAssignmentProperties_PrincipalType instead
// +kubebuilder:validation:Enum={"ForeignGroup","Group","ServicePrincipal","User"}
type RoleAssignmentProperties_PrincipalType string

const (
	RoleAssignmentProperties_PrincipalType_ForeignGroup     = RoleAssignmentProperties_PrincipalType("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_Group            = RoleAssignmentProperties_PrincipalType("Group")
	RoleAssignmentProperties_PrincipalType_ServicePrincipal = RoleAssignmentProperties_PrincipalType("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_User             = RoleAssignmentProperties_PrincipalType("User")
)

// Deprecated version of RoleAssignmentProperties_PrincipalType_STATUS. Use
// v1api20200801preview.RoleAssignmentProperties_PrincipalType_STATUS instead
type RoleAssignmentProperties_PrincipalType_STATUS string

const (
	RoleAssignmentProperties_PrincipalType_STATUS_ForeignGroup     = RoleAssignmentProperties_PrincipalType_STATUS("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_STATUS_Group            = RoleAssignmentProperties_PrincipalType_STATUS("Group")
	RoleAssignmentProperties_PrincipalType_STATUS_ServicePrincipal = RoleAssignmentProperties_PrincipalType_STATUS("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_STATUS_User             = RoleAssignmentProperties_PrincipalType_STATUS("User")
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
