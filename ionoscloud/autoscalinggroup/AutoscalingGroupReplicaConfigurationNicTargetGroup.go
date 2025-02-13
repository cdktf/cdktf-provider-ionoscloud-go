// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationNicTargetGroup struct {
	// The port for the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/autoscaling_group#port AutoscalingGroup#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The ID of the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/autoscaling_group#target_group_id AutoscalingGroup#target_group_id}
	TargetGroupId *string `field:"required" json:"targetGroupId" yaml:"targetGroupId"`
	// The weight for the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/autoscaling_group#weight AutoscalingGroup#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

