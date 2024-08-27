// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry


type ContainerRegistryGarbageCollectionSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/container_registry#days ContainerRegistry#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/container_registry#time ContainerRegistry#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

