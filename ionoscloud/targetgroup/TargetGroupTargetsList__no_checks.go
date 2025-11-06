// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package targetgroup

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TargetGroupTargetsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TargetGroupTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TargetGroupTargetsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TargetGroupTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTargetGroupTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

