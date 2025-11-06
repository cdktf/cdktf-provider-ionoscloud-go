// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigatewayCustomDomainsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApigatewayCustomDomainsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigatewayCustomDomainsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigatewayCustomDomainsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApigatewayCustomDomainsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigatewayCustomDomainsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigatewayCustomDomainsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigatewayCustomDomainsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

