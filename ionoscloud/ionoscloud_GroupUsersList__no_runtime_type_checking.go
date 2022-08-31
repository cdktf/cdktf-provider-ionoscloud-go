//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroupUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroupUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroupUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroupUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

