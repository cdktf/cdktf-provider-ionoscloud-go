// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgateway


type NatgatewayLans struct {
	// Id for the LAN connected to the NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/natgateway#id Natgateway#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// Collection of gateway IP addresses of the NAT gateway.
	//
	// Will be auto-generated if not provided. Should ideally be an IP belonging to the same subnet as the LAN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/natgateway#gateway_ips Natgateway#gateway_ips}
	GatewayIps *[]*string `field:"optional" json:"gatewayIps" yaml:"gatewayIps"`
}

