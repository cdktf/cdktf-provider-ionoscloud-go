// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/vcpuserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VcpuServerNicOutputReference interface {
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
	DeviceNumber() *float64
	Dhcp() interface{}
	SetDhcp(val interface{})
	DhcpInput() interface{}
	Dhcpv6() interface{}
	SetDhcpv6(val interface{})
	Dhcpv6Input() interface{}
	Firewall() VcpuServerNicFirewallList
	FirewallActive() interface{}
	SetFirewallActive(val interface{})
	FirewallActiveInput() interface{}
	FirewallInput() interface{}
	FirewallType() *string
	SetFirewallType(val *string)
	FirewallTypeInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *VcpuServerNic
	SetInternalValue(val *VcpuServerNic)
	Ips() *[]*string
	SetIps(val *[]*string)
	IpsInput() *[]*string
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockInput() *string
	Ipv6Ips() *[]*string
	SetIpv6Ips(val *[]*string)
	Ipv6IpsInput() *[]*string
	Lan() *float64
	SetLan(val *float64)
	LanInput() *float64
	Mac() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PciSlot() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutFirewall(value interface{})
	ResetDhcp()
	ResetDhcpv6()
	ResetFirewall()
	ResetFirewallActive()
	ResetFirewallType()
	ResetIps()
	ResetIpv6CidrBlock()
	ResetIpv6Ips()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VcpuServerNicOutputReference
type jsiiProxy_VcpuServerNicOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VcpuServerNicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) DeviceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Dhcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) DhcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Dhcpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Dhcpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Firewall() VcpuServerNicFirewallList {
	var returns VcpuServerNicFirewallList
	_jsii_.Get(
		j,
		"firewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) FirewallActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) FirewallActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) FirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) FirewallType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) FirewallTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) InternalValue() *VcpuServerNic {
	var returns *VcpuServerNic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Ipv6Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Ipv6IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6IpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Lan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) LanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Mac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) PciSlot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pciSlot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServerNicOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVcpuServerNicOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VcpuServerNicOutputReference {
	_init_.Initialize()

	if err := validateNewVcpuServerNicOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VcpuServerNicOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVcpuServerNicOutputReference_Override(v VcpuServerNicOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServerNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetDhcp(val interface{}) {
	if err := j.validateSetDhcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcp",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetDhcpv6(val interface{}) {
	if err := j.validateSetDhcpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpv6",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetFirewallActive(val interface{}) {
	if err := j.validateSetFirewallActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallActive",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetFirewallType(val *string) {
	if err := j.validateSetFirewallTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallType",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetInternalValue(val *VcpuServerNic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetIps(val *[]*string) {
	if err := j.validateSetIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ips",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetIpv6Ips(val *[]*string) {
	if err := j.validateSetIpv6IpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Ips",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetLan(val *float64) {
	if err := j.validateSetLanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lan",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VcpuServerNicOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerNicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServerNicOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) PutFirewall(value interface{}) {
	if err := v.validatePutFirewallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFirewall",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetDhcp() {
	_jsii_.InvokeVoid(
		v,
		"resetDhcp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetDhcpv6() {
	_jsii_.InvokeVoid(
		v,
		"resetDhcpv6",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetFirewall() {
	_jsii_.InvokeVoid(
		v,
		"resetFirewall",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetFirewallActive() {
	_jsii_.InvokeVoid(
		v,
		"resetFirewallActive",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetFirewallType() {
	_jsii_.InvokeVoid(
		v,
		"resetFirewallType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetIps() {
	_jsii_.InvokeVoid(
		v,
		"resetIps",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetIpv6Ips() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6Ips",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServerNicOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VcpuServerNicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

