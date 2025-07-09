// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8scluster


type K8SClusterS3Buckets struct {
	// Name of the Object Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/k8s_cluster#name K8SCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

