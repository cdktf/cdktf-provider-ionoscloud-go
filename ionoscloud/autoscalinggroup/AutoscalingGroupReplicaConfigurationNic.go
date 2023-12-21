// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationNic struct {
	// Lan ID for this replica Nic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/autoscaling_group#lan AutoscalingGroup#lan}
	Lan *float64 `field:"required" json:"lan" yaml:"lan"`
	// Name for this replica NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dhcp flag for this replica Nic.
	//
	// This is an optional attribute with default value of 'true' if not given in the request payload or given as null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/autoscaling_group#dhcp AutoscalingGroup#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
}

