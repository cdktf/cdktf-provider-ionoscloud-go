// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type ApplicationLoadbalancerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer#create ApplicationLoadbalancer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer#default ApplicationLoadbalancer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer#delete ApplicationLoadbalancer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer#update ApplicationLoadbalancer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

