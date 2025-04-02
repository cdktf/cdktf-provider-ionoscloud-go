// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nfsshare


type NfsShareClientGroups struct {
	// A singular host allowed to connect to the share.
	//
	// The host can be specified as IP address and can be either IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/nfs_share#hosts NfsShare#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// The allowed host or network to which the export is being shared.
	//
	// The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/nfs_share#ip_networks NfsShare#ip_networks}
	IpNetworks *[]*string `field:"required" json:"ipNetworks" yaml:"ipNetworks"`
	// Optional description for the clients groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/nfs_share#description NfsShare#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/nfs_share#nfs NfsShare#nfs}
	Nfs *NfsShareClientGroupsNfs `field:"optional" json:"nfs" yaml:"nfs"`
}

