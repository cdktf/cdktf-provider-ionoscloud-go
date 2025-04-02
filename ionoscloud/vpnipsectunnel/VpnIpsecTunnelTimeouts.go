// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel


type VpnIpsecTunnelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/vpn_ipsec_tunnel#create VpnIpsecTunnel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/vpn_ipsec_tunnel#default VpnIpsecTunnel#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/vpn_ipsec_tunnel#delete VpnIpsecTunnel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/vpn_ipsec_tunnel#update VpnIpsecTunnel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

