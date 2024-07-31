// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type IonoscloudProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#alias IonoscloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// To be set only for reseller accounts. Allows to run terraform on a contract number under a reseller account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#contract_number IonoscloudProvider#contract_number}
	ContractNumber *string `field:"optional" json:"contractNumber" yaml:"contractNumber"`
	// IonosCloud REST API URL.
	//
	// Usually not necessary to be set, SDKs know internally how to route requests to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#endpoint IonoscloudProvider#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// IonosCloud password for API operations. If token is provided, token is preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#password IonoscloudProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#retries IonoscloudProvider#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Access key for IONOS S3 operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#s3_access_key IonoscloudProvider#s3_access_key}
	S3AccessKey *string `field:"optional" json:"s3AccessKey" yaml:"s3AccessKey"`
	// Region for IONOS S3 operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#s3_region IonoscloudProvider#s3_region}
	S3Region *string `field:"optional" json:"s3Region" yaml:"s3Region"`
	// Secret key for IONOS S3 operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#s3_secret_key IonoscloudProvider#s3_secret_key}
	S3SecretKey *string `field:"optional" json:"s3SecretKey" yaml:"s3SecretKey"`
	// IonosCloud bearer token for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#token IonoscloudProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// IonosCloud username for API operations. If token is provided, token is preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs#username IonoscloudProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

