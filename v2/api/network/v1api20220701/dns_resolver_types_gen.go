// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"fmt"
	v1api20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701storage"
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
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}
type DnsResolver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsResolver_Spec   `json:"spec,omitempty"`
	Status            DnsResolver_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsResolver{}

// GetConditions returns the conditions of the resource
func (resolver *DnsResolver) GetConditions() conditions.Conditions {
	return resolver.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (resolver *DnsResolver) SetConditions(conditions conditions.Conditions) {
	resolver.Status.Conditions = conditions
}

var _ conversion.Convertible = &DnsResolver{}

// ConvertFrom populates our DnsResolver from the provided hub DnsResolver
func (resolver *DnsResolver) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220701s.DnsResolver)
	if !ok {
		return fmt.Errorf("expected network/v1api20220701storage/DnsResolver but received %T instead", hub)
	}

	return resolver.AssignProperties_From_DnsResolver(source)
}

// ConvertTo populates the provided hub DnsResolver from our DnsResolver
func (resolver *DnsResolver) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220701s.DnsResolver)
	if !ok {
		return fmt.Errorf("expected network/v1api20220701storage/DnsResolver but received %T instead", hub)
	}

	return resolver.AssignProperties_To_DnsResolver(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20220701-dnsresolver,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolvers,verbs=create;update,versions=v1api20220701,name=default.v1api20220701.dnsresolvers.network.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &DnsResolver{}

// Default applies defaults to the DnsResolver resource
func (resolver *DnsResolver) Default() {
	resolver.defaultImpl()
	var temp any = resolver
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (resolver *DnsResolver) defaultAzureName() {
	if resolver.Spec.AzureName == "" {
		resolver.Spec.AzureName = resolver.Name
	}
}

// defaultImpl applies the code generated defaults to the DnsResolver resource
func (resolver *DnsResolver) defaultImpl() { resolver.defaultAzureName() }

var _ genruntime.ImportableResource = &DnsResolver{}

// InitializeSpec initializes the spec for this resource from the given status
func (resolver *DnsResolver) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*DnsResolver_STATUS); ok {
		return resolver.Spec.Initialize_From_DnsResolver_STATUS(s)
	}

	return fmt.Errorf("expected Status of type DnsResolver_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &DnsResolver{}

// AzureName returns the Azure name of the resource
func (resolver *DnsResolver) AzureName() string {
	return resolver.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (resolver DnsResolver) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (resolver *DnsResolver) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (resolver *DnsResolver) GetSpec() genruntime.ConvertibleSpec {
	return &resolver.Spec
}

// GetStatus returns the status of this resource
func (resolver *DnsResolver) GetStatus() genruntime.ConvertibleStatus {
	return &resolver.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers"
func (resolver *DnsResolver) GetType() string {
	return "Microsoft.Network/dnsResolvers"
}

// NewEmptyStatus returns a new empty (blank) status
func (resolver *DnsResolver) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsResolver_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (resolver *DnsResolver) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(resolver.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  resolver.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (resolver *DnsResolver) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsResolver_STATUS); ok {
		resolver.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsResolver_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	resolver.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20220701-dnsresolver,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolvers,verbs=create;update,versions=v1api20220701,name=validate.v1api20220701.dnsresolvers.network.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &DnsResolver{}

// ValidateCreate validates the creation of the resource
func (resolver *DnsResolver) ValidateCreate() (admission.Warnings, error) {
	validations := resolver.createValidations()
	var temp any = resolver
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (resolver *DnsResolver) ValidateDelete() (admission.Warnings, error) {
	validations := resolver.deleteValidations()
	var temp any = resolver
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (resolver *DnsResolver) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := resolver.updateValidations()
	var temp any = resolver
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (resolver *DnsResolver) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){resolver.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (resolver *DnsResolver) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (resolver *DnsResolver) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return resolver.validateResourceReferences()
		},
		resolver.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (resolver *DnsResolver) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&resolver.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (resolver *DnsResolver) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*DnsResolver)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, resolver)
}

// AssignProperties_From_DnsResolver populates our DnsResolver from the provided source DnsResolver
func (resolver *DnsResolver) AssignProperties_From_DnsResolver(source *v1api20220701s.DnsResolver) error {

	// ObjectMeta
	resolver.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DnsResolver_Spec
	err := spec.AssignProperties_From_DnsResolver_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DnsResolver_Spec() to populate field Spec")
	}
	resolver.Spec = spec

	// Status
	var status DnsResolver_STATUS
	err = status.AssignProperties_From_DnsResolver_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DnsResolver_STATUS() to populate field Status")
	}
	resolver.Status = status

	// No error
	return nil
}

// AssignProperties_To_DnsResolver populates the provided destination DnsResolver from our DnsResolver
func (resolver *DnsResolver) AssignProperties_To_DnsResolver(destination *v1api20220701s.DnsResolver) error {

	// ObjectMeta
	destination.ObjectMeta = *resolver.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220701s.DnsResolver_Spec
	err := resolver.Spec.AssignProperties_To_DnsResolver_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DnsResolver_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220701s.DnsResolver_STATUS
	err = resolver.Status.AssignProperties_To_DnsResolver_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DnsResolver_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (resolver *DnsResolver) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: resolver.Spec.OriginalVersion(),
		Kind:    "DnsResolver",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}
type DnsResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsResolver `json:"items"`
}

type DnsResolver_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// +kubebuilder:validation:Required
	// VirtualNetwork: The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *DnsresolverSubResource `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ARMTransformer = &DnsResolver_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resolver *DnsResolver_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resolver == nil {
		return nil, nil
	}
	result := &DnsResolver_Spec_ARM{}

	// Set property ‘Location’:
	if resolver.Location != nil {
		location := *resolver.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if resolver.VirtualNetwork != nil {
		result.Properties = &DnsResolverProperties_ARM{}
	}
	if resolver.VirtualNetwork != nil {
		virtualNetwork_ARM, err := (*resolver.VirtualNetwork).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		virtualNetwork := *virtualNetwork_ARM.(*DnsresolverSubResource_ARM)
		result.Properties.VirtualNetwork = &virtualNetwork
	}

	// Set property ‘Tags’:
	if resolver.Tags != nil {
		result.Tags = make(map[string]string, len(resolver.Tags))
		for key, value := range resolver.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resolver *DnsResolver_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DnsResolver_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resolver *DnsResolver_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DnsResolver_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DnsResolver_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	resolver.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		resolver.Location = &location
	}

	// Set property ‘Owner’:
	resolver.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		resolver.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			resolver.Tags[key] = value
		}
	}

	// Set property ‘VirtualNetwork’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VirtualNetwork != nil {
			var virtualNetwork1 DnsresolverSubResource
			err := virtualNetwork1.PopulateFromARM(owner, *typedInput.Properties.VirtualNetwork)
			if err != nil {
				return err
			}
			virtualNetwork := virtualNetwork1
			resolver.VirtualNetwork = &virtualNetwork
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DnsResolver_Spec{}

// ConvertSpecFrom populates our DnsResolver_Spec from the provided source
func (resolver *DnsResolver_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220701s.DnsResolver_Spec)
	if ok {
		// Populate our instance from source
		return resolver.AssignProperties_From_DnsResolver_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220701s.DnsResolver_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = resolver.AssignProperties_From_DnsResolver_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DnsResolver_Spec
func (resolver *DnsResolver_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220701s.DnsResolver_Spec)
	if ok {
		// Populate destination from our instance
		return resolver.AssignProperties_To_DnsResolver_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220701s.DnsResolver_Spec{}
	err := resolver.AssignProperties_To_DnsResolver_Spec(dst)
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

// AssignProperties_From_DnsResolver_Spec populates our DnsResolver_Spec from the provided source DnsResolver_Spec
func (resolver *DnsResolver_Spec) AssignProperties_From_DnsResolver_Spec(source *v1api20220701s.DnsResolver_Spec) error {

	// AzureName
	resolver.AzureName = source.AzureName

	// Location
	resolver.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		resolver.Owner = &owner
	} else {
		resolver.Owner = nil
	}

	// Tags
	resolver.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// VirtualNetwork
	if source.VirtualNetwork != nil {
		var virtualNetwork DnsresolverSubResource
		err := virtualNetwork.AssignProperties_From_DnsresolverSubResource(source.VirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_DnsresolverSubResource() to populate field VirtualNetwork")
		}
		resolver.VirtualNetwork = &virtualNetwork
	} else {
		resolver.VirtualNetwork = nil
	}

	// No error
	return nil
}

// AssignProperties_To_DnsResolver_Spec populates the provided destination DnsResolver_Spec from our DnsResolver_Spec
func (resolver *DnsResolver_Spec) AssignProperties_To_DnsResolver_Spec(destination *v1api20220701s.DnsResolver_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = resolver.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(resolver.Location)

	// OriginalVersion
	destination.OriginalVersion = resolver.OriginalVersion()

	// Owner
	if resolver.Owner != nil {
		owner := resolver.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(resolver.Tags)

	// VirtualNetwork
	if resolver.VirtualNetwork != nil {
		var virtualNetwork v1api20220701s.DnsresolverSubResource
		err := resolver.VirtualNetwork.AssignProperties_To_DnsresolverSubResource(&virtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_DnsresolverSubResource() to populate field VirtualNetwork")
		}
		destination.VirtualNetwork = &virtualNetwork
	} else {
		destination.VirtualNetwork = nil
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

// Initialize_From_DnsResolver_STATUS populates our DnsResolver_Spec from the provided source DnsResolver_STATUS
func (resolver *DnsResolver_Spec) Initialize_From_DnsResolver_STATUS(source *DnsResolver_STATUS) error {

	// Location
	resolver.Location = genruntime.ClonePointerToString(source.Location)

	// Tags
	resolver.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// VirtualNetwork
	if source.VirtualNetwork != nil {
		var virtualNetwork DnsresolverSubResource
		err := virtualNetwork.Initialize_From_DnsresolverSubResource_STATUS(source.VirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling Initialize_From_DnsresolverSubResource_STATUS() to populate field VirtualNetwork")
		}
		resolver.VirtualNetwork = &virtualNetwork
	} else {
		resolver.VirtualNetwork = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (resolver *DnsResolver_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (resolver *DnsResolver_Spec) SetAzureName(azureName string) { resolver.AzureName = azureName }

// Describes a DNS resolver.
type DnsResolver_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DnsResolverState: The current status of the DNS resolver. This is a read-only property and any attempt to set this value
	// will be ignored.
	DnsResolverState *DnsResolverProperties_DnsResolverState_STATUS `json:"dnsResolverState,omitempty"`

	// Etag: ETag of the DNS resolver.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ProvisioningState: The current provisioning state of the DNS resolver. This is a read-only property and any attempt to
	// set this value will be ignored.
	ProvisioningState *DnsresolverProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resourceGuid property of the DNS resolver resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// VirtualNetwork: The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *DnsresolverSubResource_STATUS `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsResolver_STATUS{}

// ConvertStatusFrom populates our DnsResolver_STATUS from the provided source
func (resolver *DnsResolver_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220701s.DnsResolver_STATUS)
	if ok {
		// Populate our instance from source
		return resolver.AssignProperties_From_DnsResolver_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220701s.DnsResolver_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = resolver.AssignProperties_From_DnsResolver_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DnsResolver_STATUS
func (resolver *DnsResolver_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220701s.DnsResolver_STATUS)
	if ok {
		// Populate destination from our instance
		return resolver.AssignProperties_To_DnsResolver_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220701s.DnsResolver_STATUS{}
	err := resolver.AssignProperties_To_DnsResolver_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DnsResolver_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resolver *DnsResolver_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DnsResolver_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resolver *DnsResolver_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DnsResolver_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DnsResolver_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘DnsResolverState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DnsResolverState != nil {
			dnsResolverState := *typedInput.Properties.DnsResolverState
			resolver.DnsResolverState = &dnsResolverState
		}
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		resolver.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resolver.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		resolver.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		resolver.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			resolver.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘ResourceGuid’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ResourceGuid != nil {
			resourceGuid := *typedInput.Properties.ResourceGuid
			resolver.ResourceGuid = &resourceGuid
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		resolver.SystemData = &systemData
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		resolver.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			resolver.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		resolver.Type = &typeVar
	}

	// Set property ‘VirtualNetwork’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VirtualNetwork != nil {
			var virtualNetwork1 DnsresolverSubResource_STATUS
			err := virtualNetwork1.PopulateFromARM(owner, *typedInput.Properties.VirtualNetwork)
			if err != nil {
				return err
			}
			virtualNetwork := virtualNetwork1
			resolver.VirtualNetwork = &virtualNetwork
		}
	}

	// No error
	return nil
}

// AssignProperties_From_DnsResolver_STATUS populates our DnsResolver_STATUS from the provided source DnsResolver_STATUS
func (resolver *DnsResolver_STATUS) AssignProperties_From_DnsResolver_STATUS(source *v1api20220701s.DnsResolver_STATUS) error {

	// Conditions
	resolver.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DnsResolverState
	if source.DnsResolverState != nil {
		dnsResolverState := DnsResolverProperties_DnsResolverState_STATUS(*source.DnsResolverState)
		resolver.DnsResolverState = &dnsResolverState
	} else {
		resolver.DnsResolverState = nil
	}

	// Etag
	resolver.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resolver.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	resolver.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	resolver.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := DnsresolverProvisioningState_STATUS(*source.ProvisioningState)
		resolver.ProvisioningState = &provisioningState
	} else {
		resolver.ProvisioningState = nil
	}

	// ResourceGuid
	resolver.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		resolver.SystemData = &systemDatum
	} else {
		resolver.SystemData = nil
	}

	// Tags
	resolver.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	resolver.Type = genruntime.ClonePointerToString(source.Type)

	// VirtualNetwork
	if source.VirtualNetwork != nil {
		var virtualNetwork DnsresolverSubResource_STATUS
		err := virtualNetwork.AssignProperties_From_DnsresolverSubResource_STATUS(source.VirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_DnsresolverSubResource_STATUS() to populate field VirtualNetwork")
		}
		resolver.VirtualNetwork = &virtualNetwork
	} else {
		resolver.VirtualNetwork = nil
	}

	// No error
	return nil
}

// AssignProperties_To_DnsResolver_STATUS populates the provided destination DnsResolver_STATUS from our DnsResolver_STATUS
func (resolver *DnsResolver_STATUS) AssignProperties_To_DnsResolver_STATUS(destination *v1api20220701s.DnsResolver_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(resolver.Conditions)

	// DnsResolverState
	if resolver.DnsResolverState != nil {
		dnsResolverState := string(*resolver.DnsResolverState)
		destination.DnsResolverState = &dnsResolverState
	} else {
		destination.DnsResolverState = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resolver.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resolver.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(resolver.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(resolver.Name)

	// ProvisioningState
	if resolver.ProvisioningState != nil {
		provisioningState := string(*resolver.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(resolver.ResourceGuid)

	// SystemData
	if resolver.SystemData != nil {
		var systemDatum v1api20220701s.SystemData_STATUS
		err := resolver.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(resolver.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(resolver.Type)

	// VirtualNetwork
	if resolver.VirtualNetwork != nil {
		var virtualNetwork v1api20220701s.DnsresolverSubResource_STATUS
		err := resolver.VirtualNetwork.AssignProperties_To_DnsresolverSubResource_STATUS(&virtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_DnsresolverSubResource_STATUS() to populate field VirtualNetwork")
		}
		destination.VirtualNetwork = &virtualNetwork
	} else {
		destination.VirtualNetwork = nil
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

type DnsResolverProperties_DnsResolverState_STATUS string

const (
	DnsResolverProperties_DnsResolverState_STATUS_Connected    = DnsResolverProperties_DnsResolverState_STATUS("Connected")
	DnsResolverProperties_DnsResolverState_STATUS_Disconnected = DnsResolverProperties_DnsResolverState_STATUS("Disconnected")
)

func init() {
	SchemeBuilder.Register(&DnsResolver{}, &DnsResolverList{})
}
