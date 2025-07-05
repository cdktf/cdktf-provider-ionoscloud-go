// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VcpuServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#cores VcpuServer#cores}.
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#datacenter_id VcpuServer#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#name VcpuServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#ram VcpuServer#ram}.
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#volume VcpuServer#volume}
	Volume *VcpuServerVolume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#availability_zone VcpuServer#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The associated boot drive, if any.
	//
	// Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#boot_cdrom VcpuServer#boot_cdrom}
	BootCdrom *string `field:"optional" json:"bootCdrom" yaml:"bootCdrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#boot_image VcpuServer#boot_image}.
	BootImage *string `field:"optional" json:"bootImage" yaml:"bootImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#firewallrule_ids VcpuServer#firewallrule_ids}.
	FirewallruleIds *[]*string `field:"optional" json:"firewallruleIds" yaml:"firewallruleIds"`
	// The hostname of the resource.
	//
	// Allowed characters are a-z, 0-9 and - (minus). Hostname should not start with minus and should not be longer than 63 characters. If no value provided explicitly, it will be populated with the name of the server
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#hostname VcpuServer#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#id VcpuServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#image_name VcpuServer#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#image_password VcpuServer#image_password}.
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#label VcpuServer#label}
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#nic VcpuServer#nic}
	Nic *VcpuServerNic `field:"optional" json:"nic" yaml:"nic"`
	// The list of Security Group IDs for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#security_groups_ids VcpuServer#security_groups_ids}
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key.
	//
	// This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#ssh_keys VcpuServer#ssh_keys}
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#timeouts VcpuServer#timeouts}
	Timeouts *VcpuServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Sets the power state of the vcpu server. Possible values: `RUNNING` or `SHUTOFF`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/vcpu_server#vm_state VcpuServer#vm_state}
	VmState *string `field:"optional" json:"vmState" yaml:"vmState"`
}

