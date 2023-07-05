package dataionosclouddnsrecord


type DataIonoscloudDnsRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/data-sources/dns_record#create DataIonoscloudDnsRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/data-sources/dns_record#default DataIonoscloudDnsRecord#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/data-sources/dns_record#delete DataIonoscloudDnsRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/data-sources/dns_record#update DataIonoscloudDnsRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

