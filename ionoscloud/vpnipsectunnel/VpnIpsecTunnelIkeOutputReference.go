// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnipsectunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/vpnipsectunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnIpsecTunnelIkeOutputReference interface {
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
	InternalValue() *VpnIpsecTunnelIke
	SetInternalValue(val *VpnIpsecTunnelIke)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetDiffieHellmanGroup()
	ResetEncryptionAlgorithm()
	ResetIntegrityAlgorithm()
	ResetLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnIpsecTunnelIkeOutputReference
type jsiiProxy_VpnIpsecTunnelIkeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) DiffieHellmanGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diffieHellmanGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) DiffieHellmanGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diffieHellmanGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) IntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) IntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) InternalValue() *VpnIpsecTunnelIke {
	var returns *VpnIpsecTunnelIke
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) Lifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) LifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnIpsecTunnelIkeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnIpsecTunnelIkeOutputReference {
	_init_.Initialize()

	if err := validateNewVpnIpsecTunnelIkeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnIpsecTunnelIkeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecTunnel.VpnIpsecTunnelIkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnIpsecTunnelIkeOutputReference_Override(v VpnIpsecTunnelIkeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnIpsecTunnel.VpnIpsecTunnelIkeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetDiffieHellmanGroup(val *string) {
	if err := j.validateSetDiffieHellmanGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diffieHellmanGroup",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetIntegrityAlgorithm(val *string) {
	if err := j.validateSetIntegrityAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetInternalValue(val *VpnIpsecTunnelIke) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetLifetime(val *float64) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnIpsecTunnelIkeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ResetDiffieHellmanGroup() {
	_jsii_.InvokeVoid(
		v,
		"resetDiffieHellmanGroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ResetIntegrityAlgorithm() {
	_jsii_.InvokeVoid(
		v,
		"resetIntegrityAlgorithm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		v,
		"resetLifetime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnIpsecTunnelIkeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

