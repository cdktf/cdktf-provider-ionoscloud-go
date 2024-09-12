// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationRoutingRule struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/s3_bucket_website_configuration#condition S3BucketWebsiteConfiguration#condition}
	Condition *S3BucketWebsiteConfigurationRoutingRuleCondition `field:"optional" json:"condition" yaml:"condition"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/s3_bucket_website_configuration#redirect S3BucketWebsiteConfiguration#redirect}
	Redirect *S3BucketWebsiteConfigurationRoutingRuleRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

