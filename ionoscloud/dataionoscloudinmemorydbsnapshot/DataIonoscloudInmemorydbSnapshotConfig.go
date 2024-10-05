// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudinmemorydbsnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudInmemorydbSnapshotConfig struct {
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
	// The ID of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/inmemorydb_snapshot#id DataIonoscloudInmemorydbSnapshot#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The location of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/inmemorydb_snapshot#location DataIonoscloudInmemorydbSnapshot#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/inmemorydb_snapshot#timeouts DataIonoscloudInmemorydbSnapshot#timeouts}
	Timeouts *DataIonoscloudInmemorydbSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

