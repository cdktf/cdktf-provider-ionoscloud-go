// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleExpiration struct {
	// Specifies the date when the object expires. Required if 'days' is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/s3_bucket_lifecycle_configuration#date S3BucketLifecycleConfiguration#date}
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/s3_bucket_lifecycle_configuration#days S3BucketLifecycleConfiguration#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Indicates whether IONOS Object Storage Object Storage will remove a delete marker with no noncurrent versions.
	//
	// If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/s3_bucket_lifecycle_configuration#expired_object_delete_marker S3BucketLifecycleConfiguration#expired_object_delete_marker}
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
}

