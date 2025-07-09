// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbcluster


type MariadbClusterBackup struct {
	// The IONOS Object Storage location where the backups will be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/mariadb_cluster#location MariadbCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

