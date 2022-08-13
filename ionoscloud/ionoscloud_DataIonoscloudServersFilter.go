// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type DataIonoscloudServersFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/servers#name DataIonoscloudServers#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/servers#value DataIonoscloudServers#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

