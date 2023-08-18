package dataplatformcluster


type DataplatformClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.7/docs/resources/dataplatform_cluster#day_of_the_week DataplatformCluster#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Time at which the maintenance should start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.7/docs/resources/dataplatform_cluster#time DataplatformCluster#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

