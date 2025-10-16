// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleFilter struct {
	// Object key prefix identifying one or more objects to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/s3_bucket_lifecycle_configuration#prefix S3BucketLifecycleConfiguration#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

