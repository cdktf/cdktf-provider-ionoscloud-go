// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationIndexDocument struct {
	// A suffix that is appended to a request that is for a directory on the website endpoint (for example, if the suffix is index.html and you make a request to samplebucket/images/ the data that is returned will be for the object with the key name images/index.html) The suffix must not be empty and must not include a slash character. Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.8/docs/resources/s3_bucket_website_configuration#suffix S3BucketWebsiteConfiguration#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

