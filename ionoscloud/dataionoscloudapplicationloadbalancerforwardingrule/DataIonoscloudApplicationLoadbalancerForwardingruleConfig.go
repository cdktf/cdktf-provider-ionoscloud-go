// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudapplicationloadbalancerforwardingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudApplicationLoadbalancerForwardingruleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#application_loadbalancer_id DataIonoscloudApplicationLoadbalancerForwardingrule#application_loadbalancer_id}.
	ApplicationLoadbalancerId *string `field:"required" json:"applicationLoadbalancerId" yaml:"applicationLoadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#datacenter_id DataIonoscloudApplicationLoadbalancerForwardingrule#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#id DataIonoscloudApplicationLoadbalancerForwardingrule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the Application Load Balancer forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#name DataIonoscloudApplicationLoadbalancerForwardingrule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether partial matching is allowed or not when using name argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#partial_match DataIonoscloudApplicationLoadbalancerForwardingrule#partial_match}
	PartialMatch interface{} `field:"optional" json:"partialMatch" yaml:"partialMatch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/data-sources/application_loadbalancer_forwardingrule#timeouts DataIonoscloudApplicationLoadbalancerForwardingrule#timeouts}
	Timeouts *DataIonoscloudApplicationLoadbalancerForwardingruleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

