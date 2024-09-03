// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardpeer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnWireguardPeerConfig struct {
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
	// The subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#allowed_ips VpnWireguardPeer#allowed_ips}
	AllowedIps *[]*string `field:"required" json:"allowedIps" yaml:"allowedIps"`
	// The ID of the WireGuard Peer that the peer will connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#gateway_id VpnWireguardPeer#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The location of the WireGuard Peer. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#location VpnWireguardPeer#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The human readable name of your WireGuard Gateway Peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#name VpnWireguardPeer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// WireGuard public key of the connecting peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#public_key VpnWireguardPeer#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Human readable description of the WireGuard Gateway Peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#description VpnWireguardPeer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#endpoint VpnWireguardPeer#endpoint}
	Endpoint *VpnWireguardPeerEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#id VpnWireguardPeer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/vpn_wireguard_peer#timeouts VpnWireguardPeer#timeouts}
	Timeouts *VpnWireguardPeerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

