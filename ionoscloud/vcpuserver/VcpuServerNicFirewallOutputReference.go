// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/vcpuserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VcpuServerNicFirewallOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IcmpCode() *string
	SetIcmpCode(val *string)
	IcmpCodeInput() *string
	IcmpType() *string
	SetIcmpType(val *string)
	IcmpTypeInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PortRangeEnd() *float64
	SetPortRangeEnd(val *float64)
	PortRangeEndInput() *float64
	PortRangeStart() *float64
	SetPortRangeStart(val *float64)
	PortRangeStartInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SourceIp() *string
	SetSourceIp(val *string)
	SourceIpInput() *string
	SourceMac() *string
	SetSourceMac(val *string)
	SourceMacInput() *string
	TargetIp() *string
	SetTargetIp(val *string)
	TargetIpInput() *string
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
	ResetIcmpCode()
	ResetIcmpType()
	ResetName()
	ResetPortRangeEnd()
	ResetPortRangeStart()
	ResetSourceIp()
	ResetSourceMac()
	ResetTargetIp()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VcpuServerNicFirewallOutputReference
type jsiiProxy_VcpuServerNicFirewallOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) IcmpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) IcmpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) PortRangeEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) PortRangeEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) PortRangeStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) PortRangeStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) SourceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) SourceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) SourceMac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) SourceMacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) TargetIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) TargetIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewVcpuServerNicFirewallOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VcpuServerNicFirewallOutputReference {
	_init_.Initialize()

	if err := validateNewVcpuServerNicFirewallOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VcpuServerNicFirewallOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVcpuServerNicFirewallOutputReference_Override(v VcpuServerNicFirewallOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetIcmpCode(val *string) {
	if err := j.validateSetIcmpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpCode",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetPortRangeEnd(val *float64) {
	if err := j.validateSetPortRangeEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeEnd",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetPortRangeStart(val *float64) {
	if err := j.validateSetPortRangeStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeStart",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetSourceIp(val *string) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetSourceMac(val *string) {
	if err := j.validateSetSourceMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMac",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetTargetIp(val *string) {
	if err := j.validateSetTargetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIp",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicFirewallOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetIcmpCode() {
	_jsii_.InvokeVoid(
		v,
		"resetIcmpCode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetIcmpType() {
	_jsii_.InvokeVoid(
		v,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetPortRangeEnd() {
	_jsii_.InvokeVoid(
		v,
		"resetPortRangeEnd",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetPortRangeStart() {
	_jsii_.InvokeVoid(
		v,
		"resetPortRangeStart",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetSourceMac() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceMac",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetTargetIp() {
	_jsii_.InvokeVoid(
		v,
		"resetTargetIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		v,
		"resetType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VcpuServerNicFirewallOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

