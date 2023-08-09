package dataionoscloudfirewall


type DataIonoscloudFirewallTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/data-sources/firewall#create DataIonoscloudFirewall#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/data-sources/firewall#default DataIonoscloudFirewall#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/data-sources/firewall#delete DataIonoscloudFirewall#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/data-sources/firewall#update DataIonoscloudFirewall#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

