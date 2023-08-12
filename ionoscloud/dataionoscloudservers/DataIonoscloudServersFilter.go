package dataionoscloudservers


type DataIonoscloudServersFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/data-sources/servers#name DataIonoscloudServers#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/data-sources/servers#value DataIonoscloudServers#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

