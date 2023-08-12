package group

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v8/group/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/group ionoscloud_group}.
type Group interface {
	cdktf.TerraformResource
	AccessActivityLog() interface{}
	SetAccessActivityLog(val interface{})
	AccessActivityLogInput() interface{}
	AccessAndManageCertificates() interface{}
	SetAccessAndManageCertificates(val interface{})
	AccessAndManageCertificatesInput() interface{}
	AccessAndManageMonitoring() interface{}
	SetAccessAndManageMonitoring(val interface{})
	AccessAndManageMonitoringInput() interface{}
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
	CreateBackupUnit() interface{}
	SetCreateBackupUnit(val interface{})
	CreateBackupUnitInput() interface{}
	CreateDatacenter() interface{}
	SetCreateDatacenter(val interface{})
	CreateDatacenterInput() interface{}
	CreateFlowLog() interface{}
	SetCreateFlowLog(val interface{})
	CreateFlowLogInput() interface{}
	CreateInternetAccess() interface{}
	SetCreateInternetAccess(val interface{})
	CreateInternetAccessInput() interface{}
	CreateK8SCluster() interface{}
	SetCreateK8SCluster(val interface{})
	CreateK8SClusterInput() interface{}
	CreatePcc() interface{}
	SetCreatePcc(val interface{})
	CreatePccInput() interface{}
	CreateSnapshot() interface{}
	SetCreateSnapshot(val interface{})
	CreateSnapshotInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageDbaas() interface{}
	SetManageDbaas(val interface{})
	ManageDbaasInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	ReserveIp() interface{}
	SetReserveIp(val interface{})
	ReserveIpInput() interface{}
	S3Privilege() interface{}
	SetS3Privilege(val interface{})
	S3PrivilegeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	UserIds() *[]*string
	SetUserIds(val *[]*string)
	UserIdsInput() *[]*string
	Users() GroupUsersList
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
	PutTimeouts(value *GroupTimeouts)
	ResetAccessActivityLog()
	ResetAccessAndManageCertificates()
	ResetAccessAndManageMonitoring()
	ResetCreateBackupUnit()
	ResetCreateDatacenter()
	ResetCreateFlowLog()
	ResetCreateInternetAccess()
	ResetCreateK8SCluster()
	ResetCreatePcc()
	ResetCreateSnapshot()
	ResetId()
	ResetManageDbaas()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReserveIp()
	ResetS3Privilege()
	ResetTimeouts()
	ResetUserId()
	ResetUserIds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Group
type jsiiProxy_Group struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Group) AccessActivityLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessActivityLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AccessActivityLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessActivityLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AccessAndManageCertificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAndManageCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AccessAndManageCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAndManageCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AccessAndManageMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAndManageMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AccessAndManageMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAndManageMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateBackupUnit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createBackupUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateBackupUnitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createBackupUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateDatacenter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateDatacenterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateFlowLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createFlowLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateFlowLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createFlowLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateK8SCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createK8SCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateK8SClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createK8SClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreatePcc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPcc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreatePccInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CreateSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ManageDbaas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageDbaas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ManageDbaasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageDbaasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ReserveIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reserveIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ReserveIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reserveIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) S3Privilege() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Privilege",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) S3PrivilegeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3PrivilegeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Timeouts() GroupTimeoutsOutputReference {
	var returns GroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) UserIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) UserIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Users() GroupUsersList {
	var returns GroupUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/group ionoscloud_group} Resource.
func NewGroup(scope constructs.Construct, id *string, config *GroupConfig) Group {
	_init_.Initialize()

	if err := validateNewGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Group{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.group.Group",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.6/docs/resources/group ionoscloud_group} Resource.
func NewGroup_Override(g Group, scope constructs.Construct, id *string, config *GroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.group.Group",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Group)SetAccessActivityLog(val interface{}) {
	if err := j.validateSetAccessActivityLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessActivityLog",
		val,
	)
}

func (j *jsiiProxy_Group)SetAccessAndManageCertificates(val interface{}) {
	if err := j.validateSetAccessAndManageCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessAndManageCertificates",
		val,
	)
}

func (j *jsiiProxy_Group)SetAccessAndManageMonitoring(val interface{}) {
	if err := j.validateSetAccessAndManageMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessAndManageMonitoring",
		val,
	)
}

func (j *jsiiProxy_Group)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Group)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateBackupUnit(val interface{}) {
	if err := j.validateSetCreateBackupUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createBackupUnit",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateDatacenter(val interface{}) {
	if err := j.validateSetCreateDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatacenter",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateFlowLog(val interface{}) {
	if err := j.validateSetCreateFlowLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createFlowLog",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateInternetAccess(val interface{}) {
	if err := j.validateSetCreateInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createInternetAccess",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateK8SCluster(val interface{}) {
	if err := j.validateSetCreateK8SClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createK8SCluster",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreatePcc(val interface{}) {
	if err := j.validateSetCreatePccParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPcc",
		val,
	)
}

func (j *jsiiProxy_Group)SetCreateSnapshot(val interface{}) {
	if err := j.validateSetCreateSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSnapshot",
		val,
	)
}

func (j *jsiiProxy_Group)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Group)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Group)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Group)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Group)SetManageDbaas(val interface{}) {
	if err := j.validateSetManageDbaasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageDbaas",
		val,
	)
}

func (j *jsiiProxy_Group)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Group)SetReserveIp(val interface{}) {
	if err := j.validateSetReserveIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reserveIp",
		val,
	)
}

func (j *jsiiProxy_Group)SetS3Privilege(val interface{}) {
	if err := j.validateSetS3PrivilegeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Privilege",
		val,
	)
}

func (j *jsiiProxy_Group)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (j *jsiiProxy_Group)SetUserIds(val *[]*string) {
	if err := j.validateSetUserIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIds",
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
func Group_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.group.Group",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.group.Group",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.group.Group",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Group_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.group.Group",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Group) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Group) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Group) PutTimeouts(value *GroupTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Group) ResetAccessActivityLog() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessActivityLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAccessAndManageCertificates() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessAndManageCertificates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAccessAndManageMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessAndManageMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateBackupUnit() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateBackupUnit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateDatacenter() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateDatacenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateFlowLog() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateFlowLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateInternetAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateInternetAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateK8SCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateK8SCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreatePcc() {
	_jsii_.InvokeVoid(
		g,
		"resetCreatePcc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetCreateSnapshot() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateSnapshot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetManageDbaas() {
	_jsii_.InvokeVoid(
		g,
		"resetManageDbaas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetReserveIp() {
	_jsii_.InvokeVoid(
		g,
		"resetReserveIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetS3Privilege() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Privilege",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetUserId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetUserIds() {
	_jsii_.InvokeVoid(
		g,
		"resetUserIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

