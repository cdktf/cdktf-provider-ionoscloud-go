// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancerforwardingrule


type NetworkloadbalancerForwardingruleTargetsHealthCheck struct {
	// Check specifies whether the target VM's health is checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/networkloadbalancer_forwardingrule#check NetworkloadbalancerForwardingrule#check}
	Check interface{} `field:"optional" json:"check" yaml:"check"`
	// CheckInterval determines the duration (in milliseconds) between consecutive health checks. If unspecified a default of 2000 ms is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/networkloadbalancer_forwardingrule#check_interval NetworkloadbalancerForwardingrule#check_interval}
	CheckInterval *float64 `field:"optional" json:"checkInterval" yaml:"checkInterval"`
	// Maintenance specifies if a target VM should be marked as down, even if it is not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/networkloadbalancer_forwardingrule#maintenance NetworkloadbalancerForwardingrule#maintenance}
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
}

