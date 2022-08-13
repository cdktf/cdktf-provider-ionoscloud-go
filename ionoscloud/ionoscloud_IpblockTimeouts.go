// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type IpblockTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/ipblock#create Ipblock#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/ipblock#default Ipblock#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/ipblock#delete Ipblock#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/ipblock#update Ipblock#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

