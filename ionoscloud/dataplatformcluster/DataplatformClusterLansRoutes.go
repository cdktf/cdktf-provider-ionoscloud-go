// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplatformcluster


type DataplatformClusterLansRoutes struct {
	// IPv4 or IPv6 gateway IP for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/dataplatform_cluster#gateway DataplatformCluster#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// IPv4 or IPv6 CIDR to be routed via the interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/dataplatform_cluster#network DataplatformCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

