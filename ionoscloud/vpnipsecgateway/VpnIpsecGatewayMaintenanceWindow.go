// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsecgateway


type VpnIpsecGatewayMaintenanceWindow struct {
	// The name of the week day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/vpn_ipsec_gateway#day_of_the_week VpnIpsecGateway#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Start of the maintenance window in UTC time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/vpn_ipsec_gateway#time VpnIpsecGateway#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

