package mongouser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongoUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#cluster_id MongoUser#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The user database to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#database MongoUser#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#password MongoUser#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user database uses for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#username MongoUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#id MongoUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#roles MongoUser#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#timeouts MongoUser#timeouts}
	Timeouts *MongoUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

