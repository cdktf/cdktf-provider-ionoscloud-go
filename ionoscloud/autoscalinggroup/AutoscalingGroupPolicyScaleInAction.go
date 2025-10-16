// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupPolicyScaleInAction struct {
	// When 'amountType=ABSOLUTE' specifies the absolute number of VMs that are removed.
	//
	// The value must be between 1 to 10. 'amountType=PERCENTAGE' specifies the percentage value that is applied to the current number of replicas of the VM Auto Scaling Group. The value must be between 1 to 200. At least one VM is always removed. Note that for 'SCALE_IN' operations, volumes are not deleted after the server is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/autoscaling_group#amount AutoscalingGroup#amount}
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The type for the given amount. Possible values are: [ABSOLUTE, PERCENTAGE].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/autoscaling_group#amount_type AutoscalingGroup#amount_type}
	AmountType *string `field:"required" json:"amountType" yaml:"amountType"`
	// If set to 'true', when deleting an replica during scale in, any attached volume will also be deleted.
	//
	// When set to 'false', all volumes remain in the datacenter and must be deleted manually. Note that every scale-out creates new volumes. When they are not deleted, they will eventually use all of your contracts resource limits. At this point, scaling out would not be possible anymore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/autoscaling_group#delete_volumes AutoscalingGroup#delete_volumes}
	DeleteVolumes interface{} `field:"required" json:"deleteVolumes" yaml:"deleteVolumes"`
	// The minimum time that elapses after the start of this scaling action until the following scaling action is started.
	//
	// While a scaling action is in progress, no second action is initiated for the same VM Auto Scaling Group. Instead, the metric is re-evaluated after the current scaling action completes (either successfully or with errors). This is currently validated with a minimum value of 2 minutes and a maximum of 24 hours. The default value is 5 minutes if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/autoscaling_group#cooldown_period AutoscalingGroup#cooldown_period}
	CooldownPeriod *string `field:"optional" json:"cooldownPeriod" yaml:"cooldownPeriod"`
	// The type of termination policy for the VM Auto Scaling Group to follow a specific pattern for scaling-in replicas.
	//
	// The default termination policy is 'OLDEST_SERVER_FIRST'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/autoscaling_group#termination_policy_type AutoscalingGroup#termination_policy_type}
	TerminationPolicyType *string `field:"optional" json:"terminationPolicyType" yaml:"terminationPolicyType"`
}

