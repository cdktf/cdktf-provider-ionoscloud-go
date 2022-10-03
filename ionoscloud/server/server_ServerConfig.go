package server

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#datacenter_id Server#datacenter_id}.
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#name Server#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// nic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#nic Server#nic}
	Nic *ServerNic `field:"required" json:"nic" yaml:"nic"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#volume Server#volume}
	Volume *ServerVolume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#availability_zone Server#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#boot_cdrom Server#boot_cdrom}.
	BootCdrom *string `field:"optional" json:"bootCdrom" yaml:"bootCdrom"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#boot_image Server#boot_image}.
	BootImage *string `field:"optional" json:"bootImage" yaml:"bootImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#cores Server#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#cpu_family Server#cpu_family}.
	CpuFamily *string `field:"optional" json:"cpuFamily" yaml:"cpuFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#id Server#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#image_name Server#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#image_password Server#image_password}.
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#ram Server#ram}.
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#ssh_key_path Server#ssh_key_path}.
	SshKeyPath *[]*string `field:"optional" json:"sshKeyPath" yaml:"sshKeyPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#template_uuid Server#template_uuid}.
	TemplateUuid *string `field:"optional" json:"templateUuid" yaml:"templateUuid"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#timeouts Server#timeouts}
	Timeouts *ServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/server#type Server#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

