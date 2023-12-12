// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupunit

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupUnit) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateImportFromParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateMoveToIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (b *jsiiProxy_BackupUnit) validatePutTimeoutsParameters(value *BackupUnitTimeouts) error {
	return nil
}

func validateBackupUnit_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateBackupUnit_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBackupUnit_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBackupUnit_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetEmailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetPasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupUnit) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewBackupUnitParameters(scope constructs.Construct, id *string, config *BackupUnitConfig) error {
	return nil
}

