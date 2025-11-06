// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafkaclustertopic


type KafkaClusterTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/kafka_cluster_topic#create KafkaClusterTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/kafka_cluster_topic#default KafkaClusterTopic#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/kafka_cluster_topic#delete KafkaClusterTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.19/docs/resources/kafka_cluster_topic#update KafkaClusterTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

