// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset


type InmemorydbReplicasetCredentialsHashedPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/inmemorydb_replicaset#algorithm InmemorydbReplicaset#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.9/docs/resources/inmemorydb_replicaset#hash InmemorydbReplicaset#hash}.
	Hash *string `field:"required" json:"hash" yaml:"hash"`
}

