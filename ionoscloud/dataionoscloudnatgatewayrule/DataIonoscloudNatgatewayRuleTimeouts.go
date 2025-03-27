// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudnatgatewayrule


type DataIonoscloudNatgatewayRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/natgateway_rule#create DataIonoscloudNatgatewayRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/natgateway_rule#default DataIonoscloudNatgatewayRule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/natgateway_rule#delete DataIonoscloudNatgatewayRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/natgateway_rule#update DataIonoscloudNatgatewayRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

