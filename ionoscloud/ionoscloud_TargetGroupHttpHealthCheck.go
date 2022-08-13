// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type TargetGroupHttpHealthCheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#match_type TargetGroup#match_type}.
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// The response returned by the request, depending on the match type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#response TargetGroup#response}
	Response *string `field:"required" json:"response" yaml:"response"`
	// The method for the HTTP health check.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#method TargetGroup#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#negate TargetGroup#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The path (destination URL) for the HTTP health check request; the default is /.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#path TargetGroup#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/target_group#regex TargetGroup#regex}.
	Regex interface{} `field:"optional" json:"regex" yaml:"regex"`
}

