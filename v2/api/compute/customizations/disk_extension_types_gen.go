// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20200930 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20200930"
	v1api20200930s "github.com/Azure/azure-service-operator/v2/api/compute/v1api20200930storage"
	v20200930 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930"
	v20200930s "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DiskExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DiskExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20200930.Disk{},
		&v1api20200930s.Disk{},
		&v20200930.Disk{},
		&v20200930s.Disk{}}
}
