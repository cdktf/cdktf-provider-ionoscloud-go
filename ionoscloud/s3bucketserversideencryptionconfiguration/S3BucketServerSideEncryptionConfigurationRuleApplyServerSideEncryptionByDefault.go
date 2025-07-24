// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketserversideencryptionconfiguration


type S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// Server-side encryption algorithm to use. Valid values are 'AES256'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/s3_bucket_server_side_encryption_configuration#sse_algorithm S3BucketServerSideEncryptionConfiguration#sse_algorithm}
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

