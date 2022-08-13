// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudIpfailoverConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#datacenter_id DataIonoscloudIpfailover#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#id DataIonoscloudIpfailover#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#ip DataIonoscloudIpfailover#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#lan_id DataIonoscloudIpfailover#lan_id}.
	LanId *string `field:"optional" json:"lanId" yaml:"lanId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#nicuuid DataIonoscloudIpfailover#nicuuid}.
	Nicuuid *string `field:"optional" json:"nicuuid" yaml:"nicuuid"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/ipfailover#timeouts DataIonoscloudIpfailover#timeouts}
	Timeouts *DataIonoscloudIpfailoverTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

