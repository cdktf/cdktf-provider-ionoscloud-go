// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-ionoscloud-go/ionoscloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IpfailoverTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDefault()
	ResetDelete()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IpfailoverTimeoutsOutputReference
type jsiiProxy_IpfailoverTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewIpfailoverTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IpfailoverTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IpfailoverTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.IpfailoverTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIpfailoverTimeoutsOutputReference_Override(i IpfailoverTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.IpfailoverTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetDefault(val *string) {
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IpfailoverTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		i,
		"resetCreate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		i,
		"resetDefault",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		i,
		"resetDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpfailoverTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

