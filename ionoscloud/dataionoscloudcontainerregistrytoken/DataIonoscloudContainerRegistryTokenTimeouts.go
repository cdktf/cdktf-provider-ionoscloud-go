// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudcontainerregistrytoken


type DataIonoscloudContainerRegistryTokenTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/container_registry_token#create DataIonoscloudContainerRegistryToken#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/container_registry_token#default DataIonoscloudContainerRegistryToken#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/container_registry_token#delete DataIonoscloudContainerRegistryToken#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/container_registry_token#update DataIonoscloudContainerRegistryToken#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

