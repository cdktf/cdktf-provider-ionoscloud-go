// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// Email address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#email User#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#first_name User#first_name}.
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#last_name User#last_name}.
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// Indicates if the user is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#active User#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Indicates if the user has administrative rights.
	//
	// Administrators do not need to be managed in groups, as they automatically have access to all resources associated with the contract.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#administrator User#administrator}
	Administrator interface{} `field:"optional" json:"administrator" yaml:"administrator"`
	// Indicates if secure (two-factor) authentication is forced for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#force_sec_auth User#force_sec_auth}
	ForceSecAuth interface{} `field:"optional" json:"forceSecAuth" yaml:"forceSecAuth"`
	// Ids of the groups that the user is a member of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#group_ids User#group_ids}
	GroupIds *[]*string `field:"optional" json:"groupIds" yaml:"groupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A password for the user.
	//
	// If you are using terraform 1.11 or higher, you can use `password_wo` instead of `password` to avoid storing the password in the state file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Write-only attribute. Password for the user. To modify, must change the password_wo_version attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#password_wo User#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Version of the password_wo attribute. Must be incremented to apply changes to the password_wo attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#password_wo_version User#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/user#timeouts User#timeouts}
	Timeouts *UserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

