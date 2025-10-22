// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8snodepool


type K8SNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/k8s_node_pool#create K8SNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/k8s_node_pool#default K8SNodePool#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/k8s_node_pool#delete K8SNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/k8s_node_pool#update K8SNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

