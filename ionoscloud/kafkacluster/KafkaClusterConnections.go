// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafkacluster


type KafkaClusterConnections struct {
	// The broker addresses of the Kafka Cluster. Can be empty, but must be present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#broker_addresses KafkaCluster#broker_addresses}
	BrokerAddresses *[]*string `field:"required" json:"brokerAddresses" yaml:"brokerAddresses"`
	// The datacenter to connect your Kafka Cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#datacenter_id KafkaCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The numeric LAN ID to connect your Kafka Cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/kafka_cluster#lan_id KafkaCluster#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
}

