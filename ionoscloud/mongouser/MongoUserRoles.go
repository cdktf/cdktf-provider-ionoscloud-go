package mongouser


type MongoUserRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/mongo_user#database MongoUser#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// A list of mongodb user roles. Examples: read, readWrite, readAnyDatabase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.3.6/docs/resources/mongo_user#role MongoUser#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

