// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ipblock


type IpblockTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/ipblock#create Ipblock#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/ipblock#default Ipblock#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/ipblock#delete Ipblock#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/ipblock#update Ipblock#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

