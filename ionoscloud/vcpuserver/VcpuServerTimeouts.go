// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver


type VcpuServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/vcpu_server#create VcpuServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/vcpu_server#default VcpuServer#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/vcpu_server#delete VcpuServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/vcpu_server#update VcpuServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

