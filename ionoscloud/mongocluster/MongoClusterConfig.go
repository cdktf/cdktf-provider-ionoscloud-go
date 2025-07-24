// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongocluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongoClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#connections MongoCluster#connections}
	Connections *MongoClusterConnections `field:"required" json:"connections" yaml:"connections"`
	// The name of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#display_name MongoCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The total number of instances in the cluster (one master and n-1 standbys).
	//
	// Example: 1, 3, 5, 7. For enterprise edition at least 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#instances MongoCluster#instances}
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// The physical location where the cluster will be created.
	//
	// This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit. Update forces cluster re-creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#location MongoCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The MongoDB version of your cluster. Downgrade is not possible and will throw an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#mongodb_version MongoCluster#mongodb_version}
	MongodbVersion *string `field:"required" json:"mongodbVersion" yaml:"mongodbVersion"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#backup MongoCluster#backup}
	Backup *MongoClusterBackup `field:"optional" json:"backup" yaml:"backup"`
	// bi_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#bi_connector MongoCluster#bi_connector}
	BiConnector *MongoClusterBiConnector `field:"optional" json:"biConnector" yaml:"biConnector"`
	// The number of CPU cores per instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#cores MongoCluster#cores}
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// The cluster edition. Must be one of: playground, business, enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#edition MongoCluster#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#id MongoCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#maintenance_window MongoCluster#maintenance_window}
	MaintenanceWindow *MongoClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The amount of memory per instance in megabytes. Multiple of 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#ram MongoCluster#ram}
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// The total number of shards in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#shards MongoCluster#shards}
	Shards *float64 `field:"optional" json:"shards" yaml:"shards"`
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#storage_size MongoCluster#storage_size}
	StorageSize *float64 `field:"optional" json:"storageSize" yaml:"storageSize"`
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#storage_type MongoCluster#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory.
	//
	// You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#template_id MongoCluster#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#timeouts MongoCluster#timeouts}
	Timeouts *MongoClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The cluster type, either `replicaset` or `sharded-cluster`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.12/docs/resources/mongo_cluster#type MongoCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

