// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerNic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#lan Server#lan}.
	Lan *float64 `field:"required" json:"lan" yaml:"lan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#dhcp Server#dhcp}.
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// Indicates whether this NIC receives an IPv6 address through DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#dhcpv6 Server#dhcpv6}
	Dhcpv6 interface{} `field:"optional" json:"dhcpv6" yaml:"dhcpv6"`
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#firewall Server#firewall}
	Firewall interface{} `field:"optional" json:"firewall" yaml:"firewall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#firewall_active Server#firewall_active}.
	FirewallActive interface{} `field:"optional" json:"firewallActive" yaml:"firewallActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#firewall_type Server#firewall_type}.
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// Collection of IP addresses assigned to a nic.
	//
	// Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#ips Server#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// IPv6 CIDR block assigned to the NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#ipv6_cidr_block Server#ipv6_cidr_block}
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Collection for IPv6 addresses assigned to a nic.
	//
	// Explicitly assigned IPv6 addresses need to come from inside the IPv6 CIDR block assigned to the nic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#ipv6_ips Server#ipv6_ips}
	Ipv6Ips *[]*string `field:"optional" json:"ipv6Ips" yaml:"ipv6Ips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#mac Server#mac}.
	Mac *string `field:"optional" json:"mac" yaml:"mac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#name Server#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of Security Group IDs for the NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/server#security_groups_ids Server#security_groups_ids}
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
}

