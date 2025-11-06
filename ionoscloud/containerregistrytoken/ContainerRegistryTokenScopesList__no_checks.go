// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerregistrytoken

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryTokenScopesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryTokenScopesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryTokenScopesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenScopesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenScopesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenScopesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryTokenScopesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryTokenScopesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

