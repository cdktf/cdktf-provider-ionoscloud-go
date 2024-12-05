// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#location ContainerRegistry#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#name ContainerRegistry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subnet CIDRs that are allowed to connect to the registry.
	//
	// Specify 'a.b.c.d/32' for an individual IP address. __Note__: If this list is empty or not set, there are no restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#api_subnet_allow_list ContainerRegistry#api_subnet_allow_list}
	ApiSubnetAllowList *[]*string `field:"optional" json:"apiSubnetAllowList" yaml:"apiSubnetAllowList"`
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#features ContainerRegistry#features}
	Features *ContainerRegistryFeatures `field:"optional" json:"features" yaml:"features"`
	// garbage_collection_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#garbage_collection_schedule ContainerRegistry#garbage_collection_schedule}
	GarbageCollectionSchedule *ContainerRegistryGarbageCollectionSchedule `field:"optional" json:"garbageCollectionSchedule" yaml:"garbageCollectionSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#id ContainerRegistry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.5/docs/resources/container_registry#timeouts ContainerRegistry#timeouts}
	Timeouts *ContainerRegistryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

