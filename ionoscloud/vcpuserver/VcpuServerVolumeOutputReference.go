// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/vcpuserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VcpuServerVolumeOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupUnitId() *string
	SetBackupUnitId(val *string)
	BackupUnitIdInput() *string
	BootServer() *string
	Bus() *string
	SetBus(val *string)
	BusInput() *string
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
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	ExposeSerial() interface{}
	SetExposeSerial(val interface{})
	ExposeSerialInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VcpuServerVolume
	SetInternalValue(val *VcpuServerVolume)
	LicenceType() *string
	SetLicenceType(val *string)
	LicenceTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NicHotPlug() cdktf.IResolvable
	NicHotUnplug() cdktf.IResolvable
	PciSlot() *float64
	RamHotPlug() cdktf.IResolvable
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAvailabilityZone()
	ResetBackupUnitId()
	ResetBus()
	ResetExposeSerial()
	ResetLicenceType()
	ResetName()
	ResetSize()
	ResetUserData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VcpuServerVolumeOutputReference
type jsiiProxy_VcpuServerVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) BackupUnitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) BackupUnitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) BootServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) Bus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) BusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) CpuHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) DeviceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) DiscVirtioHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) DiscVirtioHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) ExposeSerial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeSerial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) ExposeSerialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeSerialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) InternalValue() *VcpuServerVolume {
	var returns *VcpuServerVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) LicenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) LicenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) NicHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) NicHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) PciSlot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pciSlot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) RamHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ramHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewVcpuServerVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VcpuServerVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewVcpuServerVolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VcpuServerVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVcpuServerVolumeOutputReference_Override(v VcpuServerVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetBackupUnitId(val *string) {
	if err := j.validateSetBackupUnitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupUnitId",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetBus(val *string) {
	if err := j.validateSetBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bus",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetExposeSerial(val interface{}) {
	if err := j.validateSetExposeSerialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeSerial",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetInternalValue(val *VcpuServerVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetLicenceType(val *string) {
	if err := j.validateSetLicenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenceType",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VcpuServerVolumeOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		v,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetBackupUnitId() {
	_jsii_.InvokeVoid(
		v,
		"resetBackupUnitId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetBus() {
	_jsii_.InvokeVoid(
		v,
		"resetBus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetExposeSerial() {
	_jsii_.InvokeVoid(
		v,
		"resetExposeSerial",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetLicenceType() {
	_jsii_.InvokeVoid(
		v,
		"resetLicenceType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		v,
		"resetSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		v,
		"resetUserData",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

