// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayConfig struct {
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
	// The name of the API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/apigateway#name Apigateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// custom_domains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/apigateway#custom_domains Apigateway#custom_domains}
	CustomDomains interface{} `field:"optional" json:"customDomains" yaml:"customDomains"`
	// Enable or disable logging. NOTE: Central Logging must be enabled through the Logging API to enable this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/apigateway#logs Apigateway#logs}
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// Enable or disable metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/apigateway#metrics Apigateway#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.18/docs/resources/apigateway#timeouts Apigateway#timeouts}
	Timeouts *ApigatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

