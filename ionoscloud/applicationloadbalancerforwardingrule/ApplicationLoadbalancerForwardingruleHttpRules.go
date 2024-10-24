// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancerforwardingrule


type ApplicationLoadbalancerForwardingruleHttpRules struct {
	// The unique name of the Application Load Balancer HTTP rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#name ApplicationLoadbalancerForwardingrule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the HTTP rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#type ApplicationLoadbalancerForwardingrule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#conditions ApplicationLoadbalancerForwardingrule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Valid only for STATIC actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#content_type ApplicationLoadbalancerForwardingrule#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Default is false; valid only for REDIRECT actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#drop_query ApplicationLoadbalancerForwardingrule#drop_query}
	DropQuery interface{} `field:"optional" json:"dropQuery" yaml:"dropQuery"`
	// The location for redirecting; mandatory and valid only for REDIRECT actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#location ApplicationLoadbalancerForwardingrule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The response message of the request; mandatory for STATIC actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#response_message ApplicationLoadbalancerForwardingrule#response_message}
	ResponseMessage *string `field:"optional" json:"responseMessage" yaml:"responseMessage"`
	// Valid only for REDIRECT and STATIC actions.
	//
	// For REDIRECT actions, default is 301 and possible values are 301, 302, 303, 307, and 308. For STATIC actions, default is 503 and valid range is 200 to 599.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#status_code ApplicationLoadbalancerForwardingrule#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The ID of the target group; mandatory and only valid for FORWARD actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/application_loadbalancer_forwardingrule#target_group ApplicationLoadbalancerForwardingrule#target_group}
	TargetGroup *string `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

