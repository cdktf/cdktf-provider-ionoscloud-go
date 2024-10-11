// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8snodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type K8SNodePoolConfig struct {
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
	// The compute availability zone in which the nodes should exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#availability_zone K8SNodePool#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// CPU cores count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#cores_count K8SNodePool#cores_count}
	CoresCount *float64 `field:"required" json:"coresCount" yaml:"coresCount"`
	// CPU Family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#cpu_family K8SNodePool#cpu_family}
	CpuFamily *string `field:"required" json:"cpuFamily" yaml:"cpuFamily"`
	// The UUID of the VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#datacenter_id K8SNodePool#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The UUID of an existing kubernetes cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#k8s_cluster_id K8SNodePool#k8s_cluster_id}
	K8SClusterId *string `field:"required" json:"k8SClusterId" yaml:"k8SClusterId"`
	// The desired Kubernetes Version.
	//
	// For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#k8s_version K8SNodePool#k8s_version}
	K8SVersion *string `field:"required" json:"k8SVersion" yaml:"k8SVersion"`
	// The desired name for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#name K8SNodePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of nodes in this node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#node_count K8SNodePool#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// The amount of RAM in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#ram_size K8SNodePool#ram_size}
	RamSize *float64 `field:"required" json:"ramSize" yaml:"ramSize"`
	// The total allocated storage capacity of a node in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#storage_size K8SNodePool#storage_size}
	StorageSize *float64 `field:"required" json:"storageSize" yaml:"storageSize"`
	// Storage type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#storage_type K8SNodePool#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// When set to true, allows the update of immutable fields by destroying and re-creating the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#allow_replace K8SNodePool#allow_replace}
	AllowReplace interface{} `field:"optional" json:"allowReplace" yaml:"allowReplace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#annotations K8SNodePool#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#auto_scaling K8SNodePool#auto_scaling}
	AutoScaling *K8SNodePoolAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#id K8SNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#labels K8SNodePool#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// lans block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#lans K8SNodePool#lans}
	Lans interface{} `field:"optional" json:"lans" yaml:"lans"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#maintenance_window K8SNodePool#maintenance_window}
	MaintenanceWindow *K8SNodePoolMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// A list of fixed IPs. Cannot be set on private clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#public_ips K8SNodePool#public_ips}
	PublicIps *[]*string `field:"optional" json:"publicIps" yaml:"publicIps"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.8/docs/resources/k8s_node_pool#timeouts K8SNodePool#timeouts}
	Timeouts *K8SNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

