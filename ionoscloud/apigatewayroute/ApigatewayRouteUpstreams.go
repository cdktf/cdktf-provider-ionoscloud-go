// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayroute


type ApigatewayRouteUpstreams struct {
	// The host of the upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/apigateway_route#host ApigatewayRoute#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The load balancer algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/apigateway_route#loadbalancer ApigatewayRoute#loadbalancer}
	Loadbalancer *string `field:"optional" json:"loadbalancer" yaml:"loadbalancer"`
	// The port of the upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/apigateway_route#port ApigatewayRoute#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The target URL of the upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/apigateway_route#scheme ApigatewayRoute#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Weight with which to split traffic to the upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/apigateway_route#weight ApigatewayRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

