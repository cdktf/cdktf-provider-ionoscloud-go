// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Bucket) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_S3Bucket) validatePutTimeoutsParameters(value *S3BucketTimeouts) error {
	return nil
}

func validateS3Bucket_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateS3Bucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3Bucket_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateS3Bucket_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_S3Bucket) validateSetRegionParameters(val *string) error {
	return nil
}

func validateNewS3BucketParameters(scope constructs.Construct, id *string, config *S3BucketConfig) error {
	return nil
}

