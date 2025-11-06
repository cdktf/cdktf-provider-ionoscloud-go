// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdndistribution

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnDistributionRoutingRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CdnDistributionRoutingRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnDistributionRoutingRulesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnDistributionRoutingRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnDistributionRoutingRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnDistributionRoutingRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnDistributionRoutingRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnDistributionRoutingRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

