// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/cubeserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CubeServerNicOutputReference interface {
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
	Firewall() CubeServerNicFirewallOutputReference
	FirewallActive() interface{}
	SetFirewallActive(val interface{})
	FirewallActiveInput() interface{}
	FirewallInput() *CubeServerNicFirewall
	FirewallType() *string
	SetFirewallType(val *string)
	FirewallTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CubeServerNic
	SetInternalValue(val *CubeServerNic)
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
	SetMac(val *string)
	MacInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PciSlot() *float64
	SecurityGroupsIds() *[]*string
	SetSecurityGroupsIds(val *[]*string)
	SecurityGroupsIdsInput() *[]*string
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
	PutFirewall(value *CubeServerNicFirewall)
	ResetDhcp()
	ResetDhcpv6()
	ResetFirewall()
	ResetFirewallActive()
	ResetFirewallType()
	ResetIps()
	ResetIpv6CidrBlock()
	ResetIpv6Ips()
	ResetMac()
	ResetName()
	ResetSecurityGroupsIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CubeServerNicOutputReference
type jsiiProxy_CubeServerNicOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CubeServerNicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) DeviceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Dhcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) DhcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Dhcpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Dhcpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Firewall() CubeServerNicFirewallOutputReference {
	var returns CubeServerNicFirewallOutputReference
	_jsii_.Get(
		j,
		"firewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) FirewallActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) FirewallActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) FirewallInput() *CubeServerNicFirewall {
	var returns *CubeServerNicFirewall
	_jsii_.Get(
		j,
		"firewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) FirewallType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) FirewallTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) InternalValue() *CubeServerNic {
	var returns *CubeServerNic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Ipv6Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Ipv6IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6IpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Lan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) LanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Mac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) MacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) PciSlot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pciSlot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) SecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) SecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerNicOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCubeServerNicOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CubeServerNicOutputReference {
	_init_.Initialize()

	if err := validateNewCubeServerNicOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CubeServerNicOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCubeServerNicOutputReference_Override(c CubeServerNicOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetDhcp(val interface{}) {
	if err := j.validateSetDhcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcp",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetDhcpv6(val interface{}) {
	if err := j.validateSetDhcpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpv6",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetFirewallActive(val interface{}) {
	if err := j.validateSetFirewallActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallActive",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetFirewallType(val *string) {
	if err := j.validateSetFirewallTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallType",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetInternalValue(val *CubeServerNic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetIps(val *[]*string) {
	if err := j.validateSetIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ips",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetIpv6Ips(val *[]*string) {
	if err := j.validateSetIpv6IpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Ips",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetLan(val *float64) {
	if err := j.validateSetLanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lan",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetMac(val *string) {
	if err := j.validateSetMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mac",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CubeServerNicOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CubeServerNicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CubeServerNicOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerNicOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CubeServerNicOutputReference) PutFirewall(value *CubeServerNicFirewall) {
	if err := c.validatePutFirewallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirewall",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetDhcp() {
	_jsii_.InvokeVoid(
		c,
		"resetDhcp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetDhcpv6() {
	_jsii_.InvokeVoid(
		c,
		"resetDhcpv6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetFirewall() {
	_jsii_.InvokeVoid(
		c,
		"resetFirewall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetFirewallActive() {
	_jsii_.InvokeVoid(
		c,
		"resetFirewallActive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetFirewallType() {
	_jsii_.InvokeVoid(
		c,
		"resetFirewallType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetIps() {
	_jsii_.InvokeVoid(
		c,
		"resetIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetIpv6Ips() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6Ips",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetMac() {
	_jsii_.InvokeVoid(
		c,
		"resetMac",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) ResetSecurityGroupsIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupsIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerNicOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CubeServerNicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

