// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancer


type NetworkloadbalancerFlowlog struct {
	// Specifies the traffic direction pattern. Valid values: ACCEPTED, REJECTED, ALL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/networkloadbalancer#action Networkloadbalancer#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The S3 bucket name of an existing IONOS Cloud S3 bucket. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/networkloadbalancer#bucket Networkloadbalancer#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies the traffic direction pattern. Valid values: INGRESS, EGRESS, BIDIRECTIONAL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/networkloadbalancer#direction Networkloadbalancer#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/networkloadbalancer#name Networkloadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

