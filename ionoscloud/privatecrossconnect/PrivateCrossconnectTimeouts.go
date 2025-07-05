// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecrossconnect


type PrivateCrossconnectTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/private_crossconnect#create PrivateCrossconnect#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/private_crossconnect#default PrivateCrossconnect#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/private_crossconnect#delete PrivateCrossconnect#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/private_crossconnect#update PrivateCrossconnect#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

