// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketConfig struct {
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
	// It must start and end with a letter or number and contain only lowercase alphanumeric characters, hyphens, periods and underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#name S3Bucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether all objects should be deleted from the bucket so that the bucket can be destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#force_destroy S3Bucket#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Whether object lock is enabled for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// The region of the bucket. Defaults to eu-central-3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#region S3Bucket#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A mapping of tags to assign to the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#tags S3Bucket#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.10/docs/resources/s3_bucket#timeouts S3Bucket#timeouts}
	Timeouts *S3BucketTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

