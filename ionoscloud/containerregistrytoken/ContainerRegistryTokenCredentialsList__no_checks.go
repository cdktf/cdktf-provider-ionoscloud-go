//go:build no_runtime_type_checking

package containerregistrytoken

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryTokenCredentialsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryTokenCredentialsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenCredentialsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenCredentialsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenCredentialsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryTokenCredentialsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
