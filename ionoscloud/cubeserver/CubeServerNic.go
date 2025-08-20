// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver


type CubeServerNic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#lan CubeServer#lan}.
	Lan *float64 `field:"required" json:"lan" yaml:"lan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#dhcp CubeServer#dhcp}.
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// Indicates whether this NIC receives an IPv6 address through DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#dhcpv6 CubeServer#dhcpv6}
	Dhcpv6 interface{} `field:"optional" json:"dhcpv6" yaml:"dhcpv6"`
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#firewall CubeServer#firewall}
	Firewall *CubeServerNicFirewall `field:"optional" json:"firewall" yaml:"firewall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#firewall_active CubeServer#firewall_active}.
	FirewallActive interface{} `field:"optional" json:"firewallActive" yaml:"firewallActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#firewall_type CubeServer#firewall_type}.
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#ips CubeServer#ips}.
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// IPv6 CIDR block assigned to the NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#ipv6_cidr_block CubeServer#ipv6_cidr_block}
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Collection for IPv6 addresses assigned to a nic.
	//
	// Explicitly assigned IPv6 addresses need to come from inside the IPv6 CIDR block assigned to the nic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#ipv6_ips CubeServer#ipv6_ips}
	Ipv6Ips *[]*string `field:"optional" json:"ipv6Ips" yaml:"ipv6Ips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#mac CubeServer#mac}.
	Mac *string `field:"optional" json:"mac" yaml:"mac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#name CubeServer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of Security Group IDs for the NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/cube_server#security_groups_ids CubeServer#security_groups_ids}
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
}

