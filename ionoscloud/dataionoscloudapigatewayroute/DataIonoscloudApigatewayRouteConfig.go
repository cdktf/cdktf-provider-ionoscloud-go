// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudapigatewayroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudApigatewayRouteConfig struct {
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
	// The ID of the API Gateway that the route belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/data-sources/apigateway_route#gateway_id DataIonoscloudApigatewayRoute#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The ID (UUID) of the API Gateway Route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/data-sources/apigateway_route#id DataIonoscloudApigatewayRoute#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the API Gateway Route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/data-sources/apigateway_route#name DataIonoscloudApigatewayRoute#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether partial matching is allowed or not when using the name filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/data-sources/apigateway_route#partial_match DataIonoscloudApigatewayRoute#partial_match}
	PartialMatch interface{} `field:"optional" json:"partialMatch" yaml:"partialMatch"`
}

