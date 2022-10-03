package natgatewayrule


type NatgatewayRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#create NatgatewayRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#default NatgatewayRule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#delete NatgatewayRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/natgateway_rule#update NatgatewayRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

