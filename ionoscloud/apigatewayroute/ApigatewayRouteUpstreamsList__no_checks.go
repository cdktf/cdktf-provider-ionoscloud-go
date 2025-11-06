// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigatewayroute

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigatewayRouteUpstreamsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApigatewayRouteUpstreamsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigatewayRouteUpstreamsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRouteUpstreamsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRouteUpstreamsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRouteUpstreamsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRouteUpstreamsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigatewayRouteUpstreamsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

