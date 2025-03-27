// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autocertificateprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoCertificateProviderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The email address of the certificate requester.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#email AutoCertificateProvider#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The name of the certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#name AutoCertificateProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of the certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#server AutoCertificateProvider#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// external_account_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#external_account_binding AutoCertificateProvider#external_account_binding}
	ExternalAccountBinding *AutoCertificateProviderExternalAccountBinding `field:"optional" json:"externalAccountBinding" yaml:"externalAccountBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#id AutoCertificateProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#location AutoCertificateProvider#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/auto_certificate_provider#timeouts AutoCertificateProvider#timeouts}
	Timeouts *AutoCertificateProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

