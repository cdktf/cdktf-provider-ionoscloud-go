// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancerforwardingrule


type NetworkloadbalancerForwardingruleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/networkloadbalancer_forwardingrule#create NetworkloadbalancerForwardingrule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/networkloadbalancer_forwardingrule#default NetworkloadbalancerForwardingrule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/networkloadbalancer_forwardingrule#delete NetworkloadbalancerForwardingrule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/networkloadbalancer_forwardingrule#update NetworkloadbalancerForwardingrule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

