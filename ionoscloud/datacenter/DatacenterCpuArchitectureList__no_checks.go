// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacenter

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatacenterCpuArchitectureList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatacenterCpuArchitectureList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatacenterCpuArchitectureList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatacenterCpuArchitectureList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatacenterCpuArchitectureList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatacenterCpuArchitectureList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatacenterCpuArchitectureListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

