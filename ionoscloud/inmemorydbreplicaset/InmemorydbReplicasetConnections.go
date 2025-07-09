// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset


type InmemorydbReplicasetConnections struct {
	// The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/inmemorydb_replicaset#cidr InmemorydbReplicaset#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The datacenter to connect your instance to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/inmemorydb_replicaset#datacenter_id InmemorydbReplicaset#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The numeric LAN ID to connect your instance to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/inmemorydb_replicaset#lan_id InmemorydbReplicaset#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
}

