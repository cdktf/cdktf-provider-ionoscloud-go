// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgatewayrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NatgatewayRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#datacenter_id NatgatewayRule#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Name of the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#name NatgatewayRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#natgateway_id NatgatewayRule#natgateway_id}.
	NatgatewayId *string `field:"required" json:"natgatewayId" yaml:"natgatewayId"`
	// Public IP address of the NAT gateway rule.
	//
	// Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#public_ip NatgatewayRule#public_ip}
	PublicIp *string `field:"required" json:"publicIp" yaml:"publicIp"`
	// Source subnet of the NAT gateway rule.
	//
	// For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#source_subnet NatgatewayRule#source_subnet}
	SourceSubnet *string `field:"required" json:"sourceSubnet" yaml:"sourceSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#id NatgatewayRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Protocol of the NAT gateway rule.
	//
	// Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#protocol NatgatewayRule#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// target_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#target_port_range NatgatewayRule#target_port_range}
	TargetPortRange *NatgatewayRuleTargetPortRange `field:"optional" json:"targetPortRange" yaml:"targetPortRange"`
	// Target or destination subnet of the NAT gateway rule.
	//
	// For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#target_subnet NatgatewayRule#target_subnet}
	TargetSubnet *string `field:"optional" json:"targetSubnet" yaml:"targetSubnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#timeouts NatgatewayRule#timeouts}
	Timeouts *NatgatewayRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Type of the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/natgateway_rule#type NatgatewayRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

