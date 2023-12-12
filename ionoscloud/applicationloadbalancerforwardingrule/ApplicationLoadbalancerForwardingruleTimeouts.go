// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancerforwardingrule


type ApplicationLoadbalancerForwardingruleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/resources/application_loadbalancer_forwardingrule#create ApplicationLoadbalancerForwardingrule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/resources/application_loadbalancer_forwardingrule#default ApplicationLoadbalancerForwardingrule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/resources/application_loadbalancer_forwardingrule#delete ApplicationLoadbalancerForwardingrule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.11/docs/resources/application_loadbalancer_forwardingrule#update ApplicationLoadbalancerForwardingrule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

