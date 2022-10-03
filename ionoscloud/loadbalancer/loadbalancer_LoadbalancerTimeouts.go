package loadbalancer


type LoadbalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/loadbalancer#create Loadbalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/loadbalancer#default Loadbalancer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/loadbalancer#delete Loadbalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/loadbalancer#update Loadbalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

