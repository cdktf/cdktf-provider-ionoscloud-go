// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnapshotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#datacenter_id Snapshot#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// A name of that resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#name Snapshot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#volume_id Snapshot#volume_id}.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#cpu_hot_plug Snapshot#cpu_hot_plug}.
	CpuHotPlug interface{} `field:"optional" json:"cpuHotPlug" yaml:"cpuHotPlug"`
	// Human readable description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#description Snapshot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#disc_virtio_hot_plug Snapshot#disc_virtio_hot_plug}.
	DiscVirtioHotPlug interface{} `field:"optional" json:"discVirtioHotPlug" yaml:"discVirtioHotPlug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#disc_virtio_hot_unplug Snapshot#disc_virtio_hot_unplug}.
	DiscVirtioHotUnplug interface{} `field:"optional" json:"discVirtioHotUnplug" yaml:"discVirtioHotUnplug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#id Snapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// OS type of this Snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#licence_type Snapshot#licence_type}
	LicenceType *string `field:"optional" json:"licenceType" yaml:"licenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#nic_hot_plug Snapshot#nic_hot_plug}.
	NicHotPlug interface{} `field:"optional" json:"nicHotPlug" yaml:"nicHotPlug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#nic_hot_unplug Snapshot#nic_hot_unplug}.
	NicHotUnplug interface{} `field:"optional" json:"nicHotUnplug" yaml:"nicHotUnplug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#ram_hot_plug Snapshot#ram_hot_plug}.
	RamHotPlug interface{} `field:"optional" json:"ramHotPlug" yaml:"ramHotPlug"`
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#sec_auth_protection Snapshot#sec_auth_protection}
	SecAuthProtection interface{} `field:"optional" json:"secAuthProtection" yaml:"secAuthProtection"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.6/docs/resources/snapshot#timeouts Snapshot#timeouts}
	Timeouts *SnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

