// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway


type ApigatewayCustomDomains struct {
	// The domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/apigateway#name Apigateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The certificate ID for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/apigateway#certificate_id Apigateway#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
}

