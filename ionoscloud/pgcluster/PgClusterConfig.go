// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pgcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PgClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The number of CPU cores per replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#cores PgCluster#cores}
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#credentials PgCluster#credentials}
	Credentials *PgClusterCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The friendly name of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#display_name PgCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The total number of instances in the cluster (one master and n-1 standbys).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#instances PgCluster#instances}
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// The physical location where the cluster will be created.
	//
	// This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#location PgCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The PostgreSQL version of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#postgres_version PgCluster#postgres_version}
	PostgresVersion *string `field:"required" json:"postgresVersion" yaml:"postgresVersion"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#ram PgCluster#ram}
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#storage_size PgCluster#storage_size}
	StorageSize *float64 `field:"required" json:"storageSize" yaml:"storageSize"`
	// The storage type used in your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#storage_type PgCluster#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Represents different modes of replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#synchronization_mode PgCluster#synchronization_mode}
	SynchronizationMode *string `field:"required" json:"synchronizationMode" yaml:"synchronizationMode"`
	// When set to true, allows the update of immutable fields by destroying and re-creating the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#allow_replace PgCluster#allow_replace}
	AllowReplace interface{} `field:"optional" json:"allowReplace" yaml:"allowReplace"`
	// The Object Storage location where the backups will be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#backup_location PgCluster#backup_location}
	BackupLocation *string `field:"optional" json:"backupLocation" yaml:"backupLocation"`
	// connection_pooler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#connection_pooler PgCluster#connection_pooler}
	ConnectionPooler *PgClusterConnectionPooler `field:"optional" json:"connectionPooler" yaml:"connectionPooler"`
	// connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#connections PgCluster#connections}
	Connections *PgClusterConnections `field:"optional" json:"connections" yaml:"connections"`
	// from_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#from_backup PgCluster#from_backup}
	FromBackup *PgClusterFromBackup `field:"optional" json:"fromBackup" yaml:"fromBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#id PgCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#maintenance_window PgCluster#maintenance_window}
	MaintenanceWindow *PgClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/pg_cluster#timeouts PgCluster#timeouts}
	Timeouts *PgClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

