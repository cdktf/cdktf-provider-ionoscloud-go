// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkloadbalancerForwardingruleHealthCheckOutputReference interface {
	cdktf.ComplexObject
	ClientTimeout() *float64
	SetClientTimeout(val *float64)
	ClientTimeoutInput() *float64
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
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkloadbalancerForwardingruleHealthCheck
	SetInternalValue(val *NetworkloadbalancerForwardingruleHealthCheck)
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	TargetTimeout() *float64
	SetTargetTimeout(val *float64)
	TargetTimeoutInput() *float64
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
	ResetClientTimeout()
	ResetConnectTimeout()
	ResetRetries()
	ResetTargetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkloadbalancerForwardingruleHealthCheckOutputReference
type jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) InternalValue() *NetworkloadbalancerForwardingruleHealthCheck {
	var returns *NetworkloadbalancerForwardingruleHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) TargetTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) TargetTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkloadbalancerForwardingruleHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkloadbalancerForwardingruleHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkloadbalancerForwardingruleHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.NetworkloadbalancerForwardingruleHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkloadbalancerForwardingruleHealthCheckOutputReference_Override(n NetworkloadbalancerForwardingruleHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.NetworkloadbalancerForwardingruleHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetClientTimeout(val *float64) {
	if err := j.validateSetClientTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimeout",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetInternalValue(val *NetworkloadbalancerForwardingruleHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetTargetTimeout(val *float64) {
	if err := j.validateSetTargetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTimeout",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ResetClientTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetClientTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ResetRetries() {
	_jsii_.InvokeVoid(
		n,
		"resetRetries",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ResetTargetTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkloadbalancerForwardingruleHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

