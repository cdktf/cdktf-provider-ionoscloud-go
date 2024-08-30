// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationloadbalancerforwardingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationLoadbalancerForwardingruleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#application_loadbalancer_id ApplicationLoadbalancerForwardingrule#application_loadbalancer_id}.
	ApplicationLoadbalancerId *string `field:"required" json:"applicationLoadbalancerId" yaml:"applicationLoadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#datacenter_id ApplicationLoadbalancerForwardingrule#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Listening (inbound) IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#listener_ip ApplicationLoadbalancerForwardingrule#listener_ip}
	ListenerIp *string `field:"required" json:"listenerIp" yaml:"listenerIp"`
	// Listening (inbound) port number; valid range is 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#listener_port ApplicationLoadbalancerForwardingrule#listener_port}
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
	// The name of the Application Load Balancer forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#name ApplicationLoadbalancerForwardingrule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Balancing protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#protocol ApplicationLoadbalancerForwardingrule#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The maximum time in milliseconds to wait for the client to acknowledge or send data;
	//
	// default is 50,000 (50 seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#client_timeout ApplicationLoadbalancerForwardingrule#client_timeout}
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// http_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#http_rules ApplicationLoadbalancerForwardingrule#http_rules}
	HttpRules interface{} `field:"optional" json:"httpRules" yaml:"httpRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#id ApplicationLoadbalancerForwardingrule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Array of items in the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#server_certificates ApplicationLoadbalancerForwardingrule#server_certificates}
	ServerCertificates *[]*string `field:"optional" json:"serverCertificates" yaml:"serverCertificates"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs/resources/application_loadbalancer_forwardingrule#timeouts ApplicationLoadbalancerForwardingrule#timeouts}
	Timeouts *ApplicationLoadbalancerForwardingruleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

