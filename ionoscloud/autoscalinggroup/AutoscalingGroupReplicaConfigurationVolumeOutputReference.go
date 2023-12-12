// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupReplicaConfigurationVolumeOutputReference interface {
	cdktf.ComplexObject
	BackupUnitId() *string
	SetBackupUnitId(val *string)
	BackupUnitIdInput() *string
	BootOrder() *string
	SetBootOrder(val *string)
	BootOrderInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageAlias() *string
	SetImageAlias(val *string)
	ImageAliasInput() *string
	ImageInput() *string
	ImagePassword() *string
	SetImagePassword(val *string)
	ImagePasswordInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetBackupUnitId()
	ResetBus()
	ResetImage()
	ResetImageAlias()
	ResetImagePassword()
	ResetSshKeys()
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

// The jsii proxy struct for AutoscalingGroupReplicaConfigurationVolumeOutputReference
type jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) BackupUnitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) BackupUnitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) BootOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) BootOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Bus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) BusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ImageAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ImageAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ImagePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupReplicaConfigurationVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingGroupReplicaConfigurationVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupReplicaConfigurationVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupReplicaConfigurationVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingGroupReplicaConfigurationVolumeOutputReference_Override(a AutoscalingGroupReplicaConfigurationVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupReplicaConfigurationVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetBackupUnitId(val *string) {
	if err := j.validateSetBackupUnitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupUnitId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetBootOrder(val *string) {
	if err := j.validateSetBootOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootOrder",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetBus(val *string) {
	if err := j.validateSetBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bus",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetImageAlias(val *string) {
	if err := j.validateSetImageAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageAlias",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetImagePassword(val *string) {
	if err := j.validateSetImagePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePassword",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetBackupUnitId() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupUnitId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetBus() {
	_jsii_.InvokeVoid(
		a,
		"resetBus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetImageAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetImageAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetImagePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetImagePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

