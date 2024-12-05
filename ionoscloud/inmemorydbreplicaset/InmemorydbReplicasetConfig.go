// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InmemorydbReplicasetConfig struct {
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
	// connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#connections InmemorydbReplicaset#connections}
	Connections *InmemorydbReplicasetConnections `field:"required" json:"connections" yaml:"connections"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#credentials InmemorydbReplicaset#credentials}
	Credentials *InmemorydbReplicasetCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The human readable name of your replica set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#display_name InmemorydbReplicaset#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The eviction policy for the replica set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#eviction_policy InmemorydbReplicaset#eviction_policy}
	EvictionPolicy *string `field:"required" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Specifies How and If data is persisted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#persistence_mode InmemorydbReplicaset#persistence_mode}
	PersistenceMode *string `field:"required" json:"persistenceMode" yaml:"persistenceMode"`
	// The total number of replicas in the replica set (one active and n-1 passive).
	//
	// In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#replicas InmemorydbReplicaset#replicas}
	Replicas *float64 `field:"required" json:"replicas" yaml:"replicas"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#resources InmemorydbReplicaset#resources}
	Resources *InmemorydbReplicasetResources `field:"required" json:"resources" yaml:"resources"`
	// The InMemoryDB version of your replica set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#version InmemorydbReplicaset#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#id InmemorydbReplicaset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of a snapshot to restore the replica set from.
	//
	// If set, the replica set will be created from the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#initial_snapshot_id InmemorydbReplicaset#initial_snapshot_id}
	InitialSnapshotId *string `field:"optional" json:"initialSnapshotId" yaml:"initialSnapshotId"`
	// The replica set location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#location InmemorydbReplicaset#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#maintenance_window InmemorydbReplicaset#maintenance_window}
	MaintenanceWindow *InmemorydbReplicasetMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/inmemorydb_replicaset#timeouts InmemorydbReplicaset#timeouts}
	Timeouts *InmemorydbReplicasetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

