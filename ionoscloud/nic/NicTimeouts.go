// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nic


type NicTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/nic#create Nic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/nic#default Nic#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/nic#delete Nic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/nic#update Nic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

