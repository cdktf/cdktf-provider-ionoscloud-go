// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongocluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/mongocluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/mongo_cluster ionoscloud_mongo_cluster}.
type MongoCluster interface {
	cdktf.TerraformResource
	Backup() MongoClusterBackupOutputReference
	BackupInput() *MongoClusterBackup
	BiConnector() MongoClusterBiConnectorOutputReference
	BiConnectorInput() *MongoClusterBiConnector
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Connections() MongoClusterConnectionsOutputReference
	ConnectionsInput() *MongoClusterConnections
	ConnectionString() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
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
	Instances() *float64
	SetInstances(val *float64)
	InstancesInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() MongoClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *MongoClusterMaintenanceWindow
	MongodbVersion() *string
	SetMongodbVersion(val *string)
	MongodbVersionInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Ram() *float64
	SetRam(val *float64)
	RamInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Shards() *float64
	SetShards(val *float64)
	ShardsInput() *float64
	StorageSize() *float64
	SetStorageSize(val *float64)
	StorageSizeInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	TemplateId() *string
	SetTemplateId(val *string)
	TemplateIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MongoClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutBackup(value *MongoClusterBackup)
	PutBiConnector(value *MongoClusterBiConnector)
	PutConnections(value *MongoClusterConnections)
	PutMaintenanceWindow(value *MongoClusterMaintenanceWindow)
	PutTimeouts(value *MongoClusterTimeouts)
	ResetBackup()
	ResetBiConnector()
	ResetCores()
	ResetEdition()
	ResetId()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRam()
	ResetShards()
	ResetStorageSize()
	ResetStorageType()
	ResetTemplateId()
	ResetTimeouts()
	ResetType()
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

// The jsii proxy struct for MongoCluster
type jsiiProxy_MongoCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MongoCluster) Backup() MongoClusterBackupOutputReference {
	var returns MongoClusterBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) BackupInput() *MongoClusterBackup {
	var returns *MongoClusterBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) BiConnector() MongoClusterBiConnectorOutputReference {
	var returns MongoClusterBiConnectorOutputReference
	_jsii_.Get(
		j,
		"biConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) BiConnectorInput() *MongoClusterBiConnector {
	var returns *MongoClusterBiConnector
	_jsii_.Get(
		j,
		"biConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Connections() MongoClusterConnectionsOutputReference {
	var returns MongoClusterConnectionsOutputReference
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) ConnectionsInput() *MongoClusterConnections {
	var returns *MongoClusterConnections
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) MaintenanceWindow() MongoClusterMaintenanceWindowOutputReference {
	var returns MongoClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) MaintenanceWindowInput() *MongoClusterMaintenanceWindow {
	var returns *MongoClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) MongodbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) MongodbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Ram() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) RamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Shards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) ShardsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) StorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) StorageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Timeouts() MongoClusterTimeoutsOutputReference {
	var returns MongoClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoCluster) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/mongo_cluster ionoscloud_mongo_cluster} Resource.
func NewMongoCluster(scope constructs.Construct, id *string, config *MongoClusterConfig) MongoCluster {
	_init_.Initialize()

	if err := validateNewMongoClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongoCluster{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.17/docs/resources/mongo_cluster ionoscloud_mongo_cluster} Resource.
func NewMongoCluster_Override(m MongoCluster, scope constructs.Construct, id *string, config *MongoClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MongoCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetCores(val *float64) {
	if err := j.validateSetCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetInstances(val *float64) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetMongodbVersion(val *string) {
	if err := j.validateSetMongodbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongodbVersion",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetRam(val *float64) {
	if err := j.validateSetRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ram",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetShards(val *float64) {
	if err := j.validateSetShardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shards",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetStorageSize(val *float64) {
	if err := j.validateSetStorageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSize",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetTemplateId(val *string) {
	if err := j.validateSetTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

func (j *jsiiProxy_MongoCluster)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a MongoCluster resource upon running "cdktf plan <stack-name>".
func MongoCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMongoCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
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
func MongoCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongoCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MongoCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongoCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MongoCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongoCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MongoCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.mongoCluster.MongoCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MongoCluster) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MongoCluster) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MongoCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MongoCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MongoCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MongoCluster) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MongoCluster) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MongoCluster) PutBackup(value *MongoClusterBackup) {
	if err := m.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBackup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MongoCluster) PutBiConnector(value *MongoClusterBiConnector) {
	if err := m.validatePutBiConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBiConnector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MongoCluster) PutConnections(value *MongoClusterConnections) {
	if err := m.validatePutConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putConnections",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MongoCluster) PutMaintenanceWindow(value *MongoClusterMaintenanceWindow) {
	if err := m.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MongoCluster) PutTimeouts(value *MongoClusterTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MongoCluster) ResetBackup() {
	_jsii_.InvokeVoid(
		m,
		"resetBackup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetBiConnector() {
	_jsii_.InvokeVoid(
		m,
		"resetBiConnector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetCores() {
	_jsii_.InvokeVoid(
		m,
		"resetCores",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetEdition() {
	_jsii_.InvokeVoid(
		m,
		"resetEdition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetRam() {
	_jsii_.InvokeVoid(
		m,
		"resetRam",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetShards() {
	_jsii_.InvokeVoid(
		m,
		"resetShards",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetStorageSize() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetStorageType() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetTemplateId() {
	_jsii_.InvokeVoid(
		m,
		"resetTemplateId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) ResetType() {
	_jsii_.InvokeVoid(
		m,
		"resetType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongoCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

