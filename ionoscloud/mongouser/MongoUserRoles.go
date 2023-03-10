package mongouser


type MongoUserRoles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#database MongoUser#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// A list of mongodb user roles. Examples: read, readWrite, readAnyDatabase.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_user#role MongoUser#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

