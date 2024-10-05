// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudnetworkloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudNetworkloadbalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/networkloadbalancer#datacenter_id DataIonoscloudNetworkloadbalancer#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/networkloadbalancer#id DataIonoscloudNetworkloadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/networkloadbalancer#name DataIonoscloudNetworkloadbalancer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/data-sources/networkloadbalancer#timeouts DataIonoscloudNetworkloadbalancer#timeouts}
	Timeouts *DataIonoscloudNetworkloadbalancerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

