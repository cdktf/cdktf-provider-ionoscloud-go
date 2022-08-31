//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TargetGroupTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TargetGroupTargetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTargetGroupTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

