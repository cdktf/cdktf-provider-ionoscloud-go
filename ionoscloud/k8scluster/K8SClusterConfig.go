// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8scluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type K8SClusterConfig struct {
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
	// The desired name for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#name K8SCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When set to true, allows the update of immutable fields by destroying and re-creating the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#allow_replace K8SCluster#allow_replace}
	AllowReplace interface{} `field:"optional" json:"allowReplace" yaml:"allowReplace"`
	// Access to the K8s API server is restricted to these CIDRs.
	//
	// Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#api_subnet_allow_list K8SCluster#api_subnet_allow_list}
	ApiSubnetAllowList *[]*string `field:"optional" json:"apiSubnetAllowList" yaml:"apiSubnetAllowList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#id K8SCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The desired Kubernetes Version.
	//
	// For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#k8s_version K8SCluster#k8s_version}
	K8SVersion *string `field:"optional" json:"k8SVersion" yaml:"k8SVersion"`
	// This attribute is mandatory if the cluster is private.
	//
	// The location must be enabled for your contract, or you must have a data center at that location. This attribute is immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#location K8SCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#maintenance_window K8SCluster#maintenance_window}
	MaintenanceWindow *K8SClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The NAT gateway IP of the cluster if the cluster is private.
	//
	// This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#nat_gateway_ip K8SCluster#nat_gateway_ip}
	NatGatewayIp *string `field:"optional" json:"natGatewayIp" yaml:"natGatewayIp"`
	// The node subnet of the cluster, if the cluster is private.
	//
	// This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#node_subnet K8SCluster#node_subnet}
	NodeSubnet *string `field:"optional" json:"nodeSubnet" yaml:"nodeSubnet"`
	// The indicator if the cluster is public or private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#public K8SCluster#public}
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// s3_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#s3_buckets K8SCluster#s3_buckets}
	S3Buckets interface{} `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.19/docs/resources/k8s_cluster#timeouts K8SCluster#timeouts}
	Timeouts *K8SClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

