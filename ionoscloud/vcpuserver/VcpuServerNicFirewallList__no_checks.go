// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vcpuserver

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VcpuServerNicFirewallList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VcpuServerNicFirewallList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VcpuServerNicFirewallList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcpuServerNicFirewallList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VcpuServerNicFirewallList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VcpuServerNicFirewallList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVcpuServerNicFirewallListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

