// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LanConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#datacenter_id Lan#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#id Lan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#ip_failover Lan#ip_failover}
	IpFailover interface{} `field:"optional" json:"ipFailover" yaml:"ipFailover"`
	// IPv6 CIDR block assigned to the LAN.
	//
	// Can be set to 'AUTO' for an automatically assigned address or the address can be explicitly supplied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#ipv6_cidr_block Lan#ipv6_cidr_block}
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#name Lan#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#pcc Lan#pcc}.
	Pcc *string `field:"optional" json:"pcc" yaml:"pcc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#public Lan#public}.
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/lan#timeouts Lan#timeouts}
	Timeouts *LanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

