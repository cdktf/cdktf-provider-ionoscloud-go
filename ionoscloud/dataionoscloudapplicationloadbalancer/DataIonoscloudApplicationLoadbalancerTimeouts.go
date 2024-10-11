// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudapplicationloadbalancer


type DataIonoscloudApplicationLoadbalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/data-sources/application_loadbalancer#create DataIonoscloudApplicationLoadbalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/data-sources/application_loadbalancer#default DataIonoscloudApplicationLoadbalancer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/data-sources/application_loadbalancer#delete DataIonoscloudApplicationLoadbalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/data-sources/application_loadbalancer#update DataIonoscloudApplicationLoadbalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

