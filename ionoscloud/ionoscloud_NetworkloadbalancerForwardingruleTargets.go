// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type NetworkloadbalancerForwardingruleTargets struct {
	// IP of a balanced target VM.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer_forwardingrule#ip NetworkloadbalancerForwardingrule#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Port of the balanced target service. (range: 1 to 65535).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer_forwardingrule#port NetworkloadbalancerForwardingrule#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Weight parameter is used to adjust the target VM's weight relative to other target VMs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer_forwardingrule#weight NetworkloadbalancerForwardingrule#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/networkloadbalancer_forwardingrule#health_check NetworkloadbalancerForwardingrule#health_check}
	HealthCheck *NetworkloadbalancerForwardingruleTargetsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
}

