// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type PrivateCrossconnectTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/private_crossconnect#create PrivateCrossconnect#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/private_crossconnect#default PrivateCrossconnect#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/private_crossconnect#delete PrivateCrossconnect#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/private_crossconnect#update PrivateCrossconnect#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

