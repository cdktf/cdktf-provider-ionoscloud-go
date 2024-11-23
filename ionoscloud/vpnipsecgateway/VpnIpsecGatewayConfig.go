// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsecgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnIpsecGatewayConfig struct {
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
	// connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#connections VpnIpsecGateway#connections}
	Connections interface{} `field:"required" json:"connections" yaml:"connections"`
	// Public IP address to be assigned to the gateway.
	//
	// Note: This must be an IP address in the same datacenter as the connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#gateway_ip VpnIpsecGateway#gateway_ip}
	GatewayIp *string `field:"required" json:"gatewayIp" yaml:"gatewayIp"`
	// The location of the IPSec Gateway. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#location VpnIpsecGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The human readable name of your IPSecGateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#name VpnIpsecGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The human-readable description of your IPSec Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#description VpnIpsecGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#id VpnIpsecGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#timeouts VpnIpsecGateway#timeouts}
	Timeouts *VpnIpsecGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The IKE version that is permitted for the VPN tunnels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/vpn_ipsec_gateway#version VpnIpsecGateway#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

