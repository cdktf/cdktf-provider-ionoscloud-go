// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8snodepool


type K8SNodePoolLansRoutes struct {
	// IPv4 or IPv6 Gateway IP for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/k8s_node_pool#gateway_ip K8SNodePool#gateway_ip}
	GatewayIp *string `field:"required" json:"gatewayIp" yaml:"gatewayIp"`
	// IPv4 or IPv6 CIDR to be routed via the interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/k8s_node_pool#network K8SNodePool#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

