// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vcpuserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/vcpuserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/vcpu_server ionoscloud_vcpu_server}.
type VcpuServer interface {
	cdktf.TerraformResource
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
	Label() VcpuServerLabelList
	LabelInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nic() VcpuServerNicOutputReference
	NicInput() *VcpuServerNic
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
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VcpuServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	VmState() *string
	SetVmState(val *string)
	VmStateInput() *string
	Volume() VcpuServerVolumeOutputReference
	VolumeInput() *VcpuServerVolume
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
	PutNic(value *VcpuServerNic)
	PutTimeouts(value *VcpuServerTimeouts)
	PutVolume(value *VcpuServerVolume)
	ResetAvailabilityZone()
	ResetBootCdrom()
	ResetBootImage()
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
	ResetSecurityGroupsIds()
	ResetSshKeys()
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

// The jsii proxy struct for VcpuServer
type jsiiProxy_VcpuServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VcpuServer) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) BootCdrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) BootCdromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootCdromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) BootImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) BootImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) BootVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) CpuFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) FirewallruleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) FirewallruleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewallruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) FirewallruleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewallruleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) ImagePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) InlineVolumeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inlineVolumeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Label() VcpuServerLabelList {
	var returns VcpuServerLabelList
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) LabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Nic() VcpuServerNicOutputReference {
	var returns VcpuServerNicOutputReference
	_jsii_.Get(
		j,
		"nic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) NicInput() *VcpuServerNic {
	var returns *VcpuServerNic
	_jsii_.Get(
		j,
		"nicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) PrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) PrimaryNic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Ram() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) RamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) SecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) SecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Timeouts() VcpuServerTimeoutsOutputReference {
	var returns VcpuServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) VmState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) VmStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) Volume() VcpuServerVolumeOutputReference {
	var returns VcpuServerVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcpuServer) VolumeInput() *VcpuServerVolume {
	var returns *VcpuServerVolume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/vcpu_server ionoscloud_vcpu_server} Resource.
func NewVcpuServer(scope constructs.Construct, id *string, config *VcpuServerConfig) VcpuServer {
	_init_.Initialize()

	if err := validateNewVcpuServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VcpuServer{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.7/docs/resources/vcpu_server ionoscloud_vcpu_server} Resource.
func NewVcpuServer_Override(v VcpuServer, scope constructs.Construct, id *string, config *VcpuServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VcpuServer)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetBootCdrom(val *string) {
	if err := j.validateSetBootCdromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootCdrom",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetBootImage(val *string) {
	if err := j.validateSetBootImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImage",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetCores(val *float64) {
	if err := j.validateSetCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetFirewallruleIds(val *[]*string) {
	if err := j.validateSetFirewallruleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallruleIds",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetImagePassword(val *string) {
	if err := j.validateSetImagePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePassword",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetRam(val *float64) {
	if err := j.validateSetRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ram",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_VcpuServer)SetVmState(val *string) {
	if err := j.validateSetVmStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmState",
		val,
	)
}

// Generates CDKTF code for importing a VcpuServer resource upon running "cdktf plan <stack-name>".
func VcpuServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVcpuServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
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
func VcpuServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcpuServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VcpuServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcpuServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VcpuServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcpuServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VcpuServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.vcpuServer.VcpuServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VcpuServer) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VcpuServer) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VcpuServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VcpuServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VcpuServer) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VcpuServer) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VcpuServer) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VcpuServer) PutLabel(value interface{}) {
	if err := v.validatePutLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putLabel",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VcpuServer) PutNic(value *VcpuServerNic) {
	if err := v.validatePutNicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNic",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VcpuServer) PutTimeouts(value *VcpuServerTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VcpuServer) PutVolume(value *VcpuServerVolume) {
	if err := v.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVolume",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VcpuServer) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		v,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetBootCdrom() {
	_jsii_.InvokeVoid(
		v,
		"resetBootCdrom",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetBootImage() {
	_jsii_.InvokeVoid(
		v,
		"resetBootImage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetFirewallruleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetFirewallruleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetHostname() {
	_jsii_.InvokeVoid(
		v,
		"resetHostname",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetImageName() {
	_jsii_.InvokeVoid(
		v,
		"resetImageName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetImagePassword() {
	_jsii_.InvokeVoid(
		v,
		"resetImagePassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetLabel() {
	_jsii_.InvokeVoid(
		v,
		"resetLabel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetNic() {
	_jsii_.InvokeVoid(
		v,
		"resetNic",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetSecurityGroupsIds() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityGroupsIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetSshKeys() {
	_jsii_.InvokeVoid(
		v,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) ResetVmState() {
	_jsii_.InvokeVoid(
		v,
		"resetVmState",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcpuServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcpuServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

