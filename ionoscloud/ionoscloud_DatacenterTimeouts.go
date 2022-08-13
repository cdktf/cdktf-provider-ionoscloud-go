// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type DatacenterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/datacenter#create Datacenter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/datacenter#default Datacenter#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/datacenter#delete Datacenter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/datacenter#update Datacenter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

