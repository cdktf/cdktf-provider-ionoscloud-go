// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobjectlockconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketObjectLockConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/s3_bucket_object_lock_configuration#bucket S3BucketObjectLockConfiguration#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specifies whether Object Lock is enabled for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/s3_bucket_object_lock_configuration#object_lock_enabled S3BucketObjectLockConfiguration#object_lock_enabled}
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.6/docs/resources/s3_bucket_object_lock_configuration#rule S3BucketObjectLockConfiguration#rule}
	Rule *S3BucketObjectLockConfigurationRule `field:"optional" json:"rule" yaml:"rule"`
}

