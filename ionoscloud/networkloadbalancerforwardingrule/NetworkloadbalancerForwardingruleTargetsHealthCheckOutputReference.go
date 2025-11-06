// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkloadbalancerforwardingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/networkloadbalancerforwardingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference interface {
	cdktf.ComplexObject
	Check() interface{}
	SetCheck(val interface{})
	CheckInput() interface{}
	CheckInterval() *float64
	SetCheckInterval(val *float64)
	CheckIntervalInput() *float64
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
	InternalValue() *NetworkloadbalancerForwardingruleTargetsHealthCheck
	SetInternalValue(val *NetworkloadbalancerForwardingruleTargetsHealthCheck)
	Maintenance() interface{}
	SetMaintenance(val interface{})
	MaintenanceInput() interface{}
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
	ResetCheck()
	ResetCheckInterval()
	ResetMaintenance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference
type jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) Check() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"check",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) CheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) CheckInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) CheckIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) InternalValue() *NetworkloadbalancerForwardingruleTargetsHealthCheck {
	var returns *NetworkloadbalancerForwardingruleTargetsHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) Maintenance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) MaintenanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkloadbalancerForwardingruleTargetsHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.networkloadbalancerForwardingrule.NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference_Override(n NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.networkloadbalancerForwardingrule.NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetCheck(val interface{}) {
	if err := j.validateSetCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"check",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetCheckInterval(val *float64) {
	if err := j.validateSetCheckIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkInterval",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetInternalValue(val *NetworkloadbalancerForwardingruleTargetsHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetMaintenance(val interface{}) {
	if err := j.validateSetMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenance",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ResetCheck() {
	_jsii_.InvokeVoid(
		n,
		"resetCheck",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ResetCheckInterval() {
	_jsii_.InvokeVoid(
		n,
		"resetCheckInterval",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ResetMaintenance() {
	_jsii_.InvokeVoid(
		n,
		"resetMaintenance",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleTargetsHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

