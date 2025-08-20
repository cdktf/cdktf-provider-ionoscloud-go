// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationNicFlowLog struct {
	// Specifies the traffic direction pattern. Valid values: ACCEPTED, REJECTED, ALL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/autoscaling_group#action AutoscalingGroup#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The bucket name of an existing IONOS Object Storage bucket. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/autoscaling_group#bucket AutoscalingGroup#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies the traffic direction pattern. Valid values: INGRESS, EGRESS, BIDIRECTIONAL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/autoscaling_group#direction AutoscalingGroup#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

