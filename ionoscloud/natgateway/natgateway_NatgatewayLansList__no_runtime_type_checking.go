//go:build no_runtime_type_checking

package natgateway

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatgatewayLansList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NatgatewayLansList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NatgatewayLansList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NatgatewayLansList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NatgatewayLansList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NatgatewayLansList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNatgatewayLansListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

