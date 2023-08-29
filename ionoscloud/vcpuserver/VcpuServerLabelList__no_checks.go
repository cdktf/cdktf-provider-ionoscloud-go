// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vcpuserver

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VcpuServerLabelList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VcpuServerLabelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VcpuServerLabelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcpuServerLabelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VcpuServerLabelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VcpuServerLabelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVcpuServerLabelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

