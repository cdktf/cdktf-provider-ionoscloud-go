// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NatgatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#datacenter_id Natgateway#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// lans block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#lans Natgateway#lans}
	Lans interface{} `field:"required" json:"lans" yaml:"lans"`
	// Name of the NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#name Natgateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Collection of public IP addresses of the NAT gateway. Should be customer reserved IP addresses in that location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#public_ips Natgateway#public_ips}
	PublicIps *[]*string `field:"required" json:"publicIps" yaml:"publicIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#id Natgateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/natgateway#timeouts Natgateway#timeouts}
	Timeouts *NatgatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

