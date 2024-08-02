// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/vpnipsectunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnIpsecTunnelEspOutputReference interface {
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
	DiffieHellmanGroup() *string
	SetDiffieHellmanGroup(val *string)
	DiffieHellmanGroupInput() *string
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	IntegrityAlgorithm() *string
	SetIntegrityAlgorithm(val *string)
	IntegrityAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lifetime() *float64
	SetLifetime(val *float64)
	LifetimeInput() *float64
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
	ResetDiffieHellmanGroup()
	ResetEncryptionAlgorithm()
	ResetIntegrityAlgorithm()
	ResetLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnIpsecTunnelEspOutputReference
type jsiiProxy_VpnIpsecTunnelEspOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) DiffieHellmanGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diffieHellmanGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) DiffieHellmanGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diffieHellmanGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) IntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) IntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) Lifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) LifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnIpsecTunnelEspOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnIpsecTunnelEspOutputReference {
	_init_.Initialize()

	if err := validateNewVpnIpsecTunnelEspOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnIpsecTunnelEspOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecTunnel.VpnIpsecTunnelEspOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnIpsecTunnelEspOutputReference_Override(v VpnIpsecTunnelEspOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecTunnel.VpnIpsecTunnelEspOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetDiffieHellmanGroup(val *string) {
	if err := j.validateSetDiffieHellmanGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diffieHellmanGroup",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetIntegrityAlgorithm(val *string) {
	if err := j.validateSetIntegrityAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetLifetime(val *float64) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelEspOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ResetDiffieHellmanGroup() {
	_jsii_.InvokeVoid(
		v,
		"resetDiffieHellmanGroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ResetIntegrityAlgorithm() {
	_jsii_.InvokeVoid(
		v,
		"resetIntegrityAlgorithm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		v,
		"resetLifetime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnIpsecTunnelEspOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

