// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration struct {
	// Specifies the number of days an object is noncurrent before IONOS Object Storage can perform the associated action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_days S3BucketLifecycleConfiguration#noncurrent_days}
	NoncurrentDays *float64 `field:"optional" json:"noncurrentDays" yaml:"noncurrentDays"`
}

