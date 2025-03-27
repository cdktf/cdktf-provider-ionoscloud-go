// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTokenConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#name ContainerRegistryToken#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#registry_id ContainerRegistryToken#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#expiry_date ContainerRegistryToken#expiry_date}.
	ExpiryDate *string `field:"optional" json:"expiryDate" yaml:"expiryDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#id ContainerRegistryToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Saves password to file. Only works on create. Takes as argument a file name, or a file path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#save_password_to_file ContainerRegistryToken#save_password_to_file}
	SavePasswordToFile *string `field:"optional" json:"savePasswordToFile" yaml:"savePasswordToFile"`
	// scopes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#scopes ContainerRegistryToken#scopes}
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// Can be one of enabled, disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#status ContainerRegistryToken#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/container_registry_token#timeouts ContainerRegistryToken#timeouts}
	Timeouts *ContainerRegistryTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

