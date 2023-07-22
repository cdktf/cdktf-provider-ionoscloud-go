package containerregistrytoken


type ContainerRegistryTokenTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/container_registry_token#create ContainerRegistryToken#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/container_registry_token#default ContainerRegistryToken#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/container_registry_token#delete ContainerRegistryToken#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/container_registry_token#update ContainerRegistryToken#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
