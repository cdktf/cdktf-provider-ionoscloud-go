//go:build no_runtime_type_checking

package mongouser

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MongoUserRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MongoUserRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MongoUserRolesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MongoUserRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MongoUserRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MongoUserRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMongoUserRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

