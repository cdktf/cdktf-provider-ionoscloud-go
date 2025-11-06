// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loggingpipeline

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoggingPipelineLogList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoggingPipelineLogList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoggingPipelineLogList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoggingPipelineLogListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

