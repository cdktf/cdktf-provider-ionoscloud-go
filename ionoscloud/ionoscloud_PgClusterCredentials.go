// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type PgClusterCredentials struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster#password PgCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// the username for the initial postgres user. some system usernames are restricted (e.g. "postgres", "admin", "standby").
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster#username PgCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

