// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudpguser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudPgUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/data-sources/pg_user#cluster_id DataIonoscloudPgUser#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/data-sources/pg_user#username DataIonoscloudPgUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/data-sources/pg_user#timeouts DataIonoscloudPgUser#timeouts}
	Timeouts *DataIonoscloudPgUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

