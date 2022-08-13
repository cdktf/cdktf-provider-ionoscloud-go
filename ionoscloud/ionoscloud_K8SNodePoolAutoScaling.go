// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type K8SNodePoolAutoScaling struct {
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_node_pool#max_node_count K8SNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_node_pool#min_node_count K8SNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

