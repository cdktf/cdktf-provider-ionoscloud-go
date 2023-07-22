package targetgroup


type TargetGroupHealthCheck struct {
	// The interval in milliseconds between consecutive health checks; default is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/target_group#check_interval TargetGroup#check_interval}
	CheckInterval *float64 `field:"optional" json:"checkInterval" yaml:"checkInterval"`
	// The maximum time in milliseconds to wait for a target to respond to a check.
	//
	// For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/target_group#check_timeout TargetGroup#check_timeout}
	CheckTimeout *float64 `field:"optional" json:"checkTimeout" yaml:"checkTimeout"`
	// The maximum number of attempts to reconnect to a target after a connection failure.
	//
	// Valid range is 0 to 65535, and default is three reconnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.4/docs/resources/target_group#retries TargetGroup#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
}

