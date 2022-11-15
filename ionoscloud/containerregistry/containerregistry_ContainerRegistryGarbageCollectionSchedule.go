package containerregistry


type ContainerRegistryGarbageCollectionSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/container_registry#days ContainerRegistry#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/container_registry#time ContainerRegistry#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

