// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistry) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validatePutFeaturesParameters(value *ContainerRegistryFeatures) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validatePutGarbageCollectionScheduleParameters(value *ContainerRegistryGarbageCollectionSchedule) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistry) validatePutTimeoutsParameters(value *ContainerRegistryTimeouts) error {
	return nil
}

func validateContainerRegistry_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateContainerRegistry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateContainerRegistry_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateContainerRegistry_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetApiSubnetAllowListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistry) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewContainerRegistryParameters(scope constructs.Construct, id *string, config *ContainerRegistryConfig) error {
	return nil
}

