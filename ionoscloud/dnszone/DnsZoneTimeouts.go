// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnszone


type DnsZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/dns_zone#create DnsZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/dns_zone#default DnsZone#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/dns_zone#delete DnsZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.8/docs/resources/dns_zone#update DnsZone#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

