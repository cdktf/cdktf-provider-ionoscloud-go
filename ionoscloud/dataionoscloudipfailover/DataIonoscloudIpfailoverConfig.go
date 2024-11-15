// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudipfailover

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudIpfailoverConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/data-sources/ipfailover#datacenter_id DataIonoscloudIpfailover#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/data-sources/ipfailover#ip DataIonoscloudIpfailover#ip}.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/data-sources/ipfailover#lan_id DataIonoscloudIpfailover#lan_id}.
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/data-sources/ipfailover#timeouts DataIonoscloudIpfailover#timeouts}
	Timeouts *DataIonoscloudIpfailoverTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

