package k8snodepool


type K8SNodePoolLans struct {
	// The LAN ID of an existing LAN at the related datacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/k8s_node_pool#id K8SNodePool#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/k8s_node_pool#dhcp K8SNodePool#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/k8s_node_pool#routes K8SNodePool#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

