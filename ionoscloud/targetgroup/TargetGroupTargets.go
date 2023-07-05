package targetgroup


type TargetGroupTargets struct {
	// The IP of the balanced target VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/target_group#ip TargetGroup#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The port of the balanced target service; valid range is 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/target_group#port TargetGroup#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets.
	//
	// A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/target_group#weight TargetGroup#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Makes the target available only if it accepts periodic health check TCP connection attempts;
	//
	// when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/target_group#health_check_enabled TargetGroup#health_check_enabled}
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// Maintenance mode prevents the target from receiving balanced traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/target_group#maintenance_enabled TargetGroup#maintenance_enabled}
	MaintenanceEnabled interface{} `field:"optional" json:"maintenanceEnabled" yaml:"maintenanceEnabled"`
}

