// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset


type InmemorydbReplicasetMaintenanceWindow struct {
	// The name of the week day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/inmemorydb_replicaset#day_of_the_week InmemorydbReplicaset#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Start of the maintenance window in UTC time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.1/docs/resources/inmemorydb_replicaset#time InmemorydbReplicaset#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

