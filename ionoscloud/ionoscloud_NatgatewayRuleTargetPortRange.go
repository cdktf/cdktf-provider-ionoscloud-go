// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type NatgatewayRuleTargetPortRange struct {
	// Target port range end associated with the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#end NatgatewayRule#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// Target port range start associated with the NAT gateway rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#start NatgatewayRule#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
}

