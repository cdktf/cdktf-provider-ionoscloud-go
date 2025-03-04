// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8scluster


type K8SClusterMaintenanceWindow struct {
	// Day of the week when maintenance is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/k8s_cluster#day_of_the_week K8SCluster#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// A clock time in the day when maintenance is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/k8s_cluster#time K8SCluster#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

