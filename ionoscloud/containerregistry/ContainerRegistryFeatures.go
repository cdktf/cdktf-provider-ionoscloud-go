// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry


type ContainerRegistryFeatures struct {
	// Enables vulnerability scanning for images in the container registry. Note: this feature can incur additional charges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/container_registry#vulnerability_scanning ContainerRegistry#vulnerability_scanning}
	VulnerabilityScanning interface{} `field:"optional" json:"vulnerabilityScanning" yaml:"vulnerabilityScanning"`
}

