package k8scluster


type K8SClusterS3Buckets struct {
	// Name of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.3/docs/resources/k8s_cluster#name K8SCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

