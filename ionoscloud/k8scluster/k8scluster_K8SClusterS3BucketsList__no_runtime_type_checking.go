//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8scluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_K8SClusterS3BucketsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_K8SClusterS3BucketsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_K8SClusterS3BucketsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_K8SClusterS3BucketsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_K8SClusterS3BucketsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_K8SClusterS3BucketsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewK8SClusterS3BucketsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

