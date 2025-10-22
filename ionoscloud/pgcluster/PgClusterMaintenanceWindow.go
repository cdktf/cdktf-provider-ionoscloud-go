// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pgcluster


type PgClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#day_of_the_week PgCluster#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#time PgCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

