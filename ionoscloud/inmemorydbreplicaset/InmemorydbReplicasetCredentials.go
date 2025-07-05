// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset


type InmemorydbReplicasetCredentials struct {
	// The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/inmemorydb_replicaset#username InmemorydbReplicaset#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// hashed_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/inmemorydb_replicaset#hashed_password InmemorydbReplicaset#hashed_password}
	HashedPassword *InmemorydbReplicasetCredentialsHashedPassword `field:"optional" json:"hashedPassword" yaml:"hashedPassword"`
	// The password for a InMemoryDB user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/inmemorydb_replicaset#plain_text_password InmemorydbReplicaset#plain_text_password}
	PlainTextPassword *string `field:"optional" json:"plainTextPassword" yaml:"plainTextPassword"`
}

