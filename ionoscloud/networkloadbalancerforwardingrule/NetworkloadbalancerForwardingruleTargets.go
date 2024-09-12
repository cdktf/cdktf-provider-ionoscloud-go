// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancerforwardingrule


type NetworkloadbalancerForwardingruleTargets struct {
	// IP of a balanced target VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/networkloadbalancer_forwardingrule#ip NetworkloadbalancerForwardingrule#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Port of the balanced target service. (range: 1 to 65535).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/networkloadbalancer_forwardingrule#port NetworkloadbalancerForwardingrule#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Weight parameter is used to adjust the target VM's weight relative to other target VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/networkloadbalancer_forwardingrule#weight NetworkloadbalancerForwardingrule#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/networkloadbalancer_forwardingrule#health_check NetworkloadbalancerForwardingrule#health_check}
	HealthCheck *NetworkloadbalancerForwardingruleTargetsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Proxy protocol version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/networkloadbalancer_forwardingrule#proxy_protocol NetworkloadbalancerForwardingrule#proxy_protocol}
	ProxyProtocol *string `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
}

