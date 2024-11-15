// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8snodepool


type K8SNodePoolAutoScaling struct {
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/k8s_node_pool#max_node_count K8SNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/k8s_node_pool#min_node_count K8SNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

