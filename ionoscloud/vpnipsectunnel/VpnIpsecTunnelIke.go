// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel


type VpnIpsecTunnelIke struct {
	// The Diffie-Hellman Group to use for IPSec Encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/vpn_ipsec_tunnel#diffie_hellman_group VpnIpsecTunnel#diffie_hellman_group}
	DiffieHellmanGroup *string `field:"optional" json:"diffieHellmanGroup" yaml:"diffieHellmanGroup"`
	// The encryption algorithm to use for IPSec Encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/vpn_ipsec_tunnel#encryption_algorithm VpnIpsecTunnel#encryption_algorithm}
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// The integrity algorithm to use for IPSec Encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/vpn_ipsec_tunnel#integrity_algorithm VpnIpsecTunnel#integrity_algorithm}
	IntegrityAlgorithm *string `field:"optional" json:"integrityAlgorithm" yaml:"integrityAlgorithm"`
	// The phase lifetime in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/vpn_ipsec_tunnel#lifetime VpnIpsecTunnel#lifetime}
	Lifetime *float64 `field:"optional" json:"lifetime" yaml:"lifetime"`
}

