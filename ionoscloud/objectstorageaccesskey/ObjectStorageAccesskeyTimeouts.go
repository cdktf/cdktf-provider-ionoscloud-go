// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package objectstorageaccesskey


type ObjectStorageAccesskeyTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/object_storage_accesskey#create ObjectStorageAccesskey#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/object_storage_accesskey#delete ObjectStorageAccesskey#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/object_storage_accesskey#read ObjectStorageAccesskey#read}
	Read *string `field:"optional" json:"read" yaml:"read"`
}

