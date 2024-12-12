// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway


type VpnWireguardGatewayConnections struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/vpn_wireguard_gateway#datacenter_id VpnWireguardGateway#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/vpn_wireguard_gateway#lan_id VpnWireguardGateway#lan_id}.
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/vpn_wireguard_gateway#ipv4_cidr VpnWireguardGateway#ipv4_cidr}.
	Ipv4Cidr *string `field:"optional" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/vpn_wireguard_gateway#ipv6_cidr VpnWireguardGateway#ipv6_cidr}.
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
}

