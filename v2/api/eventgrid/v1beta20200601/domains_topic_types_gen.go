// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import (
	"fmt"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
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
// Deprecated version of DomainsTopic. Use v1api20200601.DomainsTopic instead
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domains_Topic_Spec   `json:"spec,omitempty"`
	Status            Domains_Topic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (topic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20200601s.DomainsTopic

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = topic.AssignProperties_From_DomainsTopic(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to topic")
	}

	return nil
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20200601s.DomainsTopic
	err := topic.AssignProperties_To_DomainsTopic(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from topic")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-eventgrid-azure-com-v1beta20200601-domainstopic,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1beta20200601,name=default.v1beta20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &DomainsTopic{}

// Default applies defaults to the DomainsTopic resource
func (topic *DomainsTopic) Default() {
	topic.defaultImpl()
	var temp any = topic
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (topic *DomainsTopic) defaultAzureName() {
	if topic.Spec.AzureName == "" {
		topic.Spec.AzureName = topic.Name
	}
}

// defaultImpl applies the code generated defaults to the DomainsTopic resource
func (topic *DomainsTopic) defaultImpl() { topic.defaultAzureName() }

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *DomainsTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Domains_Topic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Domains_Topic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st Domains_Topic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventgrid-azure-com-v1beta20200601-domainstopic,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1beta20200601,name=validate.v1beta20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &DomainsTopic{}

// ValidateCreate validates the creation of the resource
func (topic *DomainsTopic) ValidateCreate() (admission.Warnings, error) {
	validations := topic.createValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (topic *DomainsTopic) ValidateDelete() (admission.Warnings, error) {
	validations := topic.deleteValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (topic *DomainsTopic) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := topic.updateValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (topic *DomainsTopic) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){topic.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (topic *DomainsTopic) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (topic *DomainsTopic) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return topic.validateResourceReferences()
		},
		topic.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (topic *DomainsTopic) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&topic.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (topic *DomainsTopic) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*DomainsTopic)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, topic)
}

// AssignProperties_From_DomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (topic *DomainsTopic) AssignProperties_From_DomainsTopic(source *v20200601s.DomainsTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Domains_Topic_Spec
	err := spec.AssignProperties_From_Domains_Topic_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Domains_Topic_Spec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status Domains_Topic_STATUS
	err = status.AssignProperties_From_Domains_Topic_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Domains_Topic_STATUS() to populate field Status")
	}
	topic.Status = status

	// No error
	return nil
}

// AssignProperties_To_DomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) AssignProperties_To_DomainsTopic(destination *v20200601s.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200601s.Domains_Topic_Spec
	err := topic.Spec.AssignProperties_To_Domains_Topic_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Domains_Topic_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200601s.Domains_Topic_STATUS
	err = topic.Status.AssignProperties_To_Domains_Topic_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Domains_Topic_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion(),
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of DomainsTopic. Use v1api20200601.DomainsTopic instead
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

type Domains_Topic_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventgrid.azure.com/Domain resource
	Owner *genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner,omitempty" kind:"Domain"`
}

var _ genruntime.ARMTransformer = &Domains_Topic_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (topic *Domains_Topic_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if topic == nil {
		return nil, nil
	}
	result := &Domains_Topic_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topic *Domains_Topic_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Domains_Topic_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topic *Domains_Topic_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Domains_Topic_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Domains_Topic_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	topic.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Owner":
	topic.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Domains_Topic_Spec{}

// ConvertSpecFrom populates our Domains_Topic_Spec from the provided source
func (topic *Domains_Topic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200601s.Domains_Topic_Spec)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Domains_Topic_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Domains_Topic_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Domains_Topic_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Domains_Topic_Spec
func (topic *Domains_Topic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200601s.Domains_Topic_Spec)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Domains_Topic_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Domains_Topic_Spec{}
	err := topic.AssignProperties_To_Domains_Topic_Spec(dst)
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

// AssignProperties_From_Domains_Topic_Spec populates our Domains_Topic_Spec from the provided source Domains_Topic_Spec
func (topic *Domains_Topic_Spec) AssignProperties_From_Domains_Topic_Spec(source *v20200601s.Domains_Topic_Spec) error {

	// AzureName
	topic.AzureName = source.AzureName

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topic.Owner = &owner
	} else {
		topic.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Domains_Topic_Spec populates the provided destination Domains_Topic_Spec from our Domains_Topic_Spec
func (topic *Domains_Topic_Spec) AssignProperties_To_Domains_Topic_Spec(destination *v20200601s.Domains_Topic_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = topic.AzureName

	// OriginalVersion
	destination.OriginalVersion = topic.OriginalVersion()

	// Owner
	if topic.Owner != nil {
		owner := topic.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
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
func (topic *Domains_Topic_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (topic *Domains_Topic_Spec) SetAzureName(azureName string) { topic.AzureName = azureName }

// Deprecated version of Domains_Topic_STATUS. Use v1api20200601.Domains_Topic_STATUS instead
type Domains_Topic_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions        []conditions.Condition                          `json:"conditions,omitempty"`
	Id                *string                                         `json:"id,omitempty"`
	Name              *string                                         `json:"name,omitempty"`
	ProvisioningState *DomainTopicProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS                              `json:"systemData,omitempty"`
	Type              *string                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Domains_Topic_STATUS{}

// ConvertStatusFrom populates our Domains_Topic_STATUS from the provided source
func (topic *Domains_Topic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200601s.Domains_Topic_STATUS)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Domains_Topic_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Domains_Topic_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Domains_Topic_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200601s.Domains_Topic_STATUS)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Domains_Topic_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Domains_Topic_STATUS{}
	err := topic.AssignProperties_To_Domains_Topic_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Domains_Topic_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topic *Domains_Topic_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Domains_Topic_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topic *Domains_Topic_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Domains_Topic_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Domains_Topic_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		topic.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		topic.Name = &name
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			topic.ProvisioningState = &provisioningState
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		topic.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		topic.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Domains_Topic_STATUS populates our Domains_Topic_STATUS from the provided source Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) AssignProperties_From_Domains_Topic_STATUS(source *v20200601s.Domains_Topic_STATUS) error {

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := DomainTopicProperties_ProvisioningState_STATUS(*source.ProvisioningState)
		topic.ProvisioningState = &provisioningState
	} else {
		topic.ProvisioningState = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Domains_Topic_STATUS populates the provided destination Domains_Topic_STATUS from our Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) AssignProperties_To_Domains_Topic_STATUS(destination *v20200601s.Domains_Topic_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// ProvisioningState
	if topic.ProvisioningState != nil {
		provisioningState := string(*topic.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v20200601s.SystemData_STATUS
		err := topic.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of DomainTopicProperties_ProvisioningState_STATUS. Use
// v1api20200601.DomainTopicProperties_ProvisioningState_STATUS instead
type DomainTopicProperties_ProvisioningState_STATUS string

const (
	DomainTopicProperties_ProvisioningState_STATUS_Canceled  = DomainTopicProperties_ProvisioningState_STATUS("Canceled")
	DomainTopicProperties_ProvisioningState_STATUS_Creating  = DomainTopicProperties_ProvisioningState_STATUS("Creating")
	DomainTopicProperties_ProvisioningState_STATUS_Deleting  = DomainTopicProperties_ProvisioningState_STATUS("Deleting")
	DomainTopicProperties_ProvisioningState_STATUS_Failed    = DomainTopicProperties_ProvisioningState_STATUS("Failed")
	DomainTopicProperties_ProvisioningState_STATUS_Succeeded = DomainTopicProperties_ProvisioningState_STATUS("Succeeded")
	DomainTopicProperties_ProvisioningState_STATUS_Updating  = DomainTopicProperties_ProvisioningState_STATUS("Updating")
)

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
