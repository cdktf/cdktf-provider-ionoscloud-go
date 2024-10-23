// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsecgateway


type VpnIpsecGatewayConnections struct {
	// The datacenter to connect your VPN Gateway to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/vpn_ipsec_gateway#datacenter_id VpnIpsecGateway#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Describes the private ipv4 subnet in your LAN that should be accessible by the VPN Gateway.
	//
	// Note: this should be the subnet already assigned to the LAN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/vpn_ipsec_gateway#ipv4_cidr VpnIpsecGateway#ipv4_cidr}
	Ipv4Cidr *string `field:"required" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// The numeric LAN ID to connect your VPN Gateway to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/vpn_ipsec_gateway#lan_id VpnIpsecGateway#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// Describes the ipv6 subnet in your LAN that should be accessible by the VPN Gateway.
	//
	// Note: this should be the subnet already assigned to the LAN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/vpn_ipsec_gateway#ipv6_cidr VpnIpsecGateway#ipv6_cidr}
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
}

