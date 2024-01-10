// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryStorageUsageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryStorageUsageList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryStorageUsageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryStorageUsageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryStorageUsageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryStorageUsageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryStorageUsageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

