// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20220101 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20220101"
	v1api20220101s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20220101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersConfigurationExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersConfigurationExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20220101.FlexibleServersConfiguration{},
		&v1api20220101s.FlexibleServersConfiguration{}}
}
