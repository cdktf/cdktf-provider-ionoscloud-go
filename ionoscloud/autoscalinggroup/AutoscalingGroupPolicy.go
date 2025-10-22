// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupPolicy struct {
	// The Metric that should trigger the scaling actions. Metric values are checked at fixed intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#metric AutoscalingGroup#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// scale_in_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#scale_in_action AutoscalingGroup#scale_in_action}
	ScaleInAction *AutoscalingGroupPolicyScaleInAction `field:"required" json:"scaleInAction" yaml:"scaleInAction"`
	// The upper threshold for the value of the 'metric'.
	//
	// Used with the 'greater than' (>) operator. A scale-out action is triggered when this value is exceeded, specified by the 'scale_out_action' property. The value must have a lower minimum delta to the 'scale_in_threshold', depending on the metric, to avoid competing for actions simultaneously. If 'properties.policy.unit=TOTAL', a value >= 40 must be chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#scale_in_threshold AutoscalingGroup#scale_in_threshold}
	ScaleInThreshold *float64 `field:"required" json:"scaleInThreshold" yaml:"scaleInThreshold"`
	// scale_out_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#scale_out_action AutoscalingGroup#scale_out_action}
	ScaleOutAction *AutoscalingGroupPolicyScaleOutAction `field:"required" json:"scaleOutAction" yaml:"scaleOutAction"`
	// The upper threshold for the value of the 'metric'.
	//
	// Used with the 'greater than' (>) operator. A scale-out action is triggered when this value is exceeded, specified by the 'scaleOutAction' property. The value must have a lower minimum delta to the 'scaleInThreshold', depending on the metric, to avoid competing for actions simultaneously. If 'properties.policy.unit=TOTAL', a value >= 40 must be chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#scale_out_threshold AutoscalingGroup#scale_out_threshold}
	ScaleOutThreshold *float64 `field:"required" json:"scaleOutThreshold" yaml:"scaleOutThreshold"`
	// Units of the applied Metric. Possible values are: PER_HOUR, PER_MINUTE, PER_SECOND, TOTAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#unit AutoscalingGroup#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Specifies the time range for which the samples are to be aggregated. Must be >= 2 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#range AutoscalingGroup#range}
	Range *string `field:"optional" json:"range" yaml:"range"`
}

