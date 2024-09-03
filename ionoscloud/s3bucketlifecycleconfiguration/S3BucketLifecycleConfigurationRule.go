// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRule struct {
	// abort_incomplete_multipart_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#abort_incomplete_multipart_upload S3BucketLifecycleConfiguration#abort_incomplete_multipart_upload}
	AbortIncompleteMultipartUpload *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload `field:"required" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_version_expiration S3BucketLifecycleConfiguration#noncurrent_version_expiration}
	NoncurrentVersionExpiration *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration `field:"required" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Object key prefix identifying one or more objects to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#prefix S3BucketLifecycleConfiguration#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#status S3BucketLifecycleConfiguration#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#expiration S3BucketLifecycleConfiguration#expiration}
	Expiration *S3BucketLifecycleConfigurationRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Unique identifier for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/s3_bucket_lifecycle_configuration#id S3BucketLifecycleConfiguration#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

