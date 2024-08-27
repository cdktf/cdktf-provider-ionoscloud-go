// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway


type VpnWireguardGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/vpn_wireguard_gateway#create VpnWireguardGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/vpn_wireguard_gateway#default VpnWireguardGateway#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/vpn_wireguard_gateway#delete VpnWireguardGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/vpn_wireguard_gateway#update VpnWireguardGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

