// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationLoadbalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#datacenter_id ApplicationLoadbalancer#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// ID of the listening (inbound) LAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#listener_lan ApplicationLoadbalancer#listener_lan}
	ListenerLan *float64 `field:"required" json:"listenerLan" yaml:"listenerLan"`
	// The name of the Application Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#name ApplicationLoadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the balanced private target LAN (outbound).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#target_lan ApplicationLoadbalancer#target_lan}
	TargetLan *float64 `field:"required" json:"targetLan" yaml:"targetLan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#id ApplicationLoadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Collection of the Application Load Balancer IP addresses.
	//
	// (Inbound and outbound) IPs of the listenerLan are customer-reserved public IPs for the public Load Balancers, and private IPs for the private Load Balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#ips ApplicationLoadbalancer#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// Collection of private IP addresses with the subnet mask of the Application Load Balancer.
	//
	// IPs must contain valid a subnet mask. If no IP is provided, the system will generate an IP with /24 subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#lb_private_ips ApplicationLoadbalancer#lb_private_ips}
	LbPrivateIps *[]*string `field:"optional" json:"lbPrivateIps" yaml:"lbPrivateIps"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.10/docs/resources/application_loadbalancer#timeouts ApplicationLoadbalancer#timeouts}
	Timeouts *ApplicationLoadbalancerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

