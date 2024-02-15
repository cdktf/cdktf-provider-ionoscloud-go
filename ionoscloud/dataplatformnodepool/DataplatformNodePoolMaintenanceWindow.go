// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplatformnodepool


type DataplatformNodePoolMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.13/docs/resources/dataplatform_node_pool#day_of_the_week DataplatformNodePool#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Time at which the maintenance should start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.13/docs/resources/dataplatform_node_pool#time DataplatformNodePool#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

