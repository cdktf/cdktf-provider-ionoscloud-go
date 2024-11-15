// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pgcluster


type PgClusterFromBackup struct {
	// The unique ID of the backup you want to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/pg_cluster#backup_id PgCluster#backup_id}
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp.
	//
	// If empty, the backup will be applied completely.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/pg_cluster#recovery_target_time PgCluster#recovery_target_time}
	RecoveryTargetTime *string `field:"optional" json:"recoveryTargetTime" yaml:"recoveryTargetTime"`
}

