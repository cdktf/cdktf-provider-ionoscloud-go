// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nsgfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsgFirewallruleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#datacenter_id NsgFirewallrule#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#nsg_id NsgFirewallrule#nsg_id}.
	NsgId *string `field:"required" json:"nsgId" yaml:"nsgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#protocol NsgFirewallrule#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#icmp_code NsgFirewallrule#icmp_code}.
	IcmpCode *string `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#icmp_type NsgFirewallrule#icmp_type}.
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#id NsgFirewallrule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#name NsgFirewallrule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#port_range_end NsgFirewallrule#port_range_end}.
	PortRangeEnd *float64 `field:"optional" json:"portRangeEnd" yaml:"portRangeEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#port_range_start NsgFirewallrule#port_range_start}.
	PortRangeStart *float64 `field:"optional" json:"portRangeStart" yaml:"portRangeStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#source_ip NsgFirewallrule#source_ip}.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#source_mac NsgFirewallrule#source_mac}.
	SourceMac *string `field:"optional" json:"sourceMac" yaml:"sourceMac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#target_ip NsgFirewallrule#target_ip}.
	TargetIp *string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#timeouts NsgFirewallrule#timeouts}
	Timeouts *NsgFirewallruleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule#type NsgFirewallrule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

