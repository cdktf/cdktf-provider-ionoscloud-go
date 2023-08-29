// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataionoscloudserver

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataIonoscloudServerLabelsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataIonoscloudServerLabelsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataIonoscloudServerLabelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataIonoscloudServerLabelsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataIonoscloudServerLabelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataIonoscloudServerLabelsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

