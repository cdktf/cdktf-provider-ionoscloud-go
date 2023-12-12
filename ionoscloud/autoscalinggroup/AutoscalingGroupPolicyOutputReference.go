// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v10/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupPolicyOutputReference interface {
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
	InternalValue() *AutoscalingGroupPolicy
	SetInternalValue(val *AutoscalingGroupPolicy)
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Range() *string
	SetRange(val *string)
	RangeInput() *string
	ScaleInAction() AutoscalingGroupPolicyScaleInActionOutputReference
	ScaleInActionInput() *AutoscalingGroupPolicyScaleInAction
	ScaleInThreshold() *float64
	SetScaleInThreshold(val *float64)
	ScaleInThresholdInput() *float64
	ScaleOutAction() AutoscalingGroupPolicyScaleOutActionOutputReference
	ScaleOutActionInput() *AutoscalingGroupPolicyScaleOutAction
	ScaleOutThreshold() *float64
	SetScaleOutThreshold(val *float64)
	ScaleOutThresholdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
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
	PutScaleInAction(value *AutoscalingGroupPolicyScaleInAction)
	PutScaleOutAction(value *AutoscalingGroupPolicyScaleOutAction)
	ResetRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupPolicyOutputReference
type jsiiProxy_AutoscalingGroupPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) InternalValue() *AutoscalingGroupPolicy {
	var returns *AutoscalingGroupPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleInAction() AutoscalingGroupPolicyScaleInActionOutputReference {
	var returns AutoscalingGroupPolicyScaleInActionOutputReference
	_jsii_.Get(
		j,
		"scaleInAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleInActionInput() *AutoscalingGroupPolicyScaleInAction {
	var returns *AutoscalingGroupPolicyScaleInAction
	_jsii_.Get(
		j,
		"scaleInActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleInThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleInThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleOutAction() AutoscalingGroupPolicyScaleOutActionOutputReference {
	var returns AutoscalingGroupPolicyScaleOutActionOutputReference
	_jsii_.Get(
		j,
		"scaleOutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleOutActionInput() *AutoscalingGroupPolicyScaleOutAction {
	var returns *AutoscalingGroupPolicyScaleOutAction
	_jsii_.Get(
		j,
		"scaleOutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleOutThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) ScaleOutThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupPolicyOutputReference_Override(a AutoscalingGroupPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetInternalValue(val *AutoscalingGroupPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetRange(val *string) {
	if err := j.validateSetRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"range",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetScaleInThreshold(val *float64) {
	if err := j.validateSetScaleInThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInThreshold",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetScaleOutThreshold(val *float64) {
	if err := j.validateSetScaleOutThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutThreshold",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupPolicyOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) PutScaleInAction(value *AutoscalingGroupPolicyScaleInAction) {
	if err := a.validatePutScaleInActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScaleInAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) PutScaleOutAction(value *AutoscalingGroupPolicyScaleOutAction) {
	if err := a.validatePutScaleOutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScaleOutAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		a,
		"resetRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingGroupPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

