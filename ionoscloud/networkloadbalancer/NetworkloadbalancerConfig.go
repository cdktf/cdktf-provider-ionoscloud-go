package networkloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkloadbalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#datacenter_id Networkloadbalancer#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Id of the listening LAN. (inbound).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#listener_lan Networkloadbalancer#listener_lan}
	ListenerLan *float64 `field:"required" json:"listenerLan" yaml:"listenerLan"`
	// A name of that Network Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#name Networkloadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Id of the balanced private target LAN. (outbound).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#target_lan Networkloadbalancer#target_lan}
	TargetLan *float64 `field:"required" json:"targetLan" yaml:"targetLan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#id Networkloadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Collection of IP addresses of the Network Load Balancer.
	//
	// (inbound and outbound) IP of the listenerLan must be a customer reserved IP for the public load balancer and private IP for the private load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#ips Networkloadbalancer#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// Collection of private IP addresses with subnet mask of the Network Load Balancer.
	//
	// IPs must contain valid subnet mask. If user will not provide any IP then the system will generate one IP with /24 subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#lb_private_ips Networkloadbalancer#lb_private_ips}
	LbPrivateIps *[]*string `field:"optional" json:"lbPrivateIps" yaml:"lbPrivateIps"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/networkloadbalancer#timeouts Networkloadbalancer#timeouts}
	Timeouts *NetworkloadbalancerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

