// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupPolicyScaleOutAction struct {
	// When 'amountType=ABSOLUTE' specifies the absolute number of VMs that are added.
	//
	// The value must be between 1 to 10. 'amountType=PERCENTAGE' specifies the percentage value that is applied to the current number of replicas of the VM Auto Scaling Group. The value must be between 1 to 200. At least one VM is always added or removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#amount AutoscalingGroup#amount}
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The type for the given amount. Possible values are: [ABSOLUTE, PERCENTAGE].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#amount_type AutoscalingGroup#amount_type}
	AmountType *string `field:"required" json:"amountType" yaml:"amountType"`
	// The minimum time that elapses after the start of this scaling action until the following scaling action is started.
	//
	// While a scaling action is in progress, no second action is initiated for the same VM Auto Scaling Group. Instead, the metric is re-evaluated after the current scaling action completes (either successfully or with errors). This is currently validated with a minimum value of 2 minutes and a maximum of 24 hours. The default value is 5 minutes if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/autoscaling_group#cooldown_period AutoscalingGroup#cooldown_period}
	CooldownPeriod *string `field:"optional" json:"cooldownPeriod" yaml:"cooldownPeriod"`
}

