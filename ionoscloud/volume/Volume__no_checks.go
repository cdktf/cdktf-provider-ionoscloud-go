// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package volume

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Volume) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validatePutTimeoutsParameters(value *VolumeTimeouts) error {
	return nil
}

func validateVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVolume_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVolume_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetAvailabilityZoneParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetBackupUnitIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetBusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetDatacenterIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetDiskTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetImageNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetImagePasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetLicenceTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetServerIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetSshKeyPathParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetSshKeysParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Volume) validateSetUserDataParameters(val *string) error {
	return nil
}

func validateNewVolumeParameters(scope constructs.Construct, id *string, config *VolumeConfig) error {
	return nil
}

