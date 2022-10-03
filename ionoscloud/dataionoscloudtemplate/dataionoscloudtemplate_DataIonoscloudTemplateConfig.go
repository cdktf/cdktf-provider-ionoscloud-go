package dataionoscloudtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/template#cores DataIonoscloudTemplate#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/template#name DataIonoscloudTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/template#ram DataIonoscloudTemplate#ram}.
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/template#storage_size DataIonoscloudTemplate#storage_size}.
	StorageSize *float64 `field:"optional" json:"storageSize" yaml:"storageSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/template#timeouts DataIonoscloudTemplate#timeouts}
	Timeouts *DataIonoscloudTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

