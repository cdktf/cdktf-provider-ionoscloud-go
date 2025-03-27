// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudsnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudSnapshotConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/snapshot#id DataIonoscloudSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location of that image/snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/snapshot#location DataIonoscloudSnapshot#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// A name of that resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/snapshot#name DataIonoscloudSnapshot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The size of the image in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/snapshot#size DataIonoscloudSnapshot#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/snapshot#timeouts DataIonoscloudSnapshot#timeouts}
	Timeouts *DataIonoscloudSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

