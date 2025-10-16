// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/server/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/server ionoscloud_server}.
type Server interface {
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
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuFamily() *string
	SetCpuFamily(val *string)
	CpuFamilyInput() *string
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FirewallruleId() *string
	FirewallruleIds() *[]*string
	SetFirewallruleIds(val *[]*string)
	FirewallruleIdsInput() *[]*string
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
	Label() ServerLabelList
	LabelInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nic() ServerNicOutputReference
	NicInput() *ServerNic
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
	Ram() *float64
	SetRam(val *float64)
	RamInput() *float64
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupsIds() *[]*string
	SetSecurityGroupsIds(val *[]*string)
	SecurityGroupsIdsInput() *[]*string
	SshKeyPath() *[]*string
	SetSshKeyPath(val *[]*string)
	SshKeyPathInput() *[]*string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	TemplateUuid() *string
	SetTemplateUuid(val *string)
	TemplateUuidInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VmState() *string
	SetVmState(val *string)
	VmStateInput() *string
	Volume() ServerVolumeOutputReference
	VolumeInput() *ServerVolume
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
	PutLabel(value interface{})
	PutNic(value *ServerNic)
	PutTimeouts(value *ServerTimeouts)
	PutVolume(value *ServerVolume)
	ResetAllowReplace()
	ResetAvailabilityZone()
	ResetBootCdrom()
	ResetBootImage()
	ResetCores()
	ResetCpuFamily()
	ResetFirewallruleIds()
	ResetHostname()
	ResetId()
	ResetImageName()
	ResetImagePassword()
	ResetLabel()
	ResetNic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRam()
	ResetSecurityGroupsIds()
	ResetSshKeyPath()
	ResetSshKeys()
	ResetTemplateUuid()
	ResetTimeouts()
	ResetType()
	ResetVmState()
	ResetVolume()
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

// The jsii proxy struct for Server
type jsiiProxy_Server struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Server) AllowReplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) AllowReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootCdrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootCdromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CpuFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CpuFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FirewallruleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FirewallruleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewallruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FirewallruleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewallruleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ImagePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) InlineVolumeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inlineVolumeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Label() ServerLabelList {
	var returns ServerLabelList
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) LabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Nic() ServerNicOutputReference {
	var returns ServerNicOutputReference
	_jsii_.Get(
		j,
		"nic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NicInput() *ServerNic {
	var returns *ServerNic
	_jsii_.Get(
		j,
		"nicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) PrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) PrimaryNic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Ram() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) RamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SshKeyPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SshKeyPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TemplateUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TemplateUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Timeouts() ServerTimeoutsOutputReference {
	var returns ServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) VmState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) VmStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Volume() ServerVolumeOutputReference {
	var returns ServerVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) VolumeInput() *ServerVolume {
	var returns *ServerVolume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/server ionoscloud_server} Resource.
func NewServer(scope constructs.Construct, id *string, config *ServerConfig) Server {
	_init_.Initialize()

	if err := validateNewServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Server{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.server.Server",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/server ionoscloud_server} Resource.
func NewServer_Override(s Server, scope constructs.Construct, id *string, config *ServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.server.Server",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Server)SetAllowReplace(val interface{}) {
	if err := j.validateSetAllowReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplace",
		val,
	)
}

func (j *jsiiProxy_Server)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Server)SetBootCdrom(val *string) {
	if err := j.validateSetBootCdromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootCdrom",
		val,
	)
}

func (j *jsiiProxy_Server)SetBootImage(val *string) {
	if err := j.validateSetBootImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImage",
		val,
	)
}

func (j *jsiiProxy_Server)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Server)SetCores(val *float64) {
	if err := j.validateSetCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_Server)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Server)SetCpuFamily(val *string) {
	if err := j.validateSetCpuFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuFamily",
		val,
	)
}

func (j *jsiiProxy_Server)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_Server)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Server)SetFirewallruleIds(val *[]*string) {
	if err := j.validateSetFirewallruleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallruleIds",
		val,
	)
}

func (j *jsiiProxy_Server)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Server)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_Server)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Server)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_Server)SetImagePassword(val *string) {
	if err := j.validateSetImagePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePassword",
		val,
	)
}

func (j *jsiiProxy_Server)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Server)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Server)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Server)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Server)SetRam(val *float64) {
	if err := j.validateSetRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ram",
		val,
	)
}

func (j *jsiiProxy_Server)SetSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_Server)SetSshKeyPath(val *[]*string) {
	if err := j.validateSetSshKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeyPath",
		val,
	)
}

func (j *jsiiProxy_Server)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_Server)SetTemplateUuid(val *string) {
	if err := j.validateSetTemplateUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUuid",
		val,
	)
}

func (j *jsiiProxy_Server)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Server)SetVmState(val *string) {
	if err := j.validateSetVmStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmState",
		val,
	)
}

// Generates CDKTF code for importing a Server resource upon running "cdktf plan <stack-name>".
func Server_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.server.Server",
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
func Server_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.server.Server",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Server_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.server.Server",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Server_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.server.Server",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Server_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.server.Server",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Server) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Server) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Server) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Server) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Server) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Server) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Server) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Server) PutLabel(value interface{}) {
	if err := s.validatePutLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLabel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutNic(value *ServerNic) {
	if err := s.validatePutNicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNic",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutTimeouts(value *ServerTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutVolume(value *ServerVolume) {
	if err := s.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) ResetAllowReplace() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowReplace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		s,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetBootCdrom() {
	_jsii_.InvokeVoid(
		s,
		"resetBootCdrom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetBootImage() {
	_jsii_.InvokeVoid(
		s,
		"resetBootImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetCores() {
	_jsii_.InvokeVoid(
		s,
		"resetCores",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetCpuFamily() {
	_jsii_.InvokeVoid(
		s,
		"resetCpuFamily",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetFirewallruleIds() {
	_jsii_.InvokeVoid(
		s,
		"resetFirewallruleIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetHostname() {
	_jsii_.InvokeVoid(
		s,
		"resetHostname",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetImageName() {
	_jsii_.InvokeVoid(
		s,
		"resetImageName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetImagePassword() {
	_jsii_.InvokeVoid(
		s,
		"resetImagePassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetNic() {
	_jsii_.InvokeVoid(
		s,
		"resetNic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetRam() {
	_jsii_.InvokeVoid(
		s,
		"resetRam",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetSecurityGroupsIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroupsIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetSshKeyPath() {
	_jsii_.InvokeVoid(
		s,
		"resetSshKeyPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetSshKeys() {
	_jsii_.InvokeVoid(
		s,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTemplateUuid() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplateUuid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetVmState() {
	_jsii_.InvokeVoid(
		s,
		"resetVmState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

