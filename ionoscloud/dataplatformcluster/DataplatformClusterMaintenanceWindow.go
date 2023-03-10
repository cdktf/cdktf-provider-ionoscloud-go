package dataplatformcluster


type DataplatformClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_cluster#day_of_the_week DataplatformCluster#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Time at which the maintenance should start.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_cluster#time DataplatformCluster#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

