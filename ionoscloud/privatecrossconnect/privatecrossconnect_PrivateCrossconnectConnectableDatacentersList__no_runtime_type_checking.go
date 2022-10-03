//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package privatecrossconnect

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateCrossconnectConnectableDatacentersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateCrossconnectConnectableDatacentersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

