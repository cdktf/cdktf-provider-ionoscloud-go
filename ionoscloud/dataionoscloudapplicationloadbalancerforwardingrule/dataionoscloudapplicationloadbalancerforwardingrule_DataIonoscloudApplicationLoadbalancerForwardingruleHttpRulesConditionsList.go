package dataionoscloudapplicationloadbalancerforwardingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v3/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v3/dataionoscloudapplicationloadbalancerforwardingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList
type jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList {
	_init_.Initialize()

	if err := validateNewDataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudApplicationLoadbalancerForwardingrule.DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList_Override(d DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudApplicationLoadbalancerForwardingrule.DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) Get(index *float64) DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudApplicationLoadbalancerForwardingruleHttpRulesConditionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
