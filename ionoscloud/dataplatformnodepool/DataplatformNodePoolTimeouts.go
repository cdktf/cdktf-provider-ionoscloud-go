package dataplatformnodepool


type DataplatformNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_node_pool#create DataplatformNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_node_pool#default DataplatformNodePool#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_node_pool#delete DataplatformNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/dataplatform_node_pool#update DataplatformNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

