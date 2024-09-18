// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketWebsiteConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/resources/s3_bucket_website_configuration#bucket S3BucketWebsiteConfiguration#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// error_document block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/resources/s3_bucket_website_configuration#error_document S3BucketWebsiteConfiguration#error_document}
	ErrorDocument *S3BucketWebsiteConfigurationErrorDocument `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// index_document block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/resources/s3_bucket_website_configuration#index_document S3BucketWebsiteConfiguration#index_document}
	IndexDocument *S3BucketWebsiteConfigurationIndexDocument `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// redirect_all_requests_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/resources/s3_bucket_website_configuration#redirect_all_requests_to S3BucketWebsiteConfiguration#redirect_all_requests_to}
	RedirectAllRequestsTo *S3BucketWebsiteConfigurationRedirectAllRequestsTo `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.5/docs/resources/s3_bucket_website_configuration#routing_rule S3BucketWebsiteConfiguration#routing_rule}
	RoutingRule interface{} `field:"optional" json:"routingRule" yaml:"routingRule"`
}

