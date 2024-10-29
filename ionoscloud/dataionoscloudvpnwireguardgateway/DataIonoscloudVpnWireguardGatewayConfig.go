// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudvpnwireguardgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudVpnWireguardGatewayConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the WireGuard Gateway. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/data-sources/vpn_wireguard_gateway#location DataIonoscloudVpnWireguardGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/data-sources/vpn_wireguard_gateway#description DataIonoscloudVpnWireguardGateway#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the WireGuard Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/data-sources/vpn_wireguard_gateway#id DataIonoscloudVpnWireguardGateway#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/data-sources/vpn_wireguard_gateway#name DataIonoscloudVpnWireguardGateway#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

