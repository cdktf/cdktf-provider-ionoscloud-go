// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8scluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/k8scluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/k8s_cluster ionoscloud_k8s_cluster}.
type K8SCluster interface {
	cdktf.TerraformResource
	AllowReplace() interface{}
	SetAllowReplace(val interface{})
	AllowReplaceInput() interface{}
	ApiSubnetAllowList() *[]*string
	SetApiSubnetAllowList(val *[]*string)
	ApiSubnetAllowListInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	K8SVersion() *string
	SetK8SVersion(val *string)
	K8SVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() K8SClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *K8SClusterMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	NatGatewayIp() *string
	SetNatGatewayIp(val *string)
	NatGatewayIpInput() *string
	// The tree node.
	Node() constructs.Node
	NodeSubnet() *string
	SetNodeSubnet(val *string)
	NodeSubnetInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Public() interface{}
	SetPublic(val interface{})
	PublicInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	S3Buckets() K8SClusterS3BucketsList
	S3BucketsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() K8SClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	ViableNodePoolVersions() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutMaintenanceWindow(value *K8SClusterMaintenanceWindow)
	PutS3Buckets(value interface{})
	PutTimeouts(value *K8SClusterTimeouts)
	ResetAllowReplace()
	ResetApiSubnetAllowList()
	ResetId()
	ResetK8SVersion()
	ResetLocation()
	ResetMaintenanceWindow()
	ResetNatGatewayIp()
	ResetNodeSubnet()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublic()
	ResetS3Buckets()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for K8SCluster
type jsiiProxy_K8SCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_K8SCluster) AllowReplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) AllowReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) ApiSubnetAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiSubnetAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) ApiSubnetAllowListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiSubnetAllowListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) K8SVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) K8SVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) MaintenanceWindow() K8SClusterMaintenanceWindowOutputReference {
	var returns K8SClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) MaintenanceWindowInput() *K8SClusterMaintenanceWindow {
	var returns *K8SClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) NatGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) NatGatewayIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) NodeSubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) NodeSubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Public() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) PublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) S3Buckets() K8SClusterS3BucketsList {
	var returns K8SClusterS3BucketsList
	_jsii_.Get(
		j,
		"s3Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) S3BucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) Timeouts() K8SClusterTimeoutsOutputReference {
	var returns K8SClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SCluster) ViableNodePoolVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"viableNodePoolVersions",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/k8s_cluster ionoscloud_k8s_cluster} Resource.
func NewK8SCluster(scope constructs.Construct, id *string, config *K8SClusterConfig) K8SCluster {
	_init_.Initialize()

	if err := validateNewK8SClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_K8SCluster{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.0/docs/resources/k8s_cluster ionoscloud_k8s_cluster} Resource.
func NewK8SCluster_Override(k K8SCluster, scope constructs.Construct, id *string, config *K8SClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_K8SCluster)SetAllowReplace(val interface{}) {
	if err := j.validateSetAllowReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplace",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetApiSubnetAllowList(val *[]*string) {
	if err := j.validateSetApiSubnetAllowListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiSubnetAllowList",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetK8SVersion(val *string) {
	if err := j.validateSetK8SVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k8SVersion",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetNatGatewayIp(val *string) {
	if err := j.validateSetNatGatewayIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayIp",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetNodeSubnet(val *string) {
	if err := j.validateSetNodeSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSubnet",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_K8SCluster)SetPublic(val interface{}) {
	if err := j.validateSetPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"public",
		val,
	)
}

// Generates CDKTF code for importing a K8SCluster resource upon running "cdktf plan <stack-name>".
func K8SCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateK8SCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func K8SCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func K8SCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func K8SCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func K8SCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.k8SCluster.K8SCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_K8SCluster) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_K8SCluster) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_K8SCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_K8SCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_K8SCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_K8SCluster) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_K8SCluster) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_K8SCluster) PutMaintenanceWindow(value *K8SClusterMaintenanceWindow) {
	if err := k.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SCluster) PutS3Buckets(value interface{}) {
	if err := k.validatePutS3BucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Buckets",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SCluster) PutTimeouts(value *K8SClusterTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SCluster) ResetAllowReplace() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowReplace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetApiSubnetAllowList() {
	_jsii_.InvokeVoid(
		k,
		"resetApiSubnetAllowList",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetK8SVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetK8SVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetLocation() {
	_jsii_.InvokeVoid(
		k,
		"resetLocation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetNatGatewayIp() {
	_jsii_.InvokeVoid(
		k,
		"resetNatGatewayIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetNodeSubnet() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeSubnet",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetPublic() {
	_jsii_.InvokeVoid(
		k,
		"resetPublic",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetS3Buckets() {
	_jsii_.InvokeVoid(
		k,
		"resetS3Buckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

