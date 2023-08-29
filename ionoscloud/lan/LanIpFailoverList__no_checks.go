// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lan

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LanIpFailoverList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LanIpFailoverList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LanIpFailoverList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LanIpFailoverList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LanIpFailoverList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LanIpFailoverList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLanIpFailoverListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

