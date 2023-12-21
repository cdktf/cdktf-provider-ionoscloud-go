// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type IonoscloudProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#alias IonoscloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// To be set only for reseller accounts. Allows to run terraform on a contract number under a reseller account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#contract_number IonoscloudProvider#contract_number}
	ContractNumber *string `field:"optional" json:"contractNumber" yaml:"contractNumber"`
	// IonosCloud REST API URL.
	//
	// Usually not necessary to be set, SDKs know internally how to route requests to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#endpoint IonoscloudProvider#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// IonosCloud password for API operations. If token is provided, token is preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#password IonoscloudProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#retries IonoscloudProvider#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// IonosCloud bearer token for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#token IonoscloudProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// IonosCloud username for API operations. If token is provided, token is preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.12/docs#username IonoscloudProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

