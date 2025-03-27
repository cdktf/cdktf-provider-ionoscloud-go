// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbcluster


type MariadbClusterConnections struct {
	// The IP and subnet for your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/mariadb_cluster#cidr MariadbCluster#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The datacenter to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/mariadb_cluster#datacenter_id MariadbCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The numeric LAN ID to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/mariadb_cluster#lan_id MariadbCluster#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
}

