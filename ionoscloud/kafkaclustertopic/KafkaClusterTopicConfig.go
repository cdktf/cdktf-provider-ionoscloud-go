// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafkaclustertopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KafkaClusterTopicConfig struct {
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
	// The ID of the Kafka Cluster to which the topic belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#cluster_id KafkaClusterTopic#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The name of your Kafka Cluster Topic.
	//
	// Must be 63 characters or less and must begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#name KafkaClusterTopic#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The location of your Kafka Cluster Topic. Supported locations: de/fra, de/txl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#location KafkaClusterTopic#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The number of partitions of the topic.
	//
	// Partitions allow for parallel processing of messages. The partition count must be greater than or equal to the replication factor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#number_of_partitions KafkaClusterTopic#number_of_partitions}
	NumberOfPartitions *float64 `field:"optional" json:"numberOfPartitions" yaml:"numberOfPartitions"`
	// The number of replicas of the topic.
	//
	// The replication factor determines how many copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number of brokers in the Kafka Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#replication_factor KafkaClusterTopic#replication_factor}
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// This configuration controls the maximum time we will retain a log before we will discard old log segments to free up space.
	//
	// This represents an SLA on how soon consumers must read their data. If set to -1, no time limit is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#retention_time KafkaClusterTopic#retention_time}
	RetentionTime *float64 `field:"optional" json:"retentionTime" yaml:"retentionTime"`
	// This configuration controls the segment file size for the log.
	//
	// Retention and cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over retention.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#segment_bytes KafkaClusterTopic#segment_bytes}
	SegmentBytes *float64 `field:"optional" json:"segmentBytes" yaml:"segmentBytes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.3/docs/resources/kafka_cluster_topic#timeouts KafkaClusterTopic#timeouts}
	Timeouts *KafkaClusterTopicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

