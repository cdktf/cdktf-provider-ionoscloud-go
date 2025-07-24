// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgatewayrule


type NatgatewayRuleTargetPortRange struct {
	// Target port range end associated with the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/natgateway_rule#end NatgatewayRule#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// Target port range start associated with the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/natgateway_rule#start NatgatewayRule#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
}

