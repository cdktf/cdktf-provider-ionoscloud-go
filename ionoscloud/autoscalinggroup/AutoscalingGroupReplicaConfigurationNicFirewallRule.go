// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationNicFirewallRule struct {
	// The protocol for the rule. The property cannot be modified after its creation (not allowed in update requests).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#protocol AutoscalingGroup#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Sets the allowed code (from 0 to 254) when ICMP protocol is selected. The value 'null' allows all codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#icmp_code AutoscalingGroup#icmp_code}
	IcmpCode *float64 `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Sets the allowed type (from 0 to 254) if the protocol ICMP is selected.
	//
	// The value 'null' allows all types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#icmp_type AutoscalingGroup#icmp_type}
	IcmpType *float64 `field:"optional" json:"icmpType" yaml:"icmpType"`
	// The name of the firewall rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sets the end range of the allowed port (from 1 to 65535) if the protocol TCP or UDP is selected.
	//
	// The value 'null' for 'port_range_start' and 'port_range_end' allows all ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#port_range_end AutoscalingGroup#port_range_end}
	PortRangeEnd *float64 `field:"optional" json:"portRangeEnd" yaml:"portRangeEnd"`
	// Sets the initial range of the allowed port (from 1 to 65535) if the protocol TCP or UDP is selected.
	//
	// The value 'null' for 'port_range_start' and 'port_range_end' allows all ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#port_range_start AutoscalingGroup#port_range_start}
	PortRangeStart *float64 `field:"optional" json:"portRangeStart" yaml:"portRangeStart"`
	// Only traffic originating from the respective IPv4 address is permitted. The value 'null' allows traffic from any IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#source_ip AutoscalingGroup#source_ip}
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Only traffic originating from the respective MAC address is permitted.
	//
	// Valid format: 'aa:bb:cc:dd:ee:ff'. The value 'null' allows traffic from any MAC address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#source_mac AutoscalingGroup#source_mac}
	SourceMac *string `field:"optional" json:"sourceMac" yaml:"sourceMac"`
	// If the target NIC has multiple IP addresses, only the traffic directed to the respective IP address of the NIC is allowed.
	//
	// The value 'null' allows traffic to any target IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#target_ip AutoscalingGroup#target_ip}
	TargetIp *string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// The firewall rule type. If not specified, the default value 'INGRESS' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/autoscaling_group#type AutoscalingGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

