// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type K8SClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#create K8SCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#default K8SCluster#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#delete K8SCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#update K8SCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
