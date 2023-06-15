package server

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v7/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v7/server/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerNicFirewallOutputReference interface {
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
	InternalValue() *ServerNicFirewall
	SetInternalValue(val *ServerNicFirewall)
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

// The jsii proxy struct for ServerNicFirewallOutputReference
type jsiiProxy_ServerNicFirewallOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) IcmpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) IcmpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) InternalValue() *ServerNicFirewall {
	var returns *ServerNicFirewall
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) PortRangeEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) PortRangeEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) PortRangeStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) PortRangeStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) SourceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) SourceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) SourceMac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) SourceMacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) TargetIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) TargetIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerNicFirewallOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewServerNicFirewallOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServerNicFirewallOutputReference {
	_init_.Initialize()

	if err := validateNewServerNicFirewallOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerNicFirewallOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.server.ServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServerNicFirewallOutputReference_Override(s ServerNicFirewallOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.server.ServerNicFirewallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetIcmpCode(val *string) {
	if err := j.validateSetIcmpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpCode",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetInternalValue(val *ServerNicFirewall) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetPortRangeEnd(val *float64) {
	if err := j.validateSetPortRangeEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeEnd",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetPortRangeStart(val *float64) {
	if err := j.validateSetPortRangeStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeStart",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetSourceIp(val *string) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetSourceMac(val *string) {
	if err := j.validateSetSourceMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMac",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetTargetIp(val *string) {
	if err := j.validateSetTargetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIp",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServerNicFirewallOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetIcmpCode() {
	_jsii_.InvokeVoid(
		s,
		"resetIcmpCode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetIcmpType() {
	_jsii_.InvokeVoid(
		s,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetPortRangeEnd() {
	_jsii_.InvokeVoid(
		s,
		"resetPortRangeEnd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetPortRangeStart() {
	_jsii_.InvokeVoid(
		s,
		"resetPortRangeStart",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetSourceMac() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceMac",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetTargetIp() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerNicFirewallOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

