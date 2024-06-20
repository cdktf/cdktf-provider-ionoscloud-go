// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targetgroup


type TargetGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.18/docs/resources/target_group#create TargetGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.18/docs/resources/target_group#default TargetGroup#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.18/docs/resources/target_group#delete TargetGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.18/docs/resources/target_group#update TargetGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

