// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsecgateway


type VpnIpsecGatewayConnections struct {
	// The datacenter to connect your VPN Gateway to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_ipsec_gateway#datacenter_id VpnIpsecGateway#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// A LAN IPv4 address in CIDR notation that will be assigned to the VPN Gateway.
	//
	// This will be the private gateway address for LAN clients to route traffic over the VPN Gateway, this should be within the subnet already assigned to the LAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_ipsec_gateway#ipv4_cidr VpnIpsecGateway#ipv4_cidr}
	Ipv4Cidr *string `field:"required" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// The numeric LAN ID to connect your VPN Gateway to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_ipsec_gateway#lan_id VpnIpsecGateway#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// A LAN IPv6 address in CIDR notation that will be assigned to the VPN Gateway.
	//
	// This will be the private gateway address for LAN clients to route traffic over the VPN Gateway, this should be within the subnet already assigned to the LAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_ipsec_gateway#ipv6_cidr VpnIpsecGateway#ipv6_cidr}
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
}

