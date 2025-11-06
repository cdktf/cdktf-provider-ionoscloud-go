// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsecgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/vpnipsecgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnIpsecGatewayConnectionsOutputReference interface {
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
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4Cidr() *string
	SetIpv4Cidr(val *string)
	Ipv4CidrInput() *string
	Ipv6Cidr() *string
	SetIpv6Cidr(val *string)
	Ipv6CidrInput() *string
	LanId() *string
	SetLanId(val *string)
	LanIdInput() *string
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
	ResetIpv6Cidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnIpsecGatewayConnectionsOutputReference
type jsiiProxy_VpnIpsecGatewayConnectionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Ipv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Ipv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Ipv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Ipv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) LanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) LanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnIpsecGatewayConnectionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnIpsecGatewayConnectionsOutputReference {
	_init_.Initialize()

	if err := validateNewVpnIpsecGatewayConnectionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnIpsecGatewayConnectionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecGateway.VpnIpsecGatewayConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnIpsecGatewayConnectionsOutputReference_Override(v VpnIpsecGatewayConnectionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecGateway.VpnIpsecGatewayConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetIpv4Cidr(val *string) {
	if err := j.validateSetIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetIpv6Cidr(val *string) {
	if err := j.validateSetIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetLanId(val *string) {
	if err := j.validateSetLanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanId",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) ResetIpv6Cidr() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6Cidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecGatewayConnectionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

