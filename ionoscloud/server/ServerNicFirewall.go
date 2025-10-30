// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerNicFirewall struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#protocol Server#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#icmp_code Server#icmp_code}.
	IcmpCode *string `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#icmp_type Server#icmp_type}.
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#name Server#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#port_range_end Server#port_range_end}.
	PortRangeEnd *float64 `field:"optional" json:"portRangeEnd" yaml:"portRangeEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#port_range_start Server#port_range_start}.
	PortRangeStart *float64 `field:"optional" json:"portRangeStart" yaml:"portRangeStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#source_ip Server#source_ip}.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#source_mac Server#source_mac}.
	SourceMac *string `field:"optional" json:"sourceMac" yaml:"sourceMac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#target_ip Server#target_ip}.
	TargetIp *string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/server#type Server#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

