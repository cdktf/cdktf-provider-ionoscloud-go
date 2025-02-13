// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nfsshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NfsShareConfig struct {
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
	// client_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#client_groups NfsShare#client_groups}
	ClientGroups interface{} `field:"required" json:"clientGroups" yaml:"clientGroups"`
	// The ID of the Network File Storage Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#cluster_id NfsShare#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The directory being exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#name NfsShare#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#gid NfsShare#gid}
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl, 'fr-par, 'gb-lhr, 'es/vit, 'us/las, 'us/ewr, 'us/mci'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#location NfsShare#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The quota in MiB for the export.
	//
	// The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#quota NfsShare#quota}
	Quota *float64 `field:"optional" json:"quota" yaml:"quota"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#timeouts NfsShare#timeouts}
	Timeouts *NfsShareTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/nfs_share#uid NfsShare#uid}
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

