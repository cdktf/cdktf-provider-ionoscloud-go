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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#connections MongoCluster#connections}
	Connections *MongoClusterConnections `field:"required" json:"connections" yaml:"connections"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#credentials MongoCluster#credentials}
	Credentials *MongoClusterCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The name of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#display_name MongoCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The total number of instances in the cluster (one master and n-1 standbys).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#instances MongoCluster#instances}
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// The physical location where the cluster will be created.
	//
	// This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#location MongoCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The MongoDB version of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#mongodb_version MongoCluster#mongodb_version}
	MongodbVersion *string `field:"required" json:"mongodbVersion" yaml:"mongodbVersion"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory.
	//
	// You cannot downgrade to a smaller template or minor edition (e.g. from business to playground).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#template_id MongoCluster#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#id MongoCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#maintenance_window MongoCluster#maintenance_window}
	MaintenanceWindow *MongoClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/mongo_cluster#timeouts MongoCluster#timeouts}
	Timeouts *MongoClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

