// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ipblock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpblockIpConsumersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IpblockIpConsumersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IpblockIpConsumersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IpblockIpConsumersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IpblockIpConsumersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IpblockIpConsumersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIpblockIpConsumersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

