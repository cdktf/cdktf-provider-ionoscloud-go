// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketcorsconfiguration


type S3BucketCorsConfigurationCorsRule struct {
	// An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#allowed_methods S3BucketCorsConfiguration#allowed_methods}
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#allowed_origins S3BucketCorsConfiguration#allowed_origins}
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#allowed_headers S3BucketCorsConfiguration#allowed_headers}
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// One or more headers in the response that you want customers to be able to access from their applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#expose_headers S3BucketCorsConfiguration#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Container for the Contract Number of the owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#id S3BucketCorsConfiguration#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/s3_bucket_cors_configuration#max_age_seconds S3BucketCorsConfiguration#max_age_seconds}
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}

