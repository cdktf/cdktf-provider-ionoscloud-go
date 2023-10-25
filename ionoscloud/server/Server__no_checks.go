// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Server) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutLabelParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutNicParameters(value *ServerNic) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutTimeoutsParameters(value *ServerTimeouts) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutVolumeParameters(value *ServerVolume) error {
	return nil
}

func validateServer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateServer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetAvailabilityZoneParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetBootCdromParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetBootImageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCoresParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCpuFamilyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetDatacenterIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetFirewallruleIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetImageNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetImagePasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetRamParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetSshKeyPathParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetSshKeysParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetTemplateUuidParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetVmStateParameters(val *string) error {
	return nil
}

func validateNewServerParameters(scope constructs.Construct, id *string, config *ServerConfig) error {
	return nil
}

