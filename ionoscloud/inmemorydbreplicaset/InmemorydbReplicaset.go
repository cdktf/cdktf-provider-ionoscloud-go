// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inmemorydbreplicaset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/inmemorydbreplicaset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/inmemorydb_replicaset ionoscloud_inmemorydb_replicaset}.
type InmemorydbReplicaset interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Connections() InmemorydbReplicasetConnectionsOutputReference
	ConnectionsInput() *InmemorydbReplicasetConnections
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Credentials() InmemorydbReplicasetCredentialsOutputReference
	CredentialsInput() *InmemorydbReplicasetCredentials
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DnsName() *string
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
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
	InitialSnapshotId() *string
	SetInitialSnapshotId(val *string)
	InitialSnapshotIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() InmemorydbReplicasetMaintenanceWindowOutputReference
	MaintenanceWindowInput() *InmemorydbReplicasetMaintenanceWindow
	// The tree node.
	Node() constructs.Node
	PersistenceMode() *string
	SetPersistenceMode(val *string)
	PersistenceModeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Replicas() *float64
	SetReplicas(val *float64)
	ReplicasInput() *float64
	Resources() InmemorydbReplicasetResourcesOutputReference
	ResourcesInput() *InmemorydbReplicasetResources
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() InmemorydbReplicasetTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutConnections(value *InmemorydbReplicasetConnections)
	PutCredentials(value *InmemorydbReplicasetCredentials)
	PutMaintenanceWindow(value *InmemorydbReplicasetMaintenanceWindow)
	PutResources(value *InmemorydbReplicasetResources)
	PutTimeouts(value *InmemorydbReplicasetTimeouts)
	ResetId()
	ResetInitialSnapshotId()
	ResetLocation()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for InmemorydbReplicaset
type jsiiProxy_InmemorydbReplicaset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_InmemorydbReplicaset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Connections() InmemorydbReplicasetConnectionsOutputReference {
	var returns InmemorydbReplicasetConnectionsOutputReference
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) ConnectionsInput() *InmemorydbReplicasetConnections {
	var returns *InmemorydbReplicasetConnections
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Credentials() InmemorydbReplicasetCredentialsOutputReference {
	var returns InmemorydbReplicasetCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) CredentialsInput() *InmemorydbReplicasetCredentials {
	var returns *InmemorydbReplicasetCredentials
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) InitialSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) InitialSnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialSnapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) MaintenanceWindow() InmemorydbReplicasetMaintenanceWindowOutputReference {
	var returns InmemorydbReplicasetMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) MaintenanceWindowInput() *InmemorydbReplicasetMaintenanceWindow {
	var returns *InmemorydbReplicasetMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) PersistenceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) PersistenceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) ReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Resources() InmemorydbReplicasetResourcesOutputReference {
	var returns InmemorydbReplicasetResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) ResourcesInput() *InmemorydbReplicasetResources {
	var returns *InmemorydbReplicasetResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Timeouts() InmemorydbReplicasetTimeoutsOutputReference {
	var returns InmemorydbReplicasetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InmemorydbReplicaset) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/inmemorydb_replicaset ionoscloud_inmemorydb_replicaset} Resource.
func NewInmemorydbReplicaset(scope constructs.Construct, id *string, config *InmemorydbReplicasetConfig) InmemorydbReplicaset {
	_init_.Initialize()

	if err := validateNewInmemorydbReplicasetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InmemorydbReplicaset{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.2/docs/resources/inmemorydb_replicaset ionoscloud_inmemorydb_replicaset} Resource.
func NewInmemorydbReplicaset_Override(i InmemorydbReplicaset, scope constructs.Construct, id *string, config *InmemorydbReplicasetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetInitialSnapshotId(val *string) {
	if err := j.validateSetInitialSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialSnapshotId",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetPersistenceMode(val *string) {
	if err := j.validateSetPersistenceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistenceMode",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetReplicas(val *float64) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_InmemorydbReplicaset)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a InmemorydbReplicaset resource upon running "cdktf plan <stack-name>".
func InmemorydbReplicaset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateInmemorydbReplicaset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
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
func InmemorydbReplicaset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInmemorydbReplicaset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InmemorydbReplicaset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInmemorydbReplicaset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InmemorydbReplicaset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInmemorydbReplicaset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InmemorydbReplicaset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.inmemorydbReplicaset.InmemorydbReplicaset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) PutConnections(value *InmemorydbReplicasetConnections) {
	if err := i.validatePutConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putConnections",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) PutCredentials(value *InmemorydbReplicasetCredentials) {
	if err := i.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCredentials",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) PutMaintenanceWindow(value *InmemorydbReplicasetMaintenanceWindow) {
	if err := i.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) PutResources(value *InmemorydbReplicasetResources) {
	if err := i.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResources",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) PutTimeouts(value *InmemorydbReplicasetTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetInitialSnapshotId() {
	_jsii_.InvokeVoid(
		i,
		"resetInitialSnapshotId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetLocation() {
	_jsii_.InvokeVoid(
		i,
		"resetLocation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		i,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InmemorydbReplicaset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InmemorydbReplicaset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

