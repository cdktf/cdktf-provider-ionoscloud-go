// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancerforwardingrule


type ApplicationLoadbalancerForwardingruleHttpRulesConditions struct {
	// Type of the HTTP rule condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/application_loadbalancer_forwardingrule#type ApplicationLoadbalancerForwardingrule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Matching rule for the HTTP rule condition attribute;
	//
	// mandatory for HEADER, PATH, QUERY, METHOD, HOST, and COOKIE types; must be null when type is SOURCE_IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/application_loadbalancer_forwardingrule#condition ApplicationLoadbalancerForwardingrule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Must be null when type is PATH, METHOD, HOST, or SOURCE_IP.
	//
	// Key can only be set when type is COOKIES, HEADER, or QUERY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/application_loadbalancer_forwardingrule#key ApplicationLoadbalancerForwardingrule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies whether the condition is negated or not; the default is False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/application_loadbalancer_forwardingrule#negate ApplicationLoadbalancerForwardingrule#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// Mandatory for conditions CONTAINS, EQUALS, MATCHES, STARTS_WITH, ENDS_WITH;
	//
	// must be null when condition is EXISTS; should be a valid CIDR if provided and if type is SOURCE_IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.7/docs/resources/application_loadbalancer_forwardingrule#value ApplicationLoadbalancerForwardingrule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

