// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targetgroup


type TargetGroupHttpHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#match_type TargetGroup#match_type}.
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// The response returned by the request, depending on the match type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#response TargetGroup#response}
	Response *string `field:"required" json:"response" yaml:"response"`
	// The method for the HTTP health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#method TargetGroup#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#negate TargetGroup#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The path (destination URL) for the HTTP health check request; the default is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#path TargetGroup#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/target_group#regex TargetGroup#regex}.
	Regex interface{} `field:"optional" json:"regex" yaml:"regex"`
}

