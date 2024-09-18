// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudvpnipsectunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudVpnIpsecTunnelConfig struct {
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
	// The ID of the IPSec Gateway that the tunnel belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/vpn_ipsec_tunnel#gateway_id DataIonoscloudVpnIpsecTunnel#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/vpn_ipsec_tunnel#location DataIonoscloudVpnIpsecTunnel#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the IPSec Gateway Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/vpn_ipsec_tunnel#id DataIonoscloudVpnIpsecTunnel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The human readable name of your IPSec Gateway Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/vpn_ipsec_tunnel#name DataIonoscloudVpnIpsecTunnel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

