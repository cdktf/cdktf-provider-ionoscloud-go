// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancerforwardingrule


type NetworkloadbalancerForwardingruleHealthCheck struct {
	// ClientTimeout is expressed in milliseconds.
	//
	// This inactivity timeout applies when the client is expected to acknowledge or send data. If unset the default of 50 seconds will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/networkloadbalancer_forwardingrule#client_timeout NetworkloadbalancerForwardingrule#client_timeout}
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// It specifies the maximum time (in milliseconds) to wait for a connection attempt to a target VM to succeed.
	//
	// If unset, the default of 5 seconds will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/networkloadbalancer_forwardingrule#connect_timeout NetworkloadbalancerForwardingrule#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Retries specifies the number of retries to perform on a target VM after a connection failure.
	//
	// If unset, the default value of 3 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/networkloadbalancer_forwardingrule#retries NetworkloadbalancerForwardingrule#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// TargetTimeout specifies the maximum inactivity time (in milliseconds) on the target VM side.
	//
	// If unset, the default of 50 seconds will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/networkloadbalancer_forwardingrule#target_timeout NetworkloadbalancerForwardingrule#target_timeout}
	TargetTimeout *float64 `field:"optional" json:"targetTimeout" yaml:"targetTimeout"`
}

