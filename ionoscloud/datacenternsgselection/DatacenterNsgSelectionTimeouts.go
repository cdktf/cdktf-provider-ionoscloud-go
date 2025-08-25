// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacenternsgselection


type DatacenterNsgSelectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.14/docs/resources/datacenter_nsg_selection#create DatacenterNsgSelection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.14/docs/resources/datacenter_nsg_selection#default DatacenterNsgSelection#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.14/docs/resources/datacenter_nsg_selection#delete DatacenterNsgSelection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.14/docs/resources/datacenter_nsg_selection#update DatacenterNsgSelection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

