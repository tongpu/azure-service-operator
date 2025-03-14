// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401storage

import (
	"fmt"
	v1api20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210401.StorageAccountsQueueServicesQueue
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_QueueServices_Queue_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_QueueServices_Queue_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueServicesQueue{}

// GetConditions returns the conditions of the resource
func (queue *StorageAccountsQueueServicesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *StorageAccountsQueueServicesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueServicesQueue{}

// ConvertFrom populates our StorageAccountsQueueServicesQueue from the provided hub StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220901s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignProperties_From_StorageAccountsQueueServicesQueue(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220901s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignProperties_To_StorageAccountsQueueServicesQueue(destination)
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueServicesQueue{}

// AzureName returns the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (queue *StorageAccountsQueueServicesQueue) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (queue *StorageAccountsQueueServicesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *StorageAccountsQueueServicesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *StorageAccountsQueueServicesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccounts_QueueServices_Queue_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (queue *StorageAccountsQueueServicesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (queue *StorageAccountsQueueServicesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccounts_QueueServices_Queue_STATUS); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_QueueServices_Queue_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// AssignProperties_From_StorageAccountsQueueServicesQueue populates our StorageAccountsQueueServicesQueue from the provided source StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignProperties_From_StorageAccountsQueueServicesQueue(source *v1api20220901s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccounts_QueueServices_Queue_Spec
	err := spec.AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status StorageAccounts_QueueServices_Queue_STATUS
	err = status.AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS() to populate field Status")
	}
	queue.Status = status

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueServicesQueue populates the provided destination StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignProperties_To_StorageAccountsQueueServicesQueue(destination *v1api20220901s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220901s.StorageAccounts_QueueServices_Queue_Spec
	err := queue.Spec.AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS
	err = queue.Status.AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210401.StorageAccountsQueueServicesQueue
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

type augmentConversionForStorageAccountsQueueServicesQueue interface {
	AssignPropertiesFrom(src *v1api20220901s.StorageAccountsQueueServicesQueue) error
	AssignPropertiesTo(dst *v1api20220901s.StorageAccountsQueueServicesQueue) error
}

// Storage version of v1api20210401.StorageAccounts_QueueServices_Queue_Spec
type StorageAccounts_QueueServices_Queue_Spec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string            `json:"azureName,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	OriginalVersion string            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsQueueService resource
	Owner       *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsQueueService"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_QueueServices_Queue_Spec{}

// ConvertSpecFrom populates our StorageAccounts_QueueServices_Queue_Spec from the provided source
func (queue *StorageAccounts_QueueServices_Queue_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220901s.StorageAccounts_QueueServices_Queue_Spec)
	if ok {
		// Populate our instance from source
		return queue.AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220901s.StorageAccounts_QueueServices_Queue_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queue.AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_QueueServices_Queue_Spec
func (queue *StorageAccounts_QueueServices_Queue_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220901s.StorageAccounts_QueueServices_Queue_Spec)
	if ok {
		// Populate destination from our instance
		return queue.AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220901s.StorageAccounts_QueueServices_Queue_Spec{}
	err := queue.AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec(dst)
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

// AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec populates our StorageAccounts_QueueServices_Queue_Spec from the provided source StorageAccounts_QueueServices_Queue_Spec
func (queue *StorageAccounts_QueueServices_Queue_Spec) AssignProperties_From_StorageAccounts_QueueServices_Queue_Spec(source *v1api20220901s.StorageAccounts_QueueServices_Queue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	queue.AzureName = source.AzureName

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// OriginalVersion
	queue.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		queue.Owner = &owner
	} else {
		queue.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_QueueServices_Queue_Spec interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccounts_QueueServices_Queue_Spec); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec populates the provided destination StorageAccounts_QueueServices_Queue_Spec from our StorageAccounts_QueueServices_Queue_Spec
func (queue *StorageAccounts_QueueServices_Queue_Spec) AssignProperties_To_StorageAccounts_QueueServices_Queue_Spec(destination *v1api20220901s.StorageAccounts_QueueServices_Queue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// AzureName
	destination.AzureName = queue.AzureName

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// OriginalVersion
	destination.OriginalVersion = queue.OriginalVersion

	// Owner
	if queue.Owner != nil {
		owner := queue.Owner.Copy()
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

	// Invoke the augmentConversionForStorageAccounts_QueueServices_Queue_Spec interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccounts_QueueServices_Queue_Spec); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210401.StorageAccounts_QueueServices_Queue_STATUS
type StorageAccounts_QueueServices_Queue_STATUS struct {
	ApproximateMessageCount *int                   `json:"approximateMessageCount,omitempty"`
	Conditions              []conditions.Condition `json:"conditions,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	Metadata                map[string]string      `json:"metadata,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_QueueServices_Queue_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_QueueServices_Queue_STATUS from the provided source
func (queue *StorageAccounts_QueueServices_Queue_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS)
	if ok {
		// Populate our instance from source
		return queue.AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_QueueServices_Queue_STATUS
func (queue *StorageAccounts_QueueServices_Queue_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS)
	if ok {
		// Populate destination from our instance
		return queue.AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS{}
	err := queue.AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS(dst)
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

// AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS populates our StorageAccounts_QueueServices_Queue_STATUS from the provided source StorageAccounts_QueueServices_Queue_STATUS
func (queue *StorageAccounts_QueueServices_Queue_STATUS) AssignProperties_From_StorageAccounts_QueueServices_Queue_STATUS(source *v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApproximateMessageCount
	queue.ApproximateMessageCount = genruntime.ClonePointerToInt(source.ApproximateMessageCount)

	// Conditions
	queue.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Name
	queue.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_QueueServices_Queue_STATUS interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccounts_QueueServices_Queue_STATUS); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS populates the provided destination StorageAccounts_QueueServices_Queue_STATUS from our StorageAccounts_QueueServices_Queue_STATUS
func (queue *StorageAccounts_QueueServices_Queue_STATUS) AssignProperties_To_StorageAccounts_QueueServices_Queue_STATUS(destination *v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// ApproximateMessageCount
	destination.ApproximateMessageCount = genruntime.ClonePointerToInt(queue.ApproximateMessageCount)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queue.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// Name
	destination.Name = genruntime.ClonePointerToString(queue.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_QueueServices_Queue_STATUS interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccounts_QueueServices_Queue_STATUS); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccounts_QueueServices_Queue_Spec interface {
	AssignPropertiesFrom(src *v1api20220901s.StorageAccounts_QueueServices_Queue_Spec) error
	AssignPropertiesTo(dst *v1api20220901s.StorageAccounts_QueueServices_Queue_Spec) error
}

type augmentConversionForStorageAccounts_QueueServices_Queue_STATUS interface {
	AssignPropertiesFrom(src *v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS) error
	AssignPropertiesTo(dst *v1api20220901s.StorageAccounts_QueueServices_Queue_STATUS) error
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}
