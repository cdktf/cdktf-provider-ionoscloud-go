// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cubeserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/cubeserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/cube_server ionoscloud_cube_server}.
type CubeServer interface {
	cdktf.TerraformResource
	AllowReplace() interface{}
	SetAllowReplace(val interface{})
	AllowReplaceInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BootCdrom() *string
	SetBootCdrom(val *string)
	BootCdromInput() *string
	BootImage() *string
	SetBootImage(val *string)
	BootImageInput() *string
	BootVolume() *string
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
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FirewallruleId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	ImagePassword() *string
	SetImagePassword(val *string)
	ImagePasswordInput() *string
	InlineVolumeIds() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nic() CubeServerNicOutputReference
	NicInput() *CubeServerNic
	// The tree node.
	Node() constructs.Node
	PrimaryIp() *string
	PrimaryNic() *string
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
	SecurityGroupsIds() *[]*string
	SetSecurityGroupsIds(val *[]*string)
	SecurityGroupsIdsInput() *[]*string
	SshKeyPath() *[]*string
	SetSshKeyPath(val *[]*string)
	SshKeyPathInput() *[]*string
	TemplateUuid() *string
	SetTemplateUuid(val *string)
	TemplateUuidInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CubeServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	VmState() *string
	SetVmState(val *string)
	VmStateInput() *string
	Volume() CubeServerVolumeOutputReference
	VolumeInput() *CubeServerVolume
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
	PutNic(value *CubeServerNic)
	PutTimeouts(value *CubeServerTimeouts)
	PutVolume(value *CubeServerVolume)
	ResetAllowReplace()
	ResetAvailabilityZone()
	ResetBootCdrom()
	ResetBootImage()
	ResetHostname()
	ResetId()
	ResetImageName()
	ResetImagePassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecurityGroupsIds()
	ResetSshKeyPath()
	ResetTimeouts()
	ResetVmState()
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

// The jsii proxy struct for CubeServer
type jsiiProxy_CubeServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CubeServer) AllowReplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) AllowReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) BootCdrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) BootCdromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) BootImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) BootImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) BootVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) FirewallruleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) ImagePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) InlineVolumeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inlineVolumeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Nic() CubeServerNicOutputReference {
	var returns CubeServerNicOutputReference
	_jsii_.Get(
		j,
		"nic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) NicInput() *CubeServerNic {
	var returns *CubeServerNic
	_jsii_.Get(
		j,
		"nicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) PrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) PrimaryNic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) SecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) SecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) SshKeyPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) SshKeyPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TemplateUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TemplateUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Timeouts() CubeServerTimeoutsOutputReference {
	var returns CubeServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) VmState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) VmStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) Volume() CubeServerVolumeOutputReference {
	var returns CubeServerVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServer) VolumeInput() *CubeServerVolume {
	var returns *CubeServerVolume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/cube_server ionoscloud_cube_server} Resource.
func NewCubeServer(scope constructs.Construct, id *string, config *CubeServerConfig) CubeServer {
	_init_.Initialize()

	if err := validateNewCubeServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CubeServer{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/cube_server ionoscloud_cube_server} Resource.
func NewCubeServer_Override(c CubeServer, scope constructs.Construct, id *string, config *CubeServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CubeServer)SetAllowReplace(val interface{}) {
	if err := j.validateSetAllowReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplace",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetBootCdrom(val *string) {
	if err := j.validateSetBootCdromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootCdrom",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetBootImage(val *string) {
	if err := j.validateSetBootImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImage",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetImagePassword(val *string) {
	if err := j.validateSetImagePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePassword",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetSshKeyPath(val *[]*string) {
	if err := j.validateSetSshKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeyPath",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetTemplateUuid(val *string) {
	if err := j.validateSetTemplateUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUuid",
		val,
	)
}

func (j *jsiiProxy_CubeServer)SetVmState(val *string) {
	if err := j.validateSetVmStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmState",
		val,
	)
}

// Generates CDKTF code for importing a CubeServer resource upon running "cdktf plan <stack-name>".
func CubeServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCubeServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
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
func CubeServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCubeServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CubeServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCubeServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CubeServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCubeServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CubeServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CubeServer) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CubeServer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CubeServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CubeServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CubeServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CubeServer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CubeServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CubeServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CubeServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CubeServer) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CubeServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CubeServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CubeServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CubeServer) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CubeServer) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CubeServer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CubeServer) PutNic(value *CubeServerNic) {
	if err := c.validatePutNicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNic",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CubeServer) PutTimeouts(value *CubeServerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CubeServer) PutVolume(value *CubeServerVolume) {
	if err := c.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CubeServer) ResetAllowReplace() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowReplace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetBootCdrom() {
	_jsii_.InvokeVoid(
		c,
		"resetBootCdrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetBootImage() {
	_jsii_.InvokeVoid(
		c,
		"resetBootImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetImageName() {
	_jsii_.InvokeVoid(
		c,
		"resetImageName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetImagePassword() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetSecurityGroupsIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupsIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetSshKeyPath() {
	_jsii_.InvokeVoid(
		c,
		"resetSshKeyPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) ResetVmState() {
	_jsii_.InvokeVoid(
		c,
		"resetVmState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

