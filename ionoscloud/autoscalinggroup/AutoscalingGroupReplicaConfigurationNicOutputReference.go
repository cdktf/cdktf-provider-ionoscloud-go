// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupReplicaConfigurationNicOutputReference interface {
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
	Dhcp() interface{}
	SetDhcp(val interface{})
	DhcpInput() interface{}
	FirewallActive() interface{}
	SetFirewallActive(val interface{})
	FirewallActiveInput() interface{}
	FirewallRule() AutoscalingGroupReplicaConfigurationNicFirewallRuleList
	FirewallRuleInput() interface{}
	FirewallType() *string
	SetFirewallType(val *string)
	FirewallTypeInput() *string
	FlowLog() AutoscalingGroupReplicaConfigurationNicFlowLogList
	FlowLogInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lan() *float64
	SetLan(val *float64)
	LanInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	TargetGroup() AutoscalingGroupReplicaConfigurationNicTargetGroupOutputReference
	TargetGroupInput() *AutoscalingGroupReplicaConfigurationNicTargetGroup
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
	PutFirewallRule(value interface{})
	PutFlowLog(value interface{})
	PutTargetGroup(value *AutoscalingGroupReplicaConfigurationNicTargetGroup)
	ResetDhcp()
	ResetFirewallActive()
	ResetFirewallRule()
	ResetFirewallType()
	ResetFlowLog()
	ResetTargetGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupReplicaConfigurationNicOutputReference
type jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) Dhcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) DhcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallRule() AutoscalingGroupReplicaConfigurationNicFirewallRuleList {
	var returns AutoscalingGroupReplicaConfigurationNicFirewallRuleList
	_jsii_.Get(
		j,
		"firewallRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FirewallTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FlowLog() AutoscalingGroupReplicaConfigurationNicFlowLogList {
	var returns AutoscalingGroupReplicaConfigurationNicFlowLogList
	_jsii_.Get(
		j,
		"flowLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) FlowLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) Lan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) LanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) TargetGroup() AutoscalingGroupReplicaConfigurationNicTargetGroupOutputReference {
	var returns AutoscalingGroupReplicaConfigurationNicTargetGroupOutputReference
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) TargetGroupInput() *AutoscalingGroupReplicaConfigurationNicTargetGroup {
	var returns *AutoscalingGroupReplicaConfigurationNicTargetGroup
	_jsii_.Get(
		j,
		"targetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupReplicaConfigurationNicOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingGroupReplicaConfigurationNicOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupReplicaConfigurationNicOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupReplicaConfigurationNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingGroupReplicaConfigurationNicOutputReference_Override(a AutoscalingGroupReplicaConfigurationNicOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.autoscalingGroup.AutoscalingGroupReplicaConfigurationNicOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetDhcp(val interface{}) {
	if err := j.validateSetDhcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcp",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetFirewallActive(val interface{}) {
	if err := j.validateSetFirewallActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallActive",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetFirewallType(val *string) {
	if err := j.validateSetFirewallTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetLan(val *float64) {
	if err := j.validateSetLanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lan",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) PutFirewallRule(value interface{}) {
	if err := a.validatePutFirewallRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirewallRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) PutFlowLog(value interface{}) {
	if err := a.validatePutFlowLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFlowLog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) PutTargetGroup(value *AutoscalingGroupReplicaConfigurationNicTargetGroup) {
	if err := a.validatePutTargetGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetDhcp() {
	_jsii_.InvokeVoid(
		a,
		"resetDhcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetFirewallActive() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallActive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetFirewallRule() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetFirewallType() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetFlowLog() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowLog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ResetTargetGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingGroupReplicaConfigurationNicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

