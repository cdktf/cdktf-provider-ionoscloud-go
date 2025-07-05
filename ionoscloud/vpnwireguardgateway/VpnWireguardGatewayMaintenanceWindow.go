// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway


type VpnWireguardGatewayMaintenanceWindow struct {
	// The name of the week day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vpn_wireguard_gateway#day_of_the_week VpnWireguardGateway#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Start of the maintenance window in UTC time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vpn_wireguard_gateway#time VpnWireguardGateway#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

