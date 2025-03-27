// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudmongotemplate


type DataIonoscloudMongoTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/mongo_template#create DataIonoscloudMongoTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/mongo_template#default DataIonoscloudMongoTemplate#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/mongo_template#delete DataIonoscloudMongoTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/mongo_template#update DataIonoscloudMongoTemplate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

