package dataionoscloudmongouser


type DataIonoscloudMongoUserRoles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/mongo_user#database DataIonoscloudMongoUser#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// A list of mongodb user roles. Examples: read, readWrite, readAnyDatabase.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/mongo_user#role DataIonoscloudMongoUser#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

