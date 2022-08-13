// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type K8SClusterConfig struct {
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
	// The desired name for the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#name K8SCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Access to the K8s API server is restricted to these CIDRs.
	//
	// Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#api_subnet_allow_list K8SCluster#api_subnet_allow_list}
	ApiSubnetAllowList *[]*string `field:"optional" json:"apiSubnetAllowList" yaml:"apiSubnetAllowList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#id K8SCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The desired kubernetes version.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#k8s_version K8SCluster#k8s_version}
	K8SVersion *string `field:"optional" json:"k8SVersion" yaml:"k8SVersion"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#maintenance_window K8SCluster#maintenance_window}
	MaintenanceWindow *K8SClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// s3_buckets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#s3_buckets K8SCluster#s3_buckets}
	S3Buckets interface{} `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#timeouts K8SCluster#timeouts}
	Timeouts *K8SClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// List of versions that may be used for node pools under this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#viable_node_pool_versions K8SCluster#viable_node_pool_versions}
	ViableNodePoolVersions *[]*string `field:"optional" json:"viableNodePoolVersions" yaml:"viableNodePoolVersions"`
}

