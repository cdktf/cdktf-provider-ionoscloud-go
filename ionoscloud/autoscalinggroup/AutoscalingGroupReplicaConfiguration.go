// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfiguration struct {
	// The zone where the VMs are created using this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#availability_zone AutoscalingGroup#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The total number of cores for the VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#cores AutoscalingGroup#cores}
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// The amount of memory for the VMs in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#ram AutoscalingGroup#ram}
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
	// The zone where the VMs are created using this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#cpu_family AutoscalingGroup#cpu_family}
	CpuFamily *string `field:"optional" json:"cpuFamily" yaml:"cpuFamily"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#nic AutoscalingGroup#nic}
	Nic interface{} `field:"optional" json:"nic" yaml:"nic"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/autoscaling_group#volume AutoscalingGroup#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

