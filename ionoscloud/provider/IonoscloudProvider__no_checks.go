// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IonoscloudProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_IonoscloudProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateIonoscloudProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateIonoscloudProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIonoscloudProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateIonoscloudProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_IonoscloudProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewIonoscloudProviderParameters(scope constructs.Construct, id *string, config *IonoscloudProviderConfig) error {
	return nil
}

