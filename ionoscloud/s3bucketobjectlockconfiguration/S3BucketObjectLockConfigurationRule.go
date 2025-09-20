// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobjectlockconfiguration


type S3BucketObjectLockConfigurationRule struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.15/docs/resources/s3_bucket_object_lock_configuration#default_retention S3BucketObjectLockConfiguration#default_retention}
	DefaultRetention *S3BucketObjectLockConfigurationRuleDefaultRetention `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
}

