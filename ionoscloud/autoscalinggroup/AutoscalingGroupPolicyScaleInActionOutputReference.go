// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupPolicyScaleInActionOutputReference interface {
	cdktf.ComplexObject
	Amount() *float64
	SetAmount(val *float64)
	AmountInput() *float64
	AmountType() *string
	SetAmountType(val *string)
	AmountTypeInput() *string
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
	CooldownPeriod() *string
	SetCooldownPeriod(val *string)
	CooldownPeriodInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteVolumes() interface{}
	SetDeleteVolumes(val interface{})
	DeleteVolumesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AutoscalingGroupPolicyScaleInAction
	SetInternalValue(val *AutoscalingGroupPolicyScaleInAction)
	TerminationPolicyType() *string
	SetTerminationPolicyType(val *string)
	TerminationPolicyTypeInput() *string
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
	ResetCooldownPeriod()
	ResetTerminationPolicyType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupPolicyScaleInActionOutputReference
type jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) Amount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) AmountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) AmountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) AmountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) CooldownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) CooldownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) DeleteVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) DeleteVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) InternalValue() *AutoscalingGroupPolicyScaleInAction {
	var returns *AutoscalingGroupPolicyScaleInAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) TerminationPolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationPolicyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) TerminationPolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationPolicyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupPolicyScaleInActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupPolicyScaleInActionOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupPolicyScaleInActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupPolicyScaleInActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupPolicyScaleInActionOutputReference_Override(a AutoscalingGroupPolicyScaleInActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupPolicyScaleInActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetAmount(val *float64) {
	if err := j.validateSetAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetAmountType(val *string) {
	if err := j.validateSetAmountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amountType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetCooldownPeriod(val *string) {
	if err := j.validateSetCooldownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownPeriod",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetDeleteVolumes(val interface{}) {
	if err := j.validateSetDeleteVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVolumes",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetInternalValue(val *AutoscalingGroupPolicyScaleInAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetTerminationPolicyType(val *string) {
	if err := j.validateSetTerminationPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationPolicyType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ResetCooldownPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetCooldownPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ResetTerminationPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingGroupPolicyScaleInActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

