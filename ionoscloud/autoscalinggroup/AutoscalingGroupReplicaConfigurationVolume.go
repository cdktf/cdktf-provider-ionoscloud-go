// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupReplicaConfigurationVolume struct {
	// Determines whether the volume will be used as a boot volume.
	//
	// Set to NONE, the volume will not be used as boot volume.
	// Set to PRIMARY, the volume will be used as boot volume and set to AUTO will delegate the decision to the provisioning engine to decide whether to use the volume as boot volume.
	// Notice that exactly one volume can be set to PRIMARY or all of them set to AUTO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#boot_order AutoscalingGroup#boot_order}
	BootOrder *string `field:"required" json:"bootOrder" yaml:"bootOrder"`
	// Name for this replica volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// User-defined size for this replica volume in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#size AutoscalingGroup#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Storage Type for this replica volume. Possible values: SSD, HDD, SSD_STANDARD or SSD_PREMIUM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#type AutoscalingGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The uuid of the Backup Unit that user has access to.
	//
	// The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#backup_unit_id AutoscalingGroup#backup_unit_id}
	BackupUnitId *string `field:"optional" json:"backupUnitId" yaml:"backupUnitId"`
	// The bus type of the volume. Default setting is 'VIRTIO'. The bus type 'IDE' is also supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#bus AutoscalingGroup#bus}
	Bus *string `field:"optional" json:"bus" yaml:"bus"`
	// The image installed on the disk.
	//
	// Currently, only the UUID of the image is supported. Note that either 'image' or 'imageAlias' must be specified, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#image AutoscalingGroup#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The image installed on the volume.
	//
	// Must be an 'imageAlias' as specified via the images API. Note that one of 'image' or 'imageAlias' must be set, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#image_alias AutoscalingGroup#image_alias}
	ImageAlias *string `field:"optional" json:"imageAlias" yaml:"imageAlias"`
	// Image password for this replica volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#image_password AutoscalingGroup#image_password}
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#ssh_keys AutoscalingGroup#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// User-data (Cloud Init) for this replica volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/resources/autoscaling_group#user_data AutoscalingGroup#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

