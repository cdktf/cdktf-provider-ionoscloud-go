// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autocertificateprovider


type AutoCertificateProviderExternalAccountBinding struct {
	// The key ID of the external account binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/auto_certificate_provider#key_id AutoCertificateProvider#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The secret of the external account binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.6/docs/resources/auto_certificate_provider#key_secret AutoCertificateProvider#key_secret}
	KeySecret *string `field:"required" json:"keySecret" yaml:"keySecret"`
}

