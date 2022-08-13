// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type DataIonoscloudFirewallTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/firewall#create DataIonoscloudFirewall#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/firewall#default DataIonoscloudFirewall#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/firewall#delete DataIonoscloudFirewall#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/firewall#update DataIonoscloudFirewall#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

