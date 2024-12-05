// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pgcluster


type PgClusterConnections struct {
	// The IP and subnet for the database.
	//
	// Note the following unavailable IP ranges:
	//           10.233.64.0/18
	//           10.233.0.0/18
	//           10.233.114.0/24
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/pg_cluster#cidr PgCluster#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The datacenter to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/pg_cluster#datacenter_id PgCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The LAN to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/pg_cluster#lan_id PgCluster#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
}

