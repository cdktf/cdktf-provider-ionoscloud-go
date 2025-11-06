// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/inmemorydbreplicaset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InmemorydbReplicasetCredentialsOutputReference interface {
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
	HashedPassword() InmemorydbReplicasetCredentialsHashedPasswordOutputReference
	HashedPasswordInput() *InmemorydbReplicasetCredentialsHashedPassword
	InternalValue() *InmemorydbReplicasetCredentials
	SetInternalValue(val *InmemorydbReplicasetCredentials)
	PlainTextPassword() *string
	SetPlainTextPassword(val *string)
	PlainTextPasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutHashedPassword(value *InmemorydbReplicasetCredentialsHashedPassword)
	ResetHashedPassword()
	ResetPlainTextPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InmemorydbReplicasetCredentialsOutputReference
type jsiiProxy_InmemorydbReplicasetCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) HashedPassword() InmemorydbReplicasetCredentialsHashedPasswordOutputReference {
	var returns InmemorydbReplicasetCredentialsHashedPasswordOutputReference
	_jsii_.Get(
		j,
		"hashedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) HashedPasswordInput() *InmemorydbReplicasetCredentialsHashedPassword {
	var returns *InmemorydbReplicasetCredentialsHashedPassword
	_jsii_.Get(
		j,
		"hashedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) InternalValue() *InmemorydbReplicasetCredentials {
	var returns *InmemorydbReplicasetCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) PlainTextPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainTextPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) PlainTextPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainTextPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewInmemorydbReplicasetCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InmemorydbReplicasetCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewInmemorydbReplicasetCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InmemorydbReplicasetCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicasetCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInmemorydbReplicasetCredentialsOutputReference_Override(i InmemorydbReplicasetCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicasetCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetInternalValue(val *InmemorydbReplicasetCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetPlainTextPassword(val *string) {
	if err := j.validateSetPlainTextPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plainTextPassword",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) PutHashedPassword(value *InmemorydbReplicasetCredentialsHashedPassword) {
	if err := i.validatePutHashedPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHashedPassword",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ResetHashedPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetHashedPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ResetPlainTextPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetPlainTextPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicasetCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

