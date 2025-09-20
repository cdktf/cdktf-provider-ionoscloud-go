// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudpgdatabase


type DataIonoscloudPgDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/pg_database#create DataIonoscloudPgDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/pg_database#default DataIonoscloudPgDatabase#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/pg_database#delete DataIonoscloudPgDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/data-sources/pg_database#update DataIonoscloudPgDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

