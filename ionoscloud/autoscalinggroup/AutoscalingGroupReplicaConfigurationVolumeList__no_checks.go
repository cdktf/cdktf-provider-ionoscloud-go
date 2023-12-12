// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package autoscalinggroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingGroupReplicaConfigurationVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

