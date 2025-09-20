// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerLabel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/server#key Server#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/server#value Server#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

