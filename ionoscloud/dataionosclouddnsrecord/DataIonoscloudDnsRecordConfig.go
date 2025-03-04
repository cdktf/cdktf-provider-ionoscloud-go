// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionosclouddnsrecord

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudDnsRecordConfig struct {
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
	// The UUID of an existing DNS Zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/data-sources/dns_record#zone_id DataIonoscloudDnsRecord#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The ID of your DNS Record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/data-sources/dns_record#id DataIonoscloudDnsRecord#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of your DNS Record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/data-sources/dns_record#name DataIonoscloudDnsRecord#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether partial matching is allowed or not when using name argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/data-sources/dns_record#partial_match DataIonoscloudDnsRecord#partial_match}
	PartialMatch interface{} `field:"optional" json:"partialMatch" yaml:"partialMatch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.3/docs/data-sources/dns_record#timeouts DataIonoscloudDnsRecord#timeouts}
	Timeouts *DataIonoscloudDnsRecordTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

