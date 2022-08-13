// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type NetworkloadbalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer#create Networkloadbalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer#default Networkloadbalancer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer#delete Networkloadbalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer#update Networkloadbalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

