// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package autoscalinggroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingGroupReplicaConfigurationNicListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

