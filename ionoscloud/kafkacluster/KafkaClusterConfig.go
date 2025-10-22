// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafkacluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KafkaClusterConfig struct {
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
	// connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#connections KafkaCluster#connections}
	Connections *KafkaClusterConnections `field:"required" json:"connections" yaml:"connections"`
	// The name of your Kafka Cluster.
	//
	// Must be 63 characters or less and must begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#name KafkaCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size of your Kafka Cluster.
	//
	// The size of the Kafka Cluster is given in T-shirt sizes. Valid values are: XS, S
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#size KafkaCluster#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// The desired Kafka Version. Supported versions: 3.8.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#version KafkaCluster#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The location of your Kafka Cluster. Supported locations: de/fra, de/fra/2, de/txl, fr/par, es/vit, gb/lhr, gb/bhx, us/las, us/mci, us/ewr.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#location KafkaCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#timeouts KafkaCluster#timeouts}
	Timeouts *KafkaClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

