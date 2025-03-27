// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnIpsecTunnelConfig struct {
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
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#auth VpnIpsecTunnel#auth}
	Auth *VpnIpsecTunnelAuth `field:"required" json:"auth" yaml:"auth"`
	// The network CIDRs on the "Left" side that are allowed to connect to the IPSec tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#cloud_network_cidrs VpnIpsecTunnel#cloud_network_cidrs}
	CloudNetworkCidrs *[]*string `field:"required" json:"cloudNetworkCidrs" yaml:"cloudNetworkCidrs"`
	// esp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#esp VpnIpsecTunnel#esp}
	Esp interface{} `field:"required" json:"esp" yaml:"esp"`
	// The ID of the IPSec Gateway that the tunnel belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#gateway_id VpnIpsecTunnel#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// ike block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#ike VpnIpsecTunnel#ike}
	Ike *VpnIpsecTunnelIke `field:"required" json:"ike" yaml:"ike"`
	// The human-readable name of your IPSec Gateway Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#name VpnIpsecTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network CIDRs on the "Right" side that are allowed to connect to the IPSec tunnel.
	//
	// Specify "0.0.0.0/0" or "::/0" for all addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#peer_network_cidrs VpnIpsecTunnel#peer_network_cidrs}
	PeerNetworkCidrs *[]*string `field:"required" json:"peerNetworkCidrs" yaml:"peerNetworkCidrs"`
	// The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#remote_host VpnIpsecTunnel#remote_host}
	RemoteHost *string `field:"required" json:"remoteHost" yaml:"remoteHost"`
	// The human-readable description of your IPSec Gateway Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#description VpnIpsecTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#id VpnIpsecTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit, gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#location VpnIpsecTunnel#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/vpn_ipsec_tunnel#timeouts VpnIpsecTunnel#timeouts}
	Timeouts *VpnIpsecTunnelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

