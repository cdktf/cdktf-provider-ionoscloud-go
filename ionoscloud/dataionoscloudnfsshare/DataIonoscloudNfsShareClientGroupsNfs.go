// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudnfsshare


type DataIonoscloudNfsShareClientGroupsNfs struct {
	// The squash mode for the export.
	//
	// The squash mode can be: none - No squash mode. no mapping, root-anonymous - Map root user to anonymous uid, all-anonymous - Map all users to anonymous uid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/data-sources/nfs_share#squash DataIonoscloudNfsShare#squash}
	Squash *string `field:"optional" json:"squash" yaml:"squash"`
}

