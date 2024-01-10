// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loggingpipeline

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoggingPipelineLogDestinationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoggingPipelineLogDestinationsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoggingPipelineLogDestinationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogDestinationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogDestinationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogDestinationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoggingPipelineLogDestinationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoggingPipelineLogDestinationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

