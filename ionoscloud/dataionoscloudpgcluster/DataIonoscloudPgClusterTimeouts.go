// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudpgcluster


type DataIonoscloudPgClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/pg_cluster#create DataIonoscloudPgCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/pg_cluster#default DataIonoscloudPgCluster#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/pg_cluster#delete DataIonoscloudPgCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/pg_cluster#update DataIonoscloudPgCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

