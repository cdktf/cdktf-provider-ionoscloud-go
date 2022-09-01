//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_K8SNodePoolLansList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_K8SNodePoolLansList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewK8SNodePoolLansListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
