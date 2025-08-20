// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplatformnodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplatformNodePoolConfig struct {
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
	// The UUID of an existing Dataplatform cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#cluster_id DataplatformNodePool#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The name of your node pool.
	//
	// Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#name DataplatformNodePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of nodes that make up the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#node_count DataplatformNodePool#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#annotations DataplatformNodePool#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#auto_scaling DataplatformNodePool#auto_scaling}
	AutoScaling *DataplatformNodePoolAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#availability_zone DataplatformNodePool#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of CPU cores per node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#cores_count DataplatformNodePool#cores_count}
	CoresCount *float64 `field:"optional" json:"coresCount" yaml:"coresCount"`
	// A valid CPU family name or `AUTO` if the platform shall choose the best fitting option.
	//
	// Available CPU architectures can be retrieved from the datacenter resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#cpu_family DataplatformNodePool#cpu_family}
	CpuFamily *string `field:"optional" json:"cpuFamily" yaml:"cpuFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#id DataplatformNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#labels DataplatformNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#maintenance_window DataplatformNodePool#maintenance_window}
	MaintenanceWindow *DataplatformNodePoolMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The RAM size for one node in MB.
	//
	// Must be set in multiples of 1024 MB, with a minimum size is of 2048 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#ram_size DataplatformNodePool#ram_size}
	RamSize *float64 `field:"optional" json:"ramSize" yaml:"ramSize"`
	// The size of the volume in GB. The size must be greater than 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#storage_size DataplatformNodePool#storage_size}
	StorageSize *float64 `field:"optional" json:"storageSize" yaml:"storageSize"`
	// The type of hardware for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#storage_type DataplatformNodePool#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/dataplatform_node_pool#timeouts DataplatformNodePool#timeouts}
	Timeouts *DataplatformNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

