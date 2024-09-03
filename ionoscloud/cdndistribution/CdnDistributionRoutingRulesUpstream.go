// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdndistribution


type CdnDistributionRoutingRulesUpstream struct {
	// Enable or disable caching.
	//
	// If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/cdn_distribution#caching CdnDistribution#caching}
	Caching interface{} `field:"required" json:"caching" yaml:"caching"`
	// The upstream host that handles the requests if not already cached.
	//
	// This host will be protected by the WAF if the option is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/cdn_distribution#host CdnDistribution#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Rate limit class that will be applied to limit the number of incoming requests per IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/cdn_distribution#rate_limit_class CdnDistribution#rate_limit_class}
	RateLimitClass *string `field:"required" json:"rateLimitClass" yaml:"rateLimitClass"`
	// Enable or disable WAF to protect the upstream host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/cdn_distribution#waf CdnDistribution#waf}
	Waf interface{} `field:"required" json:"waf" yaml:"waf"`
	// geo_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.3/docs/resources/cdn_distribution#geo_restrictions CdnDistribution#geo_restrictions}
	GeoRestrictions *CdnDistributionRoutingRulesUpstreamGeoRestrictions `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
}

