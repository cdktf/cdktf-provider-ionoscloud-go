// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudvolume


type DataIonoscloudVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/volume#create DataIonoscloudVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/volume#default DataIonoscloudVolume#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/volume#delete DataIonoscloudVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/data-sources/volume#update DataIonoscloudVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

