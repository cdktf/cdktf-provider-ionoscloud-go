// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TargetGroupConfig struct {
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
	// Balancing algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#algorithm TargetGroup#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The name of the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#name TargetGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Balancing protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#protocol TargetGroup#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#protocol_version TargetGroup#protocol_version}
	ProtocolVersion *string `field:"required" json:"protocolVersion" yaml:"protocolVersion"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#health_check TargetGroup#health_check}
	HealthCheck *TargetGroupHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// http_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#http_health_check TargetGroup#http_health_check}
	HttpHealthCheck *TargetGroupHttpHealthCheck `field:"optional" json:"httpHealthCheck" yaml:"httpHealthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#id TargetGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#targets TargetGroup#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/target_group#timeouts TargetGroup#timeouts}
	Timeouts *TargetGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

