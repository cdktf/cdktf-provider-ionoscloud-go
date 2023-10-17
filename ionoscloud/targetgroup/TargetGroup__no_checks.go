// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package targetgroup

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TargetGroup) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validatePutHealthCheckParameters(value *TargetGroupHealthCheck) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validatePutHttpHealthCheckParameters(value *TargetGroupHttpHealthCheck) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validatePutTargetsParameters(value interface{}) error {
	return nil
}

func (t *jsiiProxy_TargetGroup) validatePutTimeoutsParameters(value *TargetGroupTimeouts) error {
	return nil
}

func validateTargetGroup_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTargetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTargetGroup_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTargetGroup_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetAlgorithmParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetProtocolParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroup) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewTargetGroupParameters(scope constructs.Construct, id *string, config *TargetGroupConfig) error {
	return nil
}

