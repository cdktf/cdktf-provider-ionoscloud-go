// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#name Group#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#access_activity_log Group#access_activity_log}.
	AccessActivityLog interface{} `field:"optional" json:"accessActivityLog" yaml:"accessActivityLog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#access_and_manage_certificates Group#access_and_manage_certificates}.
	AccessAndManageCertificates interface{} `field:"optional" json:"accessAndManageCertificates" yaml:"accessAndManageCertificates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#access_and_manage_monitoring Group#access_and_manage_monitoring}.
	AccessAndManageMonitoring interface{} `field:"optional" json:"accessAndManageMonitoring" yaml:"accessAndManageMonitoring"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_backup_unit Group#create_backup_unit}.
	CreateBackupUnit interface{} `field:"optional" json:"createBackupUnit" yaml:"createBackupUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_datacenter Group#create_datacenter}.
	CreateDatacenter interface{} `field:"optional" json:"createDatacenter" yaml:"createDatacenter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_flow_log Group#create_flow_log}.
	CreateFlowLog interface{} `field:"optional" json:"createFlowLog" yaml:"createFlowLog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_internet_access Group#create_internet_access}.
	CreateInternetAccess interface{} `field:"optional" json:"createInternetAccess" yaml:"createInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_k8s_cluster Group#create_k8s_cluster}.
	CreateK8SCluster interface{} `field:"optional" json:"createK8SCluster" yaml:"createK8SCluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_pcc Group#create_pcc}.
	CreatePcc interface{} `field:"optional" json:"createPcc" yaml:"createPcc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#create_snapshot Group#create_snapshot}.
	CreateSnapshot interface{} `field:"optional" json:"createSnapshot" yaml:"createSnapshot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#reserve_ip Group#reserve_ip}.
	ReserveIp interface{} `field:"optional" json:"reserveIp" yaml:"reserveIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#s3_privilege Group#s3_privilege}.
	S3Privilege interface{} `field:"optional" json:"s3Privilege" yaml:"s3Privilege"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#timeouts Group#timeouts}
	Timeouts *GroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#user_id Group#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/group#user_ids Group#user_ids}.
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

