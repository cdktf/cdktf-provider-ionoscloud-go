// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudk8sclusters


type DataIonoscloudK8SClustersFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/k8s_clusters#name DataIonoscloudK8SClusters#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/data-sources/k8s_clusters#value DataIonoscloudK8SClusters#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

