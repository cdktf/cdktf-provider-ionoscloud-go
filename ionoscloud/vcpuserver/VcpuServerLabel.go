package vcpuserver


type VcpuServerLabel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/vcpu_server#key VcpuServer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/vcpu_server#value VcpuServer#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

