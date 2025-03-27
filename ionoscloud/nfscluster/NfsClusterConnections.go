// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nfscluster


type NfsClusterConnections struct {
	// The datacenter to connect your instance to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nfs_cluster#datacenter_id NfsCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The IP address and subnet for your instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nfs_cluster#ip_address NfsCluster#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The numeric LAN ID to connect your instance to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/nfs_cluster#lan NfsCluster#lan}
	Lan *string `field:"required" json:"lan" yaml:"lan"`
}

