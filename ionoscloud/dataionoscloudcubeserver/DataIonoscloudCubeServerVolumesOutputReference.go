// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudcubeserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/dataionoscloudcubeserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudCubeServerVolumesOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZone() *string
	BackupUnitId() *string
	BootServer() *string
	Bus() *string
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
	CpuHotPlug() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceNumber() *float64
	DiscVirtioHotPlug() cdktf.IResolvable
	DiscVirtioHotUnplug() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	Id() *string
	ImageName() *string
	ImagePassword() *string
	InternalValue() *DataIonoscloudCubeServerVolumes
	SetInternalValue(val *DataIonoscloudCubeServerVolumes)
	LicenceType() *string
	Name() *string
	NicHotPlug() cdktf.IResolvable
	NicHotUnplug() cdktf.IResolvable
	PciSlot() *float64
	RamHotPlug() cdktf.IResolvable
	SshKeys() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	UserData() *string
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

// The jsii proxy struct for DataIonoscloudCubeServerVolumesOutputReference
type jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) BackupUnitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) BootServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Bus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) CpuHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) DeviceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) DiscVirtioHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) DiscVirtioHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) InternalValue() *DataIonoscloudCubeServerVolumes {
	var returns *DataIonoscloudCubeServerVolumes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) LicenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) NicHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) NicHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) PciSlot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pciSlot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) RamHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ramHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


func NewDataIonoscloudCubeServerVolumesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataIonoscloudCubeServerVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewDataIonoscloudCubeServerVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudCubeServer.DataIonoscloudCubeServerVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataIonoscloudCubeServerVolumesOutputReference_Override(d DataIonoscloudCubeServerVolumesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudCubeServer.DataIonoscloudCubeServerVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference)SetInternalValue(val *DataIonoscloudCubeServerVolumes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataIonoscloudCubeServerVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

