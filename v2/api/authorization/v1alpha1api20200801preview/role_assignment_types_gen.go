// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import (
	"fmt"
	alpha20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of RoleAssignment. Use v1beta20200801preview.RoleAssignment instead
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignments_Spec  `json:"spec,omitempty"`
	Status            RoleAssignment_Status `json:"status,omitempty"`
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
	var source alpha20200801ps.RoleAssignment

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = assignment.AssignPropertiesFromRoleAssignment(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to assignment")
	}

	return nil
}

// ConvertTo populates the provided hub RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20200801ps.RoleAssignment
	err := assignment.AssignPropertiesToRoleAssignment(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from assignment")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=default.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RoleAssignment{}

// Default applies defaults to the RoleAssignment resource
func (assignment *RoleAssignment) Default() {
	assignment.defaultImpl()
	var temp interface{} = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (assignment *RoleAssignment) defaultAzureName() {
	if assignment.Spec.AzureName == "" {
		assignment.Spec.AzureName = assignment.Name
	}
}

// defaultImpl applies the code generated defaults to the RoleAssignment resource
func (assignment *RoleAssignment) defaultImpl() { assignment.defaultAzureName() }

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01-preview"
func (assignment RoleAssignment) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (assignment *RoleAssignment) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindExtension
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
	return &RoleAssignment_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
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
	if st, ok := status.(*RoleAssignment_Status); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleAssignment_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=validate.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *RoleAssignment) ValidateCreate() error {
	validations := assignment.createValidations()
	var temp interface{} = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (assignment *RoleAssignment) ValidateDelete() error {
	validations := assignment.deleteValidations()
	var temp interface{} = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (assignment *RoleAssignment) ValidateUpdate(old runtime.Object) error {
	validations := assignment.updateValidations()
	var temp interface{} = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (assignment *RoleAssignment) createValidations() []func() error {
	return []func() error{assignment.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *RoleAssignment) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *RoleAssignment) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return assignment.validateResourceReferences()
		},
		assignment.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (assignment *RoleAssignment) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&assignment.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (assignment *RoleAssignment) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RoleAssignment)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, assignment)
}

// AssignPropertiesFromRoleAssignment populates our RoleAssignment from the provided source RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesFromRoleAssignment(source *alpha20200801ps.RoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RoleAssignments_Spec
	err := spec.AssignPropertiesFromRoleAssignmentsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignmentsSpec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status RoleAssignment_Status
	err = status.AssignPropertiesFromRoleAssignmentStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignmentStatus() to populate field Status")
	}
	assignment.Status = status

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment populates the provided destination RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesToRoleAssignment(destination *alpha20200801ps.RoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20200801ps.RoleAssignments_Spec
	err := assignment.Spec.AssignPropertiesToRoleAssignmentsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignmentsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20200801ps.RoleAssignment_Status
	err = assignment.Status.AssignPropertiesToRoleAssignmentStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignmentStatus() to populate field Status")
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
// Deprecated version of RoleAssignment. Use v1beta20200801preview.RoleAssignment instead
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Deprecated version of APIVersion. Use v1beta20200801preview.APIVersion instead
// +kubebuilder:validation:Enum={"2020-08-01-preview"}
type APIVersion string

const APIVersionValue = APIVersion("2020-08-01-preview")

// Deprecated version of RoleAssignment_Status. Use v1beta20200801preview.RoleAssignment_Status instead
type RoleAssignment_Status struct {
	Condition        *string `json:"condition,omitempty"`
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	// Conditions: The observed state of the resource
	Conditions                         []conditions.Condition                       `json:"conditions,omitempty"`
	CreatedBy                          *string                                      `json:"createdBy,omitempty"`
	CreatedOn                          *string                                      `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string                                      `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                                      `json:"description,omitempty"`
	Id                                 *string                                      `json:"id,omitempty"`
	Name                               *string                                      `json:"name,omitempty"`
	PrincipalId                        *string                                      `json:"principalId,omitempty"`
	PrincipalType                      *RoleAssignmentPropertiesStatusPrincipalType `json:"principalType,omitempty"`
	RoleDefinitionId                   *string                                      `json:"roleDefinitionId,omitempty"`
	Scope                              *string                                      `json:"scope,omitempty"`
	Type                               *string                                      `json:"type,omitempty"`
	UpdatedBy                          *string                                      `json:"updatedBy,omitempty"`
	UpdatedOn                          *string                                      `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleAssignment_Status{}

// ConvertStatusFrom populates our RoleAssignment_Status from the provided source
func (assignment *RoleAssignment_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20200801ps.RoleAssignment_Status)
	if ok {
		// Populate our instance from source
		return assignment.AssignPropertiesFromRoleAssignmentStatus(src)
	}

	// Convert to an intermediate form
	src = &alpha20200801ps.RoleAssignment_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignPropertiesFromRoleAssignmentStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_Status
func (assignment *RoleAssignment_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20200801ps.RoleAssignment_Status)
	if ok {
		// Populate destination from our instance
		return assignment.AssignPropertiesToRoleAssignmentStatus(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20200801ps.RoleAssignment_Status{}
	err := assignment.AssignPropertiesToRoleAssignmentStatus(dst)
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

var _ genruntime.FromARMConverter = &RoleAssignment_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_StatusARM, got %T", armInput)
	}

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘CreatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedBy != nil {
			createdBy := *typedInput.Properties.CreatedBy
			assignment.CreatedBy = &createdBy
		}
	}

	// Set property ‘CreatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedOn != nil {
			createdOn := *typedInput.Properties.CreatedOn
			assignment.CreatedOn = &createdOn
		}
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		assignment.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		assignment.Name = &name
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// Set property ‘RoleDefinitionId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RoleDefinitionId != nil {
			roleDefinitionId := *typedInput.Properties.RoleDefinitionId
			assignment.RoleDefinitionId = &roleDefinitionId
		}
	}

	// Set property ‘Scope’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		assignment.Type = &typeVar
	}

	// Set property ‘UpdatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedBy != nil {
			updatedBy := *typedInput.Properties.UpdatedBy
			assignment.UpdatedBy = &updatedBy
		}
	}

	// Set property ‘UpdatedOn’:
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

// AssignPropertiesFromRoleAssignmentStatus populates our RoleAssignment_Status from the provided source RoleAssignment_Status
func (assignment *RoleAssignment_Status) AssignPropertiesFromRoleAssignmentStatus(source *alpha20200801ps.RoleAssignment_Status) error {

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
		principalType := RoleAssignmentPropertiesStatusPrincipalType(*source.PrincipalType)
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

// AssignPropertiesToRoleAssignmentStatus populates the provided destination RoleAssignment_Status from our RoleAssignment_Status
func (assignment *RoleAssignment_Status) AssignPropertiesToRoleAssignmentStatus(destination *alpha20200801ps.RoleAssignment_Status) error {
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

type RoleAssignments_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                          string  `json:"azureName,omitempty"`
	Condition                          *string `json:"condition,omitempty"`
	ConditionVersion                   *string `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string `json:"description,omitempty"`
	Location                           *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalId   *string                                `json:"principalId,omitempty"`
	PrincipalType *RoleAssignmentPropertiesPrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Required
	RoleDefinitionReference *genruntime.ResourceReference `armReference:"RoleDefinitionId" json:"roleDefinitionReference,omitempty"`
	Tags                    map[string]string             `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RoleAssignments_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (assignments *RoleAssignments_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if assignments == nil {
		return nil, nil
	}
	var result RoleAssignments_SpecARM

	// Set property ‘Location’:
	if assignments.Location != nil {
		location := *assignments.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if assignments.Condition != nil ||
		assignments.ConditionVersion != nil ||
		assignments.DelegatedManagedIdentityResourceId != nil ||
		assignments.Description != nil ||
		assignments.PrincipalId != nil ||
		assignments.PrincipalType != nil ||
		assignments.RoleDefinitionReference != nil {
		result.Properties = &RoleAssignmentPropertiesARM{}
	}
	if assignments.Condition != nil {
		condition := *assignments.Condition
		result.Properties.Condition = &condition
	}
	if assignments.ConditionVersion != nil {
		conditionVersion := *assignments.ConditionVersion
		result.Properties.ConditionVersion = &conditionVersion
	}
	if assignments.DelegatedManagedIdentityResourceId != nil {
		delegatedManagedIdentityResourceId := *assignments.DelegatedManagedIdentityResourceId
		result.Properties.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
	}
	if assignments.Description != nil {
		description := *assignments.Description
		result.Properties.Description = &description
	}
	if assignments.PrincipalId != nil {
		principalId := *assignments.PrincipalId
		result.Properties.PrincipalId = &principalId
	}
	if assignments.PrincipalType != nil {
		principalType := *assignments.PrincipalType
		result.Properties.PrincipalType = &principalType
	}
	if assignments.RoleDefinitionReference != nil {
		roleDefinitionIdARMID, err := resolved.ResolvedReferences.ARMIDOrErr(*assignments.RoleDefinitionReference)
		if err != nil {
			return nil, err
		}
		roleDefinitionId := roleDefinitionIdARMID
		result.Properties.RoleDefinitionId = &roleDefinitionId
	}

	// Set property ‘Tags’:
	if assignments.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range assignments.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignments *RoleAssignments_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignments_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignments *RoleAssignments_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignments_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignments_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	assignments.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignments.Condition = &condition
		}
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignments.ConditionVersion = &conditionVersion
		}
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignments.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignments.Description = &description
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		assignments.Location = &location
	}

	// Set property ‘Owner’:
	assignments.Owner = &owner

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignments.PrincipalId = &principalId
		}
	}

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignments.PrincipalType = &principalType
		}
	}

	// no assignment for property ‘RoleDefinitionReference’

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		assignments.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			assignments.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RoleAssignments_Spec{}

// ConvertSpecFrom populates our RoleAssignments_Spec from the provided source
func (assignments *RoleAssignments_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20200801ps.RoleAssignments_Spec)
	if ok {
		// Populate our instance from source
		return assignments.AssignPropertiesFromRoleAssignmentsSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20200801ps.RoleAssignments_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignments.AssignPropertiesFromRoleAssignmentsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RoleAssignments_Spec
func (assignments *RoleAssignments_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20200801ps.RoleAssignments_Spec)
	if ok {
		// Populate destination from our instance
		return assignments.AssignPropertiesToRoleAssignmentsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20200801ps.RoleAssignments_Spec{}
	err := assignments.AssignPropertiesToRoleAssignmentsSpec(dst)
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

// AssignPropertiesFromRoleAssignmentsSpec populates our RoleAssignments_Spec from the provided source RoleAssignments_Spec
func (assignments *RoleAssignments_Spec) AssignPropertiesFromRoleAssignmentsSpec(source *alpha20200801ps.RoleAssignments_Spec) error {

	// AzureName
	assignments.AzureName = source.AzureName

	// Condition
	assignments.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignments.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	assignments.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignments.Description = genruntime.ClonePointerToString(source.Description)

	// Location
	assignments.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		assignments.Owner = &owner
	} else {
		assignments.Owner = nil
	}

	// PrincipalId
	assignments.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentPropertiesPrincipalType(*source.PrincipalType)
		assignments.PrincipalType = &principalType
	} else {
		assignments.PrincipalType = nil
	}

	// RoleDefinitionReference
	if source.RoleDefinitionReference != nil {
		roleDefinitionReference := source.RoleDefinitionReference.Copy()
		assignments.RoleDefinitionReference = &roleDefinitionReference
	} else {
		assignments.RoleDefinitionReference = nil
	}

	// Tags
	assignments.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRoleAssignmentsSpec populates the provided destination RoleAssignments_Spec from our RoleAssignments_Spec
func (assignments *RoleAssignments_Spec) AssignPropertiesToRoleAssignmentsSpec(destination *alpha20200801ps.RoleAssignments_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = assignments.AzureName

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignments.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignments.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignments.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignments.Description)

	// Location
	destination.Location = genruntime.ClonePointerToString(assignments.Location)

	// OriginalVersion
	destination.OriginalVersion = assignments.OriginalVersion()

	// Owner
	if assignments.Owner != nil {
		owner := assignments.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignments.PrincipalId)

	// PrincipalType
	if assignments.PrincipalType != nil {
		principalType := string(*assignments.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionReference
	if assignments.RoleDefinitionReference != nil {
		roleDefinitionReference := assignments.RoleDefinitionReference.Copy()
		destination.RoleDefinitionReference = &roleDefinitionReference
	} else {
		destination.RoleDefinitionReference = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(assignments.Tags)

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
func (assignments *RoleAssignments_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (assignments *RoleAssignments_Spec) SetAzureName(azureName string) {
	assignments.AzureName = azureName
}

// Deprecated version of RoleAssignmentPropertiesPrincipalType. Use
// v1beta20200801preview.RoleAssignmentPropertiesPrincipalType instead
// +kubebuilder:validation:Enum={"ForeignGroup","Group","ServicePrincipal","User"}
type RoleAssignmentPropertiesPrincipalType string

const (
	RoleAssignmentPropertiesPrincipalTypeForeignGroup     = RoleAssignmentPropertiesPrincipalType("ForeignGroup")
	RoleAssignmentPropertiesPrincipalTypeGroup            = RoleAssignmentPropertiesPrincipalType("Group")
	RoleAssignmentPropertiesPrincipalTypeServicePrincipal = RoleAssignmentPropertiesPrincipalType("ServicePrincipal")
	RoleAssignmentPropertiesPrincipalTypeUser             = RoleAssignmentPropertiesPrincipalType("User")
)

// Deprecated version of RoleAssignmentPropertiesStatusPrincipalType. Use
// v1beta20200801preview.RoleAssignmentPropertiesStatusPrincipalType instead
type RoleAssignmentPropertiesStatusPrincipalType string

const (
	RoleAssignmentPropertiesStatusPrincipalTypeForeignGroup     = RoleAssignmentPropertiesStatusPrincipalType("ForeignGroup")
	RoleAssignmentPropertiesStatusPrincipalTypeGroup            = RoleAssignmentPropertiesStatusPrincipalType("Group")
	RoleAssignmentPropertiesStatusPrincipalTypeServicePrincipal = RoleAssignmentPropertiesStatusPrincipalType("ServicePrincipal")
	RoleAssignmentPropertiesStatusPrincipalTypeUser             = RoleAssignmentPropertiesStatusPrincipalType("User")
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
