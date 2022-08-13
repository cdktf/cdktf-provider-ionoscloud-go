// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type PgClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster#day_of_the_week PgCluster#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster#time PgCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

