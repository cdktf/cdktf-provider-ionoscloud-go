// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway


type VpnWireguardGatewayConnections struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_wireguard_gateway#datacenter_id VpnWireguardGateway#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_wireguard_gateway#lan_id VpnWireguardGateway#lan_id}.
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// A LAN IPv4 address in CIDR notation that will be assigned to the VPN Gateway.
	//
	// This will be the private gateway address for LAN clients to route traffic over the VPN Gateway, this should be within the subnet already assigned to the LAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_wireguard_gateway#ipv4_cidr VpnWireguardGateway#ipv4_cidr}
	Ipv4Cidr *string `field:"optional" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// A LAN IPv6 address in CIDR notation that will be assigned to the VPN Gateway.
	//
	// This will be the private gateway address for LAN clients to route traffic over the VPN Gateway, this should be within the subnet already assigned to the LAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.20/docs/resources/vpn_wireguard_gateway#ipv6_cidr VpnWireguardGateway#ipv6_cidr}
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
}

