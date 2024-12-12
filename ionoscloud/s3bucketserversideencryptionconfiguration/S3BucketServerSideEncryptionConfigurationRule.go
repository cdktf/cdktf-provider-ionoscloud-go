// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketserversideencryptionconfiguration


type S3BucketServerSideEncryptionConfigurationRule struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/s3_bucket_server_side_encryption_configuration#apply_server_side_encryption_by_default S3BucketServerSideEncryptionConfiguration#apply_server_side_encryption_by_default}
	ApplyServerSideEncryptionByDefault *S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `field:"required" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
}

