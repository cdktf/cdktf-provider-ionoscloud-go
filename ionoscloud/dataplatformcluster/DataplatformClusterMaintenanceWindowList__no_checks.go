//go:build no_runtime_type_checking

package dataplatformcluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplatformClusterMaintenanceWindowList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplatformClusterMaintenanceWindowListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
