package mongocluster


type MongoClusterCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/mongo_cluster#password MongoCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// the username for the initial mongoDB user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/mongo_cluster#username MongoCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}
