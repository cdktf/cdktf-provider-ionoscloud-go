// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupConfig struct {
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
	// Unique identifier for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#datacenter_id AutoscalingGroup#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The maximum value for the number of replicas on a VM Auto Scaling Group.
	//
	// Must be >= 0 and <= 200. Will be enforced for both automatic and manual changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#max_replica_count AutoscalingGroup#max_replica_count}
	MaxReplicaCount *float64 `field:"required" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// The minimum value for the number of replicas on a VM Auto Scaling Group.
	//
	// Must be >= 0 and <= 200. Will be enforced for both automatic and manual changes
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#min_replica_count AutoscalingGroup#min_replica_count}
	MinReplicaCount *float64 `field:"required" json:"minReplicaCount" yaml:"minReplicaCount"`
	// User-defined name for the Autoscaling Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#name AutoscalingGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#policy AutoscalingGroup#policy}
	Policy *AutoscalingGroupPolicy `field:"required" json:"policy" yaml:"policy"`
	// replica_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#replica_configuration AutoscalingGroup#replica_configuration}
	ReplicaConfiguration *AutoscalingGroupReplicaConfiguration `field:"required" json:"replicaConfiguration" yaml:"replicaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#id AutoscalingGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/autoscaling_group#timeouts AutoscalingGroup#timeouts}
	Timeouts *AutoscalingGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

