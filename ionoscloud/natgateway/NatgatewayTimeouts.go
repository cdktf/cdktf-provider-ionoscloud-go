// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgateway


type NatgatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/natgateway#create Natgateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/natgateway#default Natgateway#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/natgateway#delete Natgateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/natgateway#update Natgateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

