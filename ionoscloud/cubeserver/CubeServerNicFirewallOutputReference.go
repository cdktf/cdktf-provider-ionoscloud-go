// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v9/cubeserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CubeServerNicFirewallOutputReference interface {
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
	InternalValue() *CubeServerNicFirewall
	SetInternalValue(val *CubeServerNicFirewall)
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

// The jsii proxy struct for CubeServerNicFirewallOutputReference
type jsiiProxy_CubeServerNicFirewallOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) IcmpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) IcmpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) InternalValue() *CubeServerNicFirewall {
	var returns *CubeServerNicFirewall
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) PortRangeEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) PortRangeEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) PortRangeStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) PortRangeStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) SourceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) SourceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) SourceMac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) SourceMacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) TargetIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) TargetIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCubeServerNicFirewallOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CubeServerNicFirewallOutputReference {
	_init_.Initialize()

	if err := validateNewCubeServerNicFirewallOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CubeServerNicFirewallOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCubeServerNicFirewallOutputReference_Override(c CubeServerNicFirewallOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetIcmpCode(val *string) {
	if err := j.validateSetIcmpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpCode",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetInternalValue(val *CubeServerNicFirewall) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetPortRangeEnd(val *float64) {
	if err := j.validateSetPortRangeEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeEnd",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetPortRangeStart(val *float64) {
	if err := j.validateSetPortRangeStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeStart",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetSourceIp(val *string) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetSourceMac(val *string) {
	if err := j.validateSetSourceMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMac",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetTargetIp(val *string) {
	if err := j.validateSetTargetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIp",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicFirewallOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetIcmpCode() {
	_jsii_.InvokeVoid(
		c,
		"resetIcmpCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetIcmpType() {
	_jsii_.InvokeVoid(
		c,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetPortRangeEnd() {
	_jsii_.InvokeVoid(
		c,
		"resetPortRangeEnd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetPortRangeStart() {
	_jsii_.InvokeVoid(
		c,
		"resetPortRangeStart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetSourceMac() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceMac",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetTargetIp() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicFirewallOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

