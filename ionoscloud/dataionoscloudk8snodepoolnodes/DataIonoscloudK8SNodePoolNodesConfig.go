// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudk8snodepoolnodes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudK8SNodePoolNodesConfig struct {
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
	// The UUID of an existing kubernetes cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/k8s_node_pool_nodes#k8s_cluster_id DataIonoscloudK8SNodePoolNodes#k8s_cluster_id}
	K8SClusterId *string `field:"required" json:"k8SClusterId" yaml:"k8SClusterId"`
	// The UUID of an existing nodepool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/k8s_node_pool_nodes#node_pool_id DataIonoscloudK8SNodePoolNodes#node_pool_id}
	NodePoolId *string `field:"required" json:"nodePoolId" yaml:"nodePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/k8s_node_pool_nodes#id DataIonoscloudK8SNodePoolNodes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/k8s_node_pool_nodes#timeouts DataIonoscloudK8SNodePoolNodes#timeouts}
	Timeouts *DataIonoscloudK8SNodePoolNodesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

