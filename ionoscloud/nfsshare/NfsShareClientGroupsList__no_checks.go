// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nfsshare

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NfsShareClientGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NfsShareClientGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NfsShareClientGroupsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NfsShareClientGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NfsShareClientGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NfsShareClientGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NfsShareClientGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNfsShareClientGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

