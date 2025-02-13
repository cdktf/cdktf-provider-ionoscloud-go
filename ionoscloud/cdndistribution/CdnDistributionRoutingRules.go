// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdndistribution


type CdnDistributionRoutingRules struct {
	// The prefix of the routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/cdn_distribution#prefix CdnDistribution#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The scheme of the routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/cdn_distribution#scheme CdnDistribution#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// upstream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.2/docs/resources/cdn_distribution#upstream CdnDistribution#upstream}
	Upstream *CdnDistributionRoutingRulesUpstream `field:"required" json:"upstream" yaml:"upstream"`
}

