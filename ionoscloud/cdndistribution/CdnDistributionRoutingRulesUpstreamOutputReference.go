// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdndistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/cdndistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnDistributionRoutingRulesUpstreamOutputReference interface {
	cdktf.ComplexObject
	Caching() interface{}
	SetCaching(val interface{})
	CachingInput() interface{}
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
	GeoRestrictions() CdnDistributionRoutingRulesUpstreamGeoRestrictionsOutputReference
	GeoRestrictionsInput() *CdnDistributionRoutingRulesUpstreamGeoRestrictions
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *CdnDistributionRoutingRulesUpstream
	SetInternalValue(val *CdnDistributionRoutingRulesUpstream)
	RateLimitClass() *string
	SetRateLimitClass(val *string)
	RateLimitClassInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Waf() interface{}
	SetWaf(val interface{})
	WafInput() interface{}
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
	PutGeoRestrictions(value *CdnDistributionRoutingRulesUpstreamGeoRestrictions)
	ResetGeoRestrictions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnDistributionRoutingRulesUpstreamOutputReference
type jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) Caching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) CachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GeoRestrictions() CdnDistributionRoutingRulesUpstreamGeoRestrictionsOutputReference {
	var returns CdnDistributionRoutingRulesUpstreamGeoRestrictionsOutputReference
	_jsii_.Get(
		j,
		"geoRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GeoRestrictionsInput() *CdnDistributionRoutingRulesUpstreamGeoRestrictions {
	var returns *CdnDistributionRoutingRulesUpstreamGeoRestrictions
	_jsii_.Get(
		j,
		"geoRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) InternalValue() *CdnDistributionRoutingRulesUpstream {
	var returns *CdnDistributionRoutingRulesUpstream
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) RateLimitClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) RateLimitClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) Waf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) WafInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafInput",
		&returns,
	)
	return returns
}


func NewCdnDistributionRoutingRulesUpstreamOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnDistributionRoutingRulesUpstreamOutputReference {
	_init_.Initialize()

	if err := validateNewCdnDistributionRoutingRulesUpstreamOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cdnDistribution.CdnDistributionRoutingRulesUpstreamOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnDistributionRoutingRulesUpstreamOutputReference_Override(c CdnDistributionRoutingRulesUpstreamOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cdnDistribution.CdnDistributionRoutingRulesUpstreamOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetCaching(val interface{}) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetInternalValue(val *CdnDistributionRoutingRulesUpstream) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetRateLimitClass(val *string) {
	if err := j.validateSetRateLimitClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitClass",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference)SetWaf(val interface{}) {
	if err := j.validateSetWafParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waf",
		val,
	)
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) PutGeoRestrictions(value *CdnDistributionRoutingRulesUpstreamGeoRestrictions) {
	if err := c.validatePutGeoRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGeoRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) ResetGeoRestrictions() {
	_jsii_.InvokeVoid(
		c,
		"resetGeoRestrictions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionRoutingRulesUpstreamOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

