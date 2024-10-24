// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdndistribution


type CdnDistributionRoutingRulesUpstreamGeoRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/cdn_distribution#allow_list CdnDistribution#allow_list}.
	AllowList *[]*string `field:"optional" json:"allowList" yaml:"allowList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.0/docs/resources/cdn_distribution#block_list CdnDistribution#block_list}.
	BlockList *[]*string `field:"optional" json:"blockList" yaml:"blockList"`
}

