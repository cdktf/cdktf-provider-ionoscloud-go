// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type ApplicationLoadbalancerForwardingruleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer_forwardingrule#create ApplicationLoadbalancerForwardingrule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer_forwardingrule#default ApplicationLoadbalancerForwardingrule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer_forwardingrule#delete ApplicationLoadbalancerForwardingrule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/application_loadbalancer_forwardingrule#update ApplicationLoadbalancerForwardingrule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
