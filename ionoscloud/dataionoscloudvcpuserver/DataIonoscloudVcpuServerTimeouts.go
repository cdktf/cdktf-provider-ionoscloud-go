// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudvcpuserver


type DataIonoscloudVcpuServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/data-sources/vcpu_server#create DataIonoscloudVcpuServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/data-sources/vcpu_server#default DataIonoscloudVcpuServer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/data-sources/vcpu_server#delete DataIonoscloudVcpuServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/data-sources/vcpu_server#update DataIonoscloudVcpuServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

