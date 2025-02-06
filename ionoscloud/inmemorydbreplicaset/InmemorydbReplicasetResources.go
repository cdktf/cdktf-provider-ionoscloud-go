// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset


type InmemorydbReplicasetResources struct {
	// The number of CPU cores per instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.1/docs/resources/inmemorydb_replicaset#cores InmemorydbReplicaset#cores}
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// The amount of memory per instance in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.1/docs/resources/inmemorydb_replicaset#ram InmemorydbReplicaset#ram}
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
}

