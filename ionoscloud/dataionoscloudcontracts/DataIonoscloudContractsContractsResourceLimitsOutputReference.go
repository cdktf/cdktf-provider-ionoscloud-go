// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudcontracts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/dataionoscloudcontracts/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudContractsContractsResourceLimitsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CoresPerContract() *float64
	CoresPerServer() *float64
	CoresProvisioned() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DasVolumeProvisioned() *float64
	// Experimental.
	Fqn() *string
	HddLimitPerContract() *float64
	HddLimitPerVolume() *float64
	HddVolumeProvisioned() *float64
	InternalValue() *DataIonoscloudContractsContractsResourceLimits
	SetInternalValue(val *DataIonoscloudContractsContractsResourceLimits)
	K8SClusterLimitTotal() *float64
	K8SClustersProvisioned() *float64
	NatGatewayLimitTotal() *float64
	NatGatewayProvisioned() *float64
	NlbLimitTotal() *float64
	NlbProvisioned() *float64
	RamPerContract() *float64
	RamPerServer() *float64
	RamProvisioned() *float64
	ReservableIps() *float64
	ReservedIpsInUse() *float64
	ReservedIpsOnContract() *float64
	RulesPerSecurityGroup() *float64
	SecurityGroupsPerResource() *float64
	SecurityGroupsPerVdc() *float64
	SsdLimitPerContract() *float64
	SsdLimitPerVolume() *float64
	SsdVolumeProvisioned() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataIonoscloudContractsContractsResourceLimitsOutputReference
type jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) CoresPerContract() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresPerContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) CoresPerServer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresPerServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) CoresProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) DasVolumeProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dasVolumeProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) HddLimitPerContract() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hddLimitPerContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) HddLimitPerVolume() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hddLimitPerVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) HddVolumeProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hddVolumeProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) InternalValue() *DataIonoscloudContractsContractsResourceLimits {
	var returns *DataIonoscloudContractsContractsResourceLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) K8SClusterLimitTotal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"k8SClusterLimitTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) K8SClustersProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"k8SClustersProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) NatGatewayLimitTotal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"natGatewayLimitTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) NatGatewayProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"natGatewayProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) NlbLimitTotal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nlbLimitTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) NlbProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nlbProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) RamPerContract() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramPerContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) RamPerServer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramPerServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) RamProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ReservableIps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservableIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ReservedIpsInUse() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedIpsInUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ReservedIpsOnContract() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedIpsOnContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) RulesPerSecurityGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rulesPerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) SecurityGroupsPerResource() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityGroupsPerResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) SecurityGroupsPerVdc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityGroupsPerVdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) SsdLimitPerContract() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssdLimitPerContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) SsdLimitPerVolume() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssdLimitPerVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) SsdVolumeProvisioned() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssdVolumeProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataIonoscloudContractsContractsResourceLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataIonoscloudContractsContractsResourceLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewDataIonoscloudContractsContractsResourceLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudContracts.DataIonoscloudContractsContractsResourceLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataIonoscloudContractsContractsResourceLimitsOutputReference_Override(d DataIonoscloudContractsContractsResourceLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudContracts.DataIonoscloudContractsContractsResourceLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference)SetInternalValue(val *DataIonoscloudContractsContractsResourceLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudContractsContractsResourceLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

