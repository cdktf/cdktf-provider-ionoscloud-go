// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationNic struct {
	// Lan ID for this replica Nic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#lan AutoscalingGroup#lan}
	Lan *float64 `field:"required" json:"lan" yaml:"lan"`
	// Name for this replica NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dhcp flag for this replica Nic.
	//
	// This is an optional attribute with default value of 'true' if not given in the request payload or given as null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#dhcp AutoscalingGroup#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// Activate or deactivate the firewall.
	//
	// By default, an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, IP addresses and ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#firewall_active AutoscalingGroup#firewall_active}
	FirewallActive interface{} `field:"optional" json:"firewallActive" yaml:"firewallActive"`
	// firewall_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#firewall_rule AutoscalingGroup#firewall_rule}
	FirewallRule interface{} `field:"optional" json:"firewallRule" yaml:"firewallRule"`
	// The type of firewall rules that will be allowed on the NIC.
	//
	// If not specified, the default INGRESS value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#firewall_type AutoscalingGroup#firewall_type}
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// flow_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#flow_log AutoscalingGroup#flow_log}
	FlowLog interface{} `field:"optional" json:"flowLog" yaml:"flowLog"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.0/docs/resources/autoscaling_group#target_group AutoscalingGroup#target_group}
	TargetGroup *AutoscalingGroupReplicaConfigurationNicTargetGroup `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

