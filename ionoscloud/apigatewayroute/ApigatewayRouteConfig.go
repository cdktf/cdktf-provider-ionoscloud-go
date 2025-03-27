// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#gateway_id ApigatewayRoute#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The HTTP methods that the route should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#methods ApigatewayRoute#methods}
	Methods *[]*string `field:"required" json:"methods" yaml:"methods"`
	// The name of the API Gateway Route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#name ApigatewayRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The paths that the route should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#paths ApigatewayRoute#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// upstreams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#upstreams ApigatewayRoute#upstreams}
	Upstreams interface{} `field:"required" json:"upstreams" yaml:"upstreams"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#timeouts ApigatewayRoute#timeouts}
	Timeouts *ApigatewayRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// This field specifies the protocol used by the ingress to route traffic to the backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#type ApigatewayRoute#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// To enable websocket support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/apigateway_route#websocket ApigatewayRoute#websocket}
	Websocket interface{} `field:"optional" json:"websocket" yaml:"websocket"`
}

