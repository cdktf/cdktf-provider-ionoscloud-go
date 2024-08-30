// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplatformcluster


type DataplatformClusterLans struct {
	// The LAN ID of an existing LAN at the related data center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/dataplatform_cluster#lan_id DataplatformCluster#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/dataplatform_cluster#dhcp DataplatformCluster#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/dataplatform_cluster#routes DataplatformCluster#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

