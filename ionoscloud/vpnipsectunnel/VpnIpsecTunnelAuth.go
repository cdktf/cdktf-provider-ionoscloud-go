// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel


type VpnIpsecTunnelAuth struct {
	// The Authentication Method to use for IPSec Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/vpn_ipsec_tunnel#method VpnIpsecTunnel#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The Pre-Shared Key to use for IPSec Authentication. Note: Required if method is PSK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/vpn_ipsec_tunnel#psk_key VpnIpsecTunnel#psk_key}
	PskKey *string `field:"optional" json:"pskKey" yaml:"pskKey"`
}

