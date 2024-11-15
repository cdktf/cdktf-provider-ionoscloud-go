// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketversioning

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketVersioningConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/s3_bucket_versioning#bucket S3BucketVersioning#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// versioning_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/s3_bucket_versioning#versioning_configuration S3BucketVersioning#versioning_configuration}
	VersioningConfiguration *S3BucketVersioningVersioningConfiguration `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
}

