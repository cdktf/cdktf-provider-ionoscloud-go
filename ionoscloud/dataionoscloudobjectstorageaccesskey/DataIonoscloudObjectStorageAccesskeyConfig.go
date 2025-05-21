// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudobjectstorageaccesskey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudObjectStorageAccesskeyConfig struct {
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
	// Access key metadata is a string of 92 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/data-sources/object_storage_accesskey#accesskey DataIonoscloudObjectStorageAccesskey#accesskey}
	Accesskey *string `field:"optional" json:"accesskey" yaml:"accesskey"`
	// Description of the Access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/data-sources/object_storage_accesskey#description DataIonoscloudObjectStorageAccesskey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID (UUID) of the AccessKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/data-sources/object_storage_accesskey#id DataIonoscloudObjectStorageAccesskey#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

