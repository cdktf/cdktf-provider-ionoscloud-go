package pgcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v3/pgcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster ionoscloud_pg_cluster}.
type PgCluster interface {
	cdktf.TerraformResource
	BackupLocation() *string
	SetBackupLocation(val *string)
	BackupLocationInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Connections() PgClusterConnectionsOutputReference
	ConnectionsInput() *PgClusterConnections
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Credentials() PgClusterCredentialsOutputReference
	CredentialsInput() *PgClusterCredentials
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FromBackup() PgClusterFromBackupOutputReference
	FromBackupInput() *PgClusterFromBackup
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
	MaintenanceWindow() PgClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *PgClusterMaintenanceWindow
	// The tree node.
	Node() constructs.Node
	PostgresVersion() *string
	SetPostgresVersion(val *string)
	PostgresVersionInput() *string
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
	StorageSize() *float64
	SetStorageSize(val *float64)
	StorageSizeInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	SynchronizationMode() *string
	SetSynchronizationMode(val *string)
	SynchronizationModeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PgClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConnections(value *PgClusterConnections)
	PutCredentials(value *PgClusterCredentials)
	PutFromBackup(value *PgClusterFromBackup)
	PutMaintenanceWindow(value *PgClusterMaintenanceWindow)
	PutTimeouts(value *PgClusterTimeouts)
	ResetBackupLocation()
	ResetConnections()
	ResetFromBackup()
	ResetId()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PgCluster
type jsiiProxy_PgCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PgCluster) BackupLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) BackupLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Connections() PgClusterConnectionsOutputReference {
	var returns PgClusterConnectionsOutputReference
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) ConnectionsInput() *PgClusterConnections {
	var returns *PgClusterConnections
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Credentials() PgClusterCredentialsOutputReference {
	var returns PgClusterCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) CredentialsInput() *PgClusterCredentials {
	var returns *PgClusterCredentials
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) FromBackup() PgClusterFromBackupOutputReference {
	var returns PgClusterFromBackupOutputReference
	_jsii_.Get(
		j,
		"fromBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) FromBackupInput() *PgClusterFromBackup {
	var returns *PgClusterFromBackup
	_jsii_.Get(
		j,
		"fromBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) MaintenanceWindow() PgClusterMaintenanceWindowOutputReference {
	var returns PgClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) MaintenanceWindowInput() *PgClusterMaintenanceWindow {
	var returns *PgClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) PostgresVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) PostgresVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Ram() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) RamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) StorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) StorageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) SynchronizationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronizationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) SynchronizationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronizationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) Timeouts() PgClusterTimeoutsOutputReference {
	var returns PgClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster ionoscloud_pg_cluster} Resource.
func NewPgCluster(scope constructs.Construct, id *string, config *PgClusterConfig) PgCluster {
	_init_.Initialize()

	if err := validateNewPgClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PgCluster{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.pgCluster.PgCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/ionoscloud/r/pg_cluster ionoscloud_pg_cluster} Resource.
func NewPgCluster_Override(p PgCluster, scope constructs.Construct, id *string, config *PgClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.pgCluster.PgCluster",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PgCluster)SetBackupLocation(val *string) {
	if err := j.validateSetBackupLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupLocation",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetCores(val *float64) {
	if err := j.validateSetCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetInstances(val *float64) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetPostgresVersion(val *string) {
	if err := j.validateSetPostgresVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresVersion",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetRam(val *float64) {
	if err := j.validateSetRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ram",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetStorageSize(val *float64) {
	if err := j.validateSetStorageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSize",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_PgCluster)SetSynchronizationMode(val *string) {
	if err := j.validateSetSynchronizationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synchronizationMode",
		val,
	)
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
func PgCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePgCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.pgCluster.PgCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PgCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.pgCluster.PgCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PgCluster) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PgCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PgCluster) PutConnections(value *PgClusterConnections) {
	if err := p.validatePutConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConnections",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PgCluster) PutCredentials(value *PgClusterCredentials) {
	if err := p.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCredentials",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PgCluster) PutFromBackup(value *PgClusterFromBackup) {
	if err := p.validatePutFromBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFromBackup",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PgCluster) PutMaintenanceWindow(value *PgClusterMaintenanceWindow) {
	if err := p.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PgCluster) PutTimeouts(value *PgClusterTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PgCluster) ResetBackupLocation() {
	_jsii_.InvokeVoid(
		p,
		"resetBackupLocation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetConnections() {
	_jsii_.InvokeVoid(
		p,
		"resetConnections",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetFromBackup() {
	_jsii_.InvokeVoid(
		p,
		"resetFromBackup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

