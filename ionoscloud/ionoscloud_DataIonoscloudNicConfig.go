// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudNicConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#datacenter_id DataIonoscloudNic#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#server_id DataIonoscloudNic#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#dhcp DataIonoscloudNic#dhcp}.
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#firewall_active DataIonoscloudNic#firewall_active}.
	FirewallActive interface{} `field:"optional" json:"firewallActive" yaml:"firewallActive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#firewall_type DataIonoscloudNic#firewall_type}.
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#id DataIonoscloudNic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#ips DataIonoscloudNic#ips}.
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#lan DataIonoscloudNic#lan}.
	Lan *float64 `field:"optional" json:"lan" yaml:"lan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#name DataIonoscloudNic#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/nic#timeouts DataIonoscloudNic#timeouts}
	Timeouts *DataIonoscloudNicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
