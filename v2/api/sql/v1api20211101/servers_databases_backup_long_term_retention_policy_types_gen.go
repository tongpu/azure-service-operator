// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v1api20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/LongTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/default
type ServersDatabasesBackupLongTermRetentionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_BackupLongTermRetentionPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_BackupLongTermRetentionPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesBackupLongTermRetentionPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersDatabasesBackupLongTermRetentionPolicy{}

// ConvertFrom populates our ServersDatabasesBackupLongTermRetentionPolicy from the provided hub ServersDatabasesBackupLongTermRetentionPolicy
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20211101s.ServersDatabasesBackupLongTermRetentionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersDatabasesBackupLongTermRetentionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ServersDatabasesBackupLongTermRetentionPolicy(source)
}

// ConvertTo populates the provided hub ServersDatabasesBackupLongTermRetentionPolicy from our ServersDatabasesBackupLongTermRetentionPolicy
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20211101s.ServersDatabasesBackupLongTermRetentionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersDatabasesBackupLongTermRetentionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ServersDatabasesBackupLongTermRetentionPolicy(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabasesbackuplongtermretentionpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesbackuplongtermretentionpolicies,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabasesbackuplongtermretentionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersDatabasesBackupLongTermRetentionPolicy{}

// Default applies defaults to the ServersDatabasesBackupLongTermRetentionPolicy resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) Default() {
	policy.defaultImpl()
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersDatabasesBackupLongTermRetentionPolicy resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersDatabasesBackupLongTermRetentionPolicy{}

// InitializeSpec initializes the spec for this resource from the given status
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Databases_BackupLongTermRetentionPolicy_STATUS); ok {
		return policy.Spec.Initialize_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Databases_BackupLongTermRetentionPolicy_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersDatabasesBackupLongTermRetentionPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesBackupLongTermRetentionPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) GetType() string {
	return "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  policy.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_BackupLongTermRetentionPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_BackupLongTermRetentionPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabasesbackuplongtermretentionpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesbackuplongtermretentionpolicies,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabasesbackuplongtermretentionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersDatabasesBackupLongTermRetentionPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ValidateCreate() (admission.Warnings, error) {
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ValidateDelete() (admission.Warnings, error) {
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){policy.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersDatabasesBackupLongTermRetentionPolicy)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_ServersDatabasesBackupLongTermRetentionPolicy populates our ServersDatabasesBackupLongTermRetentionPolicy from the provided source ServersDatabasesBackupLongTermRetentionPolicy
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) AssignProperties_From_ServersDatabasesBackupLongTermRetentionPolicy(source *v1api20211101s.ServersDatabasesBackupLongTermRetentionPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Databases_BackupLongTermRetentionPolicy_Spec
	err := spec.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Servers_Databases_BackupLongTermRetentionPolicy_STATUS
	err = status.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesBackupLongTermRetentionPolicy populates the provided destination ServersDatabasesBackupLongTermRetentionPolicy from our ServersDatabasesBackupLongTermRetentionPolicy
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) AssignProperties_To_ServersDatabasesBackupLongTermRetentionPolicy(destination *v1api20211101s.ServersDatabasesBackupLongTermRetentionPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec
	err := policy.Spec.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS
	err = policy.Status.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "ServersDatabasesBackupLongTermRetentionPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/LongTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/default
type ServersDatabasesBackupLongTermRetentionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesBackupLongTermRetentionPolicy `json:"items"`
}

type Servers_Databases_BackupLongTermRetentionPolicy_Spec struct {
	// MonthlyRetention: The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `json:"monthlyRetention,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// WeekOfYear: The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `json:"weekOfYear,omitempty"`

	// WeeklyRetention: The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `json:"weeklyRetention,omitempty"`

	// YearlyRetention: The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `json:"yearlyRetention,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Databases_BackupLongTermRetentionPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if policy.MonthlyRetention != nil ||
		policy.WeekOfYear != nil ||
		policy.WeeklyRetention != nil ||
		policy.YearlyRetention != nil {
		result.Properties = &BaseLongTermRetentionPolicyProperties_ARM{}
	}
	if policy.MonthlyRetention != nil {
		monthlyRetention := *policy.MonthlyRetention
		result.Properties.MonthlyRetention = &monthlyRetention
	}
	if policy.WeekOfYear != nil {
		weekOfYear := *policy.WeekOfYear
		result.Properties.WeekOfYear = &weekOfYear
	}
	if policy.WeeklyRetention != nil {
		weeklyRetention := *policy.WeeklyRetention
		result.Properties.WeeklyRetention = &weeklyRetention
	}
	if policy.YearlyRetention != nil {
		yearlyRetention := *policy.YearlyRetention
		result.Properties.YearlyRetention = &yearlyRetention
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM, got %T", armInput)
	}

	// Set property ‘MonthlyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.MonthlyRetention != nil {
			monthlyRetention := *typedInput.Properties.MonthlyRetention
			policy.MonthlyRetention = &monthlyRetention
		}
	}

	// Set property ‘Owner’:
	policy.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘WeekOfYear’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.WeekOfYear != nil {
			weekOfYear := *typedInput.Properties.WeekOfYear
			policy.WeekOfYear = &weekOfYear
		}
	}

	// Set property ‘WeeklyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.WeeklyRetention != nil {
			weeklyRetention := *typedInput.Properties.WeeklyRetention
			policy.WeeklyRetention = &weeklyRetention
		}
	}

	// Set property ‘YearlyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.YearlyRetention != nil {
			yearlyRetention := *typedInput.Properties.YearlyRetention
			policy.YearlyRetention = &yearlyRetention
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_BackupLongTermRetentionPolicy_Spec{}

// ConvertSpecFrom populates our Servers_Databases_BackupLongTermRetentionPolicy_Spec from the provided source
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_BackupLongTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec{}
	err := policy.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec(dst)
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

// AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec populates our Servers_Databases_BackupLongTermRetentionPolicy_Spec from the provided source Servers_Databases_BackupLongTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_Spec(source *v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec) error {

	// MonthlyRetention
	policy.MonthlyRetention = genruntime.ClonePointerToString(source.MonthlyRetention)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// WeekOfYear
	policy.WeekOfYear = genruntime.ClonePointerToInt(source.WeekOfYear)

	// WeeklyRetention
	policy.WeeklyRetention = genruntime.ClonePointerToString(source.WeeklyRetention)

	// YearlyRetention
	policy.YearlyRetention = genruntime.ClonePointerToString(source.YearlyRetention)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec populates the provided destination Servers_Databases_BackupLongTermRetentionPolicy_Spec from our Servers_Databases_BackupLongTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_Spec(destination *v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// MonthlyRetention
	destination.MonthlyRetention = genruntime.ClonePointerToString(policy.MonthlyRetention)

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// WeekOfYear
	destination.WeekOfYear = genruntime.ClonePointerToInt(policy.WeekOfYear)

	// WeeklyRetention
	destination.WeeklyRetention = genruntime.ClonePointerToString(policy.WeeklyRetention)

	// YearlyRetention
	destination.YearlyRetention = genruntime.ClonePointerToString(policy.YearlyRetention)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS populates our Servers_Databases_BackupLongTermRetentionPolicy_Spec from the provided source Servers_Databases_BackupLongTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) Initialize_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(source *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) error {

	// MonthlyRetention
	policy.MonthlyRetention = genruntime.ClonePointerToString(source.MonthlyRetention)

	// WeekOfYear
	policy.WeekOfYear = genruntime.ClonePointerToInt(source.WeekOfYear)

	// WeeklyRetention
	policy.WeeklyRetention = genruntime.ClonePointerToString(source.WeeklyRetention)

	// YearlyRetention
	policy.YearlyRetention = genruntime.ClonePointerToString(source.YearlyRetention)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Databases_BackupLongTermRetentionPolicy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// MonthlyRetention: The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `json:"monthlyRetention,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// WeekOfYear: The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `json:"weekOfYear,omitempty"`

	// WeeklyRetention: The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `json:"weeklyRetention,omitempty"`

	// YearlyRetention: The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `json:"yearlyRetention,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_BackupLongTermRetentionPolicy_STATUS from the provided source
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_BackupLongTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}
	err := policy.AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Databases_BackupLongTermRetentionPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_BackupLongTermRetentionPolicy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_BackupLongTermRetentionPolicy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_BackupLongTermRetentionPolicy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property ‘MonthlyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.MonthlyRetention != nil {
			monthlyRetention := *typedInput.Properties.MonthlyRetention
			policy.MonthlyRetention = &monthlyRetention
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// Set property ‘WeekOfYear’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.WeekOfYear != nil {
			weekOfYear := *typedInput.Properties.WeekOfYear
			policy.WeekOfYear = &weekOfYear
		}
	}

	// Set property ‘WeeklyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.WeeklyRetention != nil {
			weeklyRetention := *typedInput.Properties.WeeklyRetention
			policy.WeeklyRetention = &weeklyRetention
		}
	}

	// Set property ‘YearlyRetention’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.YearlyRetention != nil {
			yearlyRetention := *typedInput.Properties.YearlyRetention
			policy.YearlyRetention = &yearlyRetention
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS populates our Servers_Databases_BackupLongTermRetentionPolicy_STATUS from the provided source Servers_Databases_BackupLongTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) AssignProperties_From_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(source *v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// MonthlyRetention
	policy.MonthlyRetention = genruntime.ClonePointerToString(source.MonthlyRetention)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// WeekOfYear
	policy.WeekOfYear = genruntime.ClonePointerToInt(source.WeekOfYear)

	// WeeklyRetention
	policy.WeeklyRetention = genruntime.ClonePointerToString(source.WeeklyRetention)

	// YearlyRetention
	policy.YearlyRetention = genruntime.ClonePointerToString(source.YearlyRetention)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS populates the provided destination Servers_Databases_BackupLongTermRetentionPolicy_STATUS from our Servers_Databases_BackupLongTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_STATUS) AssignProperties_To_Servers_Databases_BackupLongTermRetentionPolicy_STATUS(destination *v1api20211101s.Servers_Databases_BackupLongTermRetentionPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// MonthlyRetention
	destination.MonthlyRetention = genruntime.ClonePointerToString(policy.MonthlyRetention)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

	// WeekOfYear
	destination.WeekOfYear = genruntime.ClonePointerToInt(policy.WeekOfYear)

	// WeeklyRetention
	destination.WeeklyRetention = genruntime.ClonePointerToString(policy.WeeklyRetention)

	// YearlyRetention
	destination.YearlyRetention = genruntime.ClonePointerToString(policy.YearlyRetention)

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
	SchemeBuilder.Register(&ServersDatabasesBackupLongTermRetentionPolicy{}, &ServersDatabasesBackupLongTermRetentionPolicyList{})
}
