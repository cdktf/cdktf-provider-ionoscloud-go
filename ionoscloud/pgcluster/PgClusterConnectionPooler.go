// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pgcluster


type PgClusterConnectionPooler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/pg_cluster#enabled PgCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Represents different modes of connection pooling for the connection pooler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/pg_cluster#pool_mode PgCluster#pool_mode}
	PoolMode *string `field:"required" json:"poolMode" yaml:"poolMode"`
}

