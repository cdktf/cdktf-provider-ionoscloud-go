// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver


type VcpuServerNicFirewall struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#protocol VcpuServer#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#icmp_code VcpuServer#icmp_code}.
	IcmpCode *string `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#icmp_type VcpuServer#icmp_type}.
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#name VcpuServer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#port_range_end VcpuServer#port_range_end}.
	PortRangeEnd *float64 `field:"optional" json:"portRangeEnd" yaml:"portRangeEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#port_range_start VcpuServer#port_range_start}.
	PortRangeStart *float64 `field:"optional" json:"portRangeStart" yaml:"portRangeStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#source_ip VcpuServer#source_ip}.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#source_mac VcpuServer#source_mac}.
	SourceMac *string `field:"optional" json:"sourceMac" yaml:"sourceMac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#target_ip VcpuServer#target_ip}.
	TargetIp *string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/vcpu_server#type VcpuServer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

