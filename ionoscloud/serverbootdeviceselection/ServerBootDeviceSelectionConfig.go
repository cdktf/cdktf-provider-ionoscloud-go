// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverbootdeviceselection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerBootDeviceSelectionConfig struct {
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
	// ID of the Datacenter that holds the server for which the boot volume is selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/server_boot_device_selection#datacenter_id ServerBootDeviceSelection#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// ID of the Server for which the boot device will be selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/server_boot_device_selection#server_id ServerBootDeviceSelection#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// ID of the entity to set as primary boot device.
	//
	// Possible boot devices are CDROM Images and Volumes. If omitted, server will boot from PXE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/server_boot_device_selection#boot_device_id ServerBootDeviceSelection#boot_device_id}
	BootDeviceId *string `field:"optional" json:"bootDeviceId" yaml:"bootDeviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/server_boot_device_selection#id ServerBootDeviceSelection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/server_boot_device_selection#timeouts ServerBootDeviceSelection#timeouts}
	Timeouts *ServerBootDeviceSelectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

