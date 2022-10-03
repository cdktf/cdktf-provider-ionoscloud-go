package user


type UserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/user#create User#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/user#default User#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/user#delete User#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/user#update User#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

