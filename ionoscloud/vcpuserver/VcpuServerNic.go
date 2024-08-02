// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver


type VcpuServerNic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#lan VcpuServer#lan}.
	Lan *float64 `field:"required" json:"lan" yaml:"lan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#dhcp VcpuServer#dhcp}.
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#dhcpv6 VcpuServer#dhcpv6}.
	Dhcpv6 interface{} `field:"optional" json:"dhcpv6" yaml:"dhcpv6"`
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#firewall VcpuServer#firewall}
	Firewall interface{} `field:"optional" json:"firewall" yaml:"firewall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#firewall_active VcpuServer#firewall_active}.
	FirewallActive interface{} `field:"optional" json:"firewallActive" yaml:"firewallActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#firewall_type VcpuServer#firewall_type}.
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// Collection of IP addresses assigned to a nic.
	//
	// Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#ips VcpuServer#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#ipv6_cidr_block VcpuServer#ipv6_cidr_block}.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#ipv6_ips VcpuServer#ipv6_ips}.
	Ipv6Ips *[]*string `field:"optional" json:"ipv6Ips" yaml:"ipv6Ips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/vcpu_server#name VcpuServer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

