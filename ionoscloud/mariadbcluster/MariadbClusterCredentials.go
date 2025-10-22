// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbcluster


type MariadbClusterCredentials struct {
	// The password for a MariaDB user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/mariadb_cluster#password MariadbCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username for the initial MariaDB user. Some system usernames are restricted (e.g 'mariadb', 'admin', 'standby').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/mariadb_cluster#username MariadbCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

