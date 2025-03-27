// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nsgfirewallrule


type NsgFirewallruleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nsg_firewallrule#create NsgFirewallrule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nsg_firewallrule#default NsgFirewallrule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nsg_firewallrule#delete NsgFirewallrule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nsg_firewallrule#update NsgFirewallrule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

