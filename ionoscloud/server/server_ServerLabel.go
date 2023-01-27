package server


type ServerLabel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#key Server#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#value Server#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

