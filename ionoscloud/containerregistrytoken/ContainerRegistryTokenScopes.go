package containerregistrytoken


type ContainerRegistryTokenScopes struct {
	// Example: ["pull", "push", "delete"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/resources/container_registry_token#actions ContainerRegistryToken#actions}
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/resources/container_registry_token#name ContainerRegistryToken#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.5/docs/resources/container_registry_token#type ContainerRegistryToken#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

