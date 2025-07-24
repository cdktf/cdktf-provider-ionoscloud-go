// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#datacenter_id Server#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#name Server#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When set to true, allows the update of immutable fields by destroying and re-creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#allow_replace Server#allow_replace}
	AllowReplace interface{} `field:"optional" json:"allowReplace" yaml:"allowReplace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#availability_zone Server#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The associated boot drive, if any.
	//
	// Must be the UUID of a bootable CDROM image that you can retrieve using the image data source
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#boot_cdrom Server#boot_cdrom}
	BootCdrom *string `field:"optional" json:"bootCdrom" yaml:"bootCdrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#boot_image Server#boot_image}.
	BootImage *string `field:"optional" json:"bootImage" yaml:"bootImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#cores Server#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#cpu_family Server#cpu_family}.
	CpuFamily *string `field:"optional" json:"cpuFamily" yaml:"cpuFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#firewallrule_ids Server#firewallrule_ids}.
	FirewallruleIds *[]*string `field:"optional" json:"firewallruleIds" yaml:"firewallruleIds"`
	// The hostname of the resource.
	//
	// Allowed characters are a-z, 0-9 and - (minus). Hostname should not start with minus and should not be longer than 63 characters. If no value provided explicitly, it will be populated with the name of the server
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#hostname Server#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#id Server#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#image_name Server#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#image_password Server#image_password}.
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#label Server#label}
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#nic Server#nic}
	Nic *ServerNic `field:"optional" json:"nic" yaml:"nic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#ram Server#ram}.
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// The list of Security Group IDs for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#security_groups_ids Server#security_groups_ids}
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
	// Immutable List of absolute or relative paths to files containing public SSH key that will be injected into IonosCloud provided Linux images.
	//
	// Does not support `~` expansion to homedir in the given path. Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. This property is immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#ssh_key_path Server#ssh_key_path}
	SshKeyPath *[]*string `field:"optional" json:"sshKeyPath" yaml:"sshKeyPath"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key.
	//
	// This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#ssh_keys Server#ssh_keys}
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#template_uuid Server#template_uuid}.
	TemplateUuid *string `field:"optional" json:"templateUuid" yaml:"templateUuid"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#timeouts Server#timeouts}
	Timeouts *ServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// server usages: ENTERPRISE or CUBE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#type Server#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Sets the power state of the server.
	//
	// Possible values: `RUNNING`, `SHUTOFF` or `SUSPENDED`. SUSPENDED state is only valid for cube. SHUTOFF state is only valid for enterprise
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#vm_state Server#vm_state}
	VmState *string `field:"optional" json:"vmState" yaml:"vmState"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/server#volume Server#volume}
	Volume *ServerVolume `field:"optional" json:"volume" yaml:"volume"`
}

