// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501"
	v1api20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501storage"
	v1api20230201 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	v1api20230201s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201storage"
	v1api20230202p "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	v1api20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202previewstorage"
	v20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ManagedClustersAgentPoolExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ManagedClustersAgentPoolExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20210501.ManagedClustersAgentPool{},
		&v1api20210501s.ManagedClustersAgentPool{},
		&v1api20230201.ManagedClustersAgentPool{},
		&v1api20230201s.ManagedClustersAgentPool{},
		&v1api20230202p.ManagedClustersAgentPool{},
		&v1api20230202ps.ManagedClustersAgentPool{},
		&v20210501.ManagedClustersAgentPool{},
		&v20210501s.ManagedClustersAgentPool{}}
}
