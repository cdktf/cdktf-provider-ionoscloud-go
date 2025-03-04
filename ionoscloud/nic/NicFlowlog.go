// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nic


type NicFlowlog struct {
	// Specifies the traffic direction pattern. Valid values: ACCEPTED, REJECTED, ALL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/nic#action Nic#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The bucket name of an existing IONOS Object Storage bucket. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/nic#bucket Nic#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies the traffic direction pattern. Valid values: INGRESS, EGRESS, BIDIRECTIONAL. Immutable, forces re-recreation of the nic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/nic#direction Nic#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/nic#name Nic#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

