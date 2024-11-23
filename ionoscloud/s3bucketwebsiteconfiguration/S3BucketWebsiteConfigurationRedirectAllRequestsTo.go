// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationRedirectAllRequestsTo struct {
	// The host name to use in the redirect request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/s3_bucket_website_configuration#host_name S3BucketWebsiteConfiguration#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/s3_bucket_website_configuration#protocol S3BucketWebsiteConfiguration#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

