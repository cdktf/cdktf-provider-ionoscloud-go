// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver


type CubeServerNicFirewall struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#protocol CubeServer#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#icmp_code CubeServer#icmp_code}.
	IcmpCode *string `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#icmp_type CubeServer#icmp_type}.
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#name CubeServer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#port_range_end CubeServer#port_range_end}.
	PortRangeEnd *float64 `field:"optional" json:"portRangeEnd" yaml:"portRangeEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#port_range_start CubeServer#port_range_start}.
	PortRangeStart *float64 `field:"optional" json:"portRangeStart" yaml:"portRangeStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#source_ip CubeServer#source_ip}.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#source_mac CubeServer#source_mac}.
	SourceMac *string `field:"optional" json:"sourceMac" yaml:"sourceMac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#target_ip CubeServer#target_ip}.
	TargetIp *string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/cube_server#type CubeServer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

