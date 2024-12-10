// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CubeServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#datacenter_id CubeServer#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#name CubeServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#nic CubeServer#nic}
	Nic *CubeServerNic `field:"required" json:"nic" yaml:"nic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#template_uuid CubeServer#template_uuid}.
	TemplateUuid *string `field:"required" json:"templateUuid" yaml:"templateUuid"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#volume CubeServer#volume}
	Volume *CubeServerVolume `field:"required" json:"volume" yaml:"volume"`
	// When set to true, allows the update of immutable fields by destroying and re-creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#allow_replace CubeServer#allow_replace}
	AllowReplace interface{} `field:"optional" json:"allowReplace" yaml:"allowReplace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#availability_zone CubeServer#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#boot_cdrom CubeServer#boot_cdrom}.
	BootCdrom *string `field:"optional" json:"bootCdrom" yaml:"bootCdrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#boot_image CubeServer#boot_image}.
	BootImage *string `field:"optional" json:"bootImage" yaml:"bootImage"`
	// The hostname of the resource.
	//
	// Allowed characters are a-z, 0-9 and - (minus). Hostname should not start with minus and should not be longer than 63 characters. If no value provided explicitly, it will be populated with the name of the server
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#hostname CubeServer#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#id CubeServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#image_name CubeServer#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#image_password CubeServer#image_password}.
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// The list of Security Group IDs for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#security_groups_ids CubeServer#security_groups_ids}
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#ssh_key_path CubeServer#ssh_key_path}.
	SshKeyPath *[]*string `field:"optional" json:"sshKeyPath" yaml:"sshKeyPath"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#timeouts CubeServer#timeouts}
	Timeouts *CubeServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Sets the power state of the cube server. Possible values: `RUNNING` or `SUSPENDED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/cube_server#vm_state CubeServer#vm_state}
	VmState *string `field:"optional" json:"vmState" yaml:"vmState"`
}

