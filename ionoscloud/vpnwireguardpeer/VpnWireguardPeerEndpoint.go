// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardpeer


type VpnWireguardPeerEndpoint struct {
	// Hostname or IPV4 address that the WireGuard Server will connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/vpn_wireguard_peer#host VpnWireguardPeer#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port that the WireGuard Server will connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/vpn_wireguard_peer#port VpnWireguardPeer#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

