// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplatformnodepool


type DataplatformNodePoolAutoScaling struct {
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/dataplatform_node_pool#max_node_count DataplatformNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/dataplatform_node_pool#min_node_count DataplatformNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

