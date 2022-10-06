package mongocluster


type MongoClusterCredentials struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_cluster#password MongoCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// the username for the initial mongoDB user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_cluster#username MongoCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

