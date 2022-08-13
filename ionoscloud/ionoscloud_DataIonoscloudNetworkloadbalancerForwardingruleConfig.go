// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudNetworkloadbalancerForwardingruleConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/networkloadbalancer_forwardingrule#datacenter_id DataIonoscloudNetworkloadbalancerForwardingrule#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/networkloadbalancer_forwardingrule#networkloadbalancer_id DataIonoscloudNetworkloadbalancerForwardingrule#networkloadbalancer_id}.
	NetworkloadbalancerId *string `field:"required" json:"networkloadbalancerId" yaml:"networkloadbalancerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/networkloadbalancer_forwardingrule#id DataIonoscloudNetworkloadbalancerForwardingrule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/networkloadbalancer_forwardingrule#name DataIonoscloudNetworkloadbalancerForwardingrule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/networkloadbalancer_forwardingrule#timeouts DataIonoscloudNetworkloadbalancerForwardingrule#timeouts}
	Timeouts *DataIonoscloudNetworkloadbalancerForwardingruleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

