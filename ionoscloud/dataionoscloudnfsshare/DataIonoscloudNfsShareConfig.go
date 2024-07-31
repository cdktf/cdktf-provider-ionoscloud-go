// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudnfsshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudNfsShareConfig struct {
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
	// The ID of the Network File Storage Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#cluster_id DataIonoscloudNfsShare#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#location DataIonoscloudNfsShare#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// client_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#client_groups DataIonoscloudNfsShare#client_groups}
	ClientGroups interface{} `field:"optional" json:"clientGroups" yaml:"clientGroups"`
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#gid DataIonoscloudNfsShare#gid}
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The ID of the Network File Storage Share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#id DataIonoscloudNfsShare#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the Network File Storage Share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#name DataIonoscloudNfsShare#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether partial matching is allowed or not when using the name filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#partial_match DataIonoscloudNfsShare#partial_match}
	PartialMatch interface{} `field:"optional" json:"partialMatch" yaml:"partialMatch"`
	// The quota in MiB for the export.
	//
	// The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#quota DataIonoscloudNfsShare#quota}
	Quota *float64 `field:"optional" json:"quota" yaml:"quota"`
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/data-sources/nfs_share#uid DataIonoscloudNfsShare#uid}
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

