// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type K8SClusterS3Buckets struct {
	// Name of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/k8s_cluster#name K8SCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

