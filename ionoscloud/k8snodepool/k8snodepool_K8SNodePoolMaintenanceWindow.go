package k8snodepool


type K8SNodePoolMaintenanceWindow struct {
	// Day of the week when maintenance is allowed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_node_pool#day_of_the_week K8SNodePool#day_of_the_week}
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// A clock time in the day when maintenance is allowed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_node_pool#time K8SNodePool#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

