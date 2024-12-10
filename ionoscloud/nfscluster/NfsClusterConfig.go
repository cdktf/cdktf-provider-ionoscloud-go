// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nfscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NfsClusterConfig struct {
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
	// connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#connections NfsCluster#connections}
	Connections *NfsClusterConnections `field:"required" json:"connections" yaml:"connections"`
	// The name of the Network File Storage Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#name NfsCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size of the Network File Storage Cluster. Minimum size is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#size NfsCluster#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#location NfsCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#nfs NfsCluster#nfs}
	Nfs *NfsClusterNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/nfs_cluster#timeouts NfsCluster#timeouts}
	Timeouts *NfsClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

