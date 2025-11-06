// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package k8snodepool

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_K8SNodePoolLansRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_K8SNodePoolLansRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_K8SNodePoolLansRoutesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansRoutesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_K8SNodePoolLansRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewK8SNodePoolLansRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

