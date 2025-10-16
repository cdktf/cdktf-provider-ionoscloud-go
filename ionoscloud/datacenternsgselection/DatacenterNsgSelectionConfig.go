// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacenternsgselection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatacenterNsgSelectionConfig struct {
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
	// ID of the Datacenter to which the NSG will be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/datacenter_nsg_selection#datacenter_id DatacenterNsgSelection#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// ID of the NSG which will be attached to the datacenter.
	//
	// If an empty string is specified and a NSG was attached previously, it will be unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/datacenter_nsg_selection#nsg_id DatacenterNsgSelection#nsg_id}
	NsgId *string `field:"required" json:"nsgId" yaml:"nsgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/datacenter_nsg_selection#id DatacenterNsgSelection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/datacenter_nsg_selection#timeouts DatacenterNsgSelection#timeouts}
	Timeouts *DatacenterNsgSelectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

