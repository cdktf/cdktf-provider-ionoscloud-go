// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certificate

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Certificate) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutTimeoutsParameters(value *CertificateTimeouts) error {
	return nil
}

func validateCertificate_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCertificate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCertificate_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCertificate_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCertificateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCertificateChainParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetPrivateKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewCertificateParameters(scope constructs.Construct, id *string, config *CertificateConfig) error {
	return nil
}

