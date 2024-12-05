// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudnfsshare


type DataIonoscloudNfsShareClientGroups struct {
	// Optional description for the clients groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/data-sources/nfs_share#description DataIonoscloudNfsShare#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A singular host allowed to connect to the share.
	//
	// The host can be specified as IP address and can be either IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/data-sources/nfs_share#hosts DataIonoscloudNfsShare#hosts}
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// The allowed host or network to which the export is being shared.
	//
	// The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/data-sources/nfs_share#ip_networks DataIonoscloudNfsShare#ip_networks}
	IpNetworks *[]*string `field:"optional" json:"ipNetworks" yaml:"ipNetworks"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/data-sources/nfs_share#nfs DataIonoscloudNfsShare#nfs}
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
}

