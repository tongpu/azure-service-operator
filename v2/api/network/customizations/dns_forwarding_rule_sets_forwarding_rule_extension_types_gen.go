// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	v1api20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DnsForwardingRuleSetsForwardingRuleExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DnsForwardingRuleSetsForwardingRuleExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20220701.DnsForwardingRuleSetsForwardingRule{},
		&v1api20220701s.DnsForwardingRuleSetsForwardingRule{}}
}
