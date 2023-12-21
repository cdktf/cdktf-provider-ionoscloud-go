// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongocluster


type MongoClusterBackup struct {
	// The location where the cluster backups will be stored.
	//
	// If not set, the backup is stored in the nearest location of the cluster. Examples: de, eu-sounth-2, eu-central-2
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/mongo_cluster#location MongoCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Number of hours in the past for which a point-in-time snapshot can be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/mongo_cluster#point_in_time_window_hours MongoCluster#point_in_time_window_hours}
	PointInTimeWindowHours *float64 `field:"optional" json:"pointInTimeWindowHours" yaml:"pointInTimeWindowHours"`
	// Number of hours between snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs/resources/mongo_cluster#snapshot_interval_hours MongoCluster#snapshot_interval_hours}
	SnapshotIntervalHours *float64 `field:"optional" json:"snapshotIntervalHours" yaml:"snapshotIntervalHours"`
}

