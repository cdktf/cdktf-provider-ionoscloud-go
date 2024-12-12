// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketversioning


type S3BucketVersioningVersioningConfiguration struct {
	// The versioning status of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/s3_bucket_versioning#status S3BucketVersioning#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// The MFA delete status of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/s3_bucket_versioning#mfa_delete S3BucketVersioning#mfa_delete}
	MfaDelete *string `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

