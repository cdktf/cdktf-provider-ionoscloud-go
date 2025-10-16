// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nsgfirewallrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/nsgfirewallrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule ionoscloud_nsg_firewallrule}.
type NsgFirewallrule interface {
	cdktf.TerraformResource
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IcmpCode() *string
	SetIcmpCode(val *string)
	IcmpCodeInput() *string
	IcmpType() *string
	SetIcmpType(val *string)
	IcmpTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NsgId() *string
	SetNsgId(val *string)
	NsgIdInput() *string
	PortRangeEnd() *float64
	SetPortRangeEnd(val *float64)
	PortRangeEndInput() *float64
	PortRangeStart() *float64
	SetPortRangeStart(val *float64)
	PortRangeStartInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SourceIp() *string
	SetSourceIp(val *string)
	SourceIpInput() *string
	SourceMac() *string
	SetSourceMac(val *string)
	SourceMacInput() *string
	TargetIp() *string
	SetTargetIp(val *string)
	TargetIpInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NsgFirewallruleTimeoutsOutputReference
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
	PutTimeouts(value *NsgFirewallruleTimeouts)
	ResetIcmpCode()
	ResetIcmpType()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortRangeEnd()
	ResetPortRangeStart()
	ResetSourceIp()
	ResetSourceMac()
	ResetTargetIp()
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

// The jsii proxy struct for NsgFirewallrule
type jsiiProxy_NsgFirewallrule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsgFirewallrule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) IcmpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) IcmpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) NsgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) NsgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) PortRangeEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) PortRangeEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) PortRangeStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) PortRangeStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) SourceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) SourceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) SourceMac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) SourceMacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TargetIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TargetIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Timeouts() NsgFirewallruleTimeoutsOutputReference {
	var returns NsgFirewallruleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsgFirewallrule) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule ionoscloud_nsg_firewallrule} Resource.
func NewNsgFirewallrule(scope constructs.Construct, id *string, config *NsgFirewallruleConfig) NsgFirewallrule {
	_init_.Initialize()

	if err := validateNewNsgFirewallruleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsgFirewallrule{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/nsg_firewallrule ionoscloud_nsg_firewallrule} Resource.
func NewNsgFirewallrule_Override(n NsgFirewallrule, scope constructs.Construct, id *string, config *NsgFirewallruleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetIcmpCode(val *string) {
	if err := j.validateSetIcmpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpCode",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetNsgId(val *string) {
	if err := j.validateSetNsgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsgId",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetPortRangeEnd(val *float64) {
	if err := j.validateSetPortRangeEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeEnd",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetPortRangeStart(val *float64) {
	if err := j.validateSetPortRangeStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeStart",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetSourceIp(val *string) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetSourceMac(val *string) {
	if err := j.validateSetSourceMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMac",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetTargetIp(val *string) {
	if err := j.validateSetTargetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIp",
		val,
	)
}

func (j *jsiiProxy_NsgFirewallrule)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a NsgFirewallrule resource upon running "cdktf plan <stack-name>".
func NsgFirewallrule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNsgFirewallrule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
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
func NsgFirewallrule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsgFirewallrule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsgFirewallrule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsgFirewallrule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsgFirewallrule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsgFirewallrule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsgFirewallrule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.nsgFirewallrule.NsgFirewallrule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsgFirewallrule) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NsgFirewallrule) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsgFirewallrule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NsgFirewallrule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NsgFirewallrule) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NsgFirewallrule) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NsgFirewallrule) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsgFirewallrule) PutTimeouts(value *NsgFirewallruleTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetIcmpCode() {
	_jsii_.InvokeVoid(
		n,
		"resetIcmpCode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetIcmpType() {
	_jsii_.InvokeVoid(
		n,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetName() {
	_jsii_.InvokeVoid(
		n,
		"resetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetPortRangeEnd() {
	_jsii_.InvokeVoid(
		n,
		"resetPortRangeEnd",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetPortRangeStart() {
	_jsii_.InvokeVoid(
		n,
		"resetPortRangeStart",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetSourceIp() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetSourceMac() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceMac",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetTargetIp() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsgFirewallrule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsgFirewallrule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

