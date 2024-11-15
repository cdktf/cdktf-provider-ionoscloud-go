// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbcluster


type MariadbClusterMaintenanceWindow struct {
	// The name of the week day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/mariadb_cluster#day_of_the_week MariadbCluster#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Start of the maintenance window in UTC time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/mariadb_cluster#time MariadbCluster#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

