// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudautocertificateprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudAutoCertificateProviderConfig struct {
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
	// The ID of the auto-certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/auto_certificate_provider#id DataIonoscloudAutoCertificateProvider#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the auto-certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/auto_certificate_provider#location DataIonoscloudAutoCertificateProvider#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the auto-certificate provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/auto_certificate_provider#name DataIonoscloudAutoCertificateProvider#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/auto_certificate_provider#timeouts DataIonoscloudAutoCertificateProvider#timeouts}
	Timeouts *DataIonoscloudAutoCertificateProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

