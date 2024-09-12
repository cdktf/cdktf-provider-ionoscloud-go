// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MariadbClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#connections MariadbCluster#connections}
	Connections *MariadbClusterConnections `field:"required" json:"connections" yaml:"connections"`
	// The number of CPU cores per instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#cores MariadbCluster#cores}
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#credentials MariadbCluster#credentials}
	Credentials *MariadbClusterCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The friendly name of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#display_name MariadbCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The total number of instances in the cluster (one primary and n-1 secondary).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#instances MariadbCluster#instances}
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// The MariaDB version of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#mariadb_version MariadbCluster#mariadb_version}
	MariadbVersion *string `field:"required" json:"mariadbVersion" yaml:"mariadbVersion"`
	// The amount of memory per instance in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#ram MariadbCluster#ram}
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
	// The amount of storage per instance in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#storage_size MariadbCluster#storage_size}
	StorageSize *float64 `field:"required" json:"storageSize" yaml:"storageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#id MariadbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The cluster location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#location MariadbCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#maintenance_window MariadbCluster#maintenance_window}
	MaintenanceWindow *MariadbClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.4/docs/resources/mariadb_cluster#timeouts MariadbCluster#timeouts}
	Timeouts *MariadbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

