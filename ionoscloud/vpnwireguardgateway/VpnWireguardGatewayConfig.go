// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnWireguardGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#connections VpnWireguardGateway#connections}
	Connections interface{} `field:"required" json:"connections" yaml:"connections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#gateway_ip VpnWireguardGateway#gateway_ip}.
	GatewayIp *string `field:"required" json:"gatewayIp" yaml:"gatewayIp"`
	// The location of the WireGuard Gateway. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#location VpnWireguardGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#name VpnWireguardGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// PrivateKey used for WireGuard Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#private_key VpnWireguardGateway#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#description VpnWireguardGateway#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#id VpnWireguardGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IPV4 address (with CIDR mask) to be assigned to the WireGuard interface.
	//
	// __Note__: either interfaceIPv4CIDR or interfaceIPv6CIDR is __required__.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#interface_ipv4_cidr VpnWireguardGateway#interface_ipv4_cidr}
	InterfaceIpv4Cidr *string `field:"optional" json:"interfaceIpv4Cidr" yaml:"interfaceIpv4Cidr"`
	// The IPV6 address (with CIDR mask) to be assigned to the WireGuard interface.
	//
	// __Note__: either interfaceIPv6CIDR or interfaceIPv4CIDR is __required__.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#interface_ipv6_cidr VpnWireguardGateway#interface_ipv6_cidr}
	InterfaceIpv6Cidr *string `field:"optional" json:"interfaceIpv6Cidr" yaml:"interfaceIpv6Cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#listen_port VpnWireguardGateway#listen_port}.
	ListenPort *float64 `field:"optional" json:"listenPort" yaml:"listenPort"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/vpn_wireguard_gateway#timeouts VpnWireguardGateway#timeouts}
	Timeouts *VpnWireguardGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

