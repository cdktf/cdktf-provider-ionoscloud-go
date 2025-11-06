// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongouser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongoUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#cluster_id MongoUser#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#password MongoUser#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#username MongoUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#id MongoUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#roles MongoUser#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/mongo_user#timeouts MongoUser#timeouts}
	Timeouts *MongoUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

