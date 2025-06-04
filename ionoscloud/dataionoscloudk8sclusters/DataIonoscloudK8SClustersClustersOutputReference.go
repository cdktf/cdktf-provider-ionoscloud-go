// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionoscloudk8sclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/dataionoscloudk8sclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIonoscloudK8SClustersClustersOutputReference interface {
	cdktf.ComplexObject
	ApiSubnetAllowList() *[]*string
	AvailableUpgradeVersions() *[]*string
	CaCrt() *string
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
	Config() DataIonoscloudK8SClustersClustersConfigList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataIonoscloudK8SClustersClusters
	SetInternalValue(val *DataIonoscloudK8SClustersClusters)
	K8SVersion() *string
	KubeConfig() *string
	Location() *string
	MaintenanceWindow() DataIonoscloudK8SClustersClustersMaintenanceWindowList
	Name() *string
	NatGatewayIp() *string
	NodePools() *[]*string
	NodeSubnet() *string
	Public() cdktf.IResolvable
	S3Buckets() DataIonoscloudK8SClustersClustersS3BucketsList
	Server() *string
	State() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserTokens() cdktf.StringMap
	ViableNodePoolVersions() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataIonoscloudK8SClustersClustersOutputReference
type jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ApiSubnetAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiSubnetAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) CaCrt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCrt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Config() DataIonoscloudK8SClustersClustersConfigList {
	var returns DataIonoscloudK8SClustersClustersConfigList
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) InternalValue() *DataIonoscloudK8SClustersClusters {
	var returns *DataIonoscloudK8SClustersClusters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) K8SVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) KubeConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) MaintenanceWindow() DataIonoscloudK8SClustersClustersMaintenanceWindowList {
	var returns DataIonoscloudK8SClustersClustersMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) NatGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) NodePools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) NodeSubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Public() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) S3Buckets() DataIonoscloudK8SClustersClustersS3BucketsList {
	var returns DataIonoscloudK8SClustersClustersS3BucketsList
	_jsii_.Get(
		j,
		"s3Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) UserTokens() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"userTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ViableNodePoolVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"viableNodePoolVersions",
		&returns,
	)
	return returns
}


func NewDataIonoscloudK8SClustersClustersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataIonoscloudK8SClustersClustersOutputReference {
	_init_.Initialize()

	if err := validateNewDataIonoscloudK8SClustersClustersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudK8SClusters.DataIonoscloudK8SClustersClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataIonoscloudK8SClustersClustersOutputReference_Override(d DataIonoscloudK8SClustersClustersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudK8SClusters.DataIonoscloudK8SClustersClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference)SetInternalValue(val *DataIonoscloudK8SClustersClusters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataIonoscloudK8SClustersClustersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

