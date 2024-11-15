// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nfscluster


type NfsClusterNfs struct {
	// The minimum Network File Storage version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/nfs_cluster#min_version NfsCluster#min_version}
	MinVersion *string `field:"optional" json:"minVersion" yaml:"minVersion"`
}

