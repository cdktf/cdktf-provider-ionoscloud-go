// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#name Group#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#access_activity_log Group#access_activity_log}.
	AccessActivityLog interface{} `field:"optional" json:"accessActivityLog" yaml:"accessActivityLog"`
	// Privilege for a group to access and manage certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#access_and_manage_certificates Group#access_and_manage_certificates}
	AccessAndManageCertificates interface{} `field:"optional" json:"accessAndManageCertificates" yaml:"accessAndManageCertificates"`
	// Privilege for a group to access and manage monitoring related functionality (access metrics, CRUD on alarms, alarm-actions etc) using Monotoring-as-a-Service (MaaS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#access_and_manage_monitoring Group#access_and_manage_monitoring}
	AccessAndManageMonitoring interface{} `field:"optional" json:"accessAndManageMonitoring" yaml:"accessAndManageMonitoring"`
	// Create backup unit privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_backup_unit Group#create_backup_unit}
	CreateBackupUnit interface{} `field:"optional" json:"createBackupUnit" yaml:"createBackupUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_datacenter Group#create_datacenter}.
	CreateDatacenter interface{} `field:"optional" json:"createDatacenter" yaml:"createDatacenter"`
	// Create Flow Logs privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_flow_log Group#create_flow_log}
	CreateFlowLog interface{} `field:"optional" json:"createFlowLog" yaml:"createFlowLog"`
	// Create internet access privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_internet_access Group#create_internet_access}
	CreateInternetAccess interface{} `field:"optional" json:"createInternetAccess" yaml:"createInternetAccess"`
	// Create Kubernetes cluster privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_k8s_cluster Group#create_k8s_cluster}
	CreateK8SCluster interface{} `field:"optional" json:"createK8SCluster" yaml:"createK8SCluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_pcc Group#create_pcc}.
	CreatePcc interface{} `field:"optional" json:"createPcc" yaml:"createPcc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#create_snapshot Group#create_snapshot}.
	CreateSnapshot interface{} `field:"optional" json:"createSnapshot" yaml:"createSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Privilege for a group to manage DBaaS related functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#manage_dbaas Group#manage_dbaas}
	ManageDbaas interface{} `field:"optional" json:"manageDbaas" yaml:"manageDbaas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#reserve_ip Group#reserve_ip}.
	ReserveIp interface{} `field:"optional" json:"reserveIp" yaml:"reserveIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#s3_privilege Group#s3_privilege}.
	S3Privilege interface{} `field:"optional" json:"s3Privilege" yaml:"s3Privilege"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#timeouts Group#timeouts}
	Timeouts *GroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#user_id Group#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.1/docs/resources/group#user_ids Group#user_ids}.
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

