// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnwireguardgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v12/vpnwireguardgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/vpn_wireguard_gateway ionoscloud_vpn_wireguard_gateway}.
type VpnWireguardGateway interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Connections() VpnWireguardGatewayConnectionsList
	ConnectionsInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayIp() *string
	SetGatewayIp(val *string)
	GatewayIpInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterfaceIpv4Cidr() *string
	SetInterfaceIpv4Cidr(val *string)
	InterfaceIpv4CidrInput() *string
	InterfaceIpv6Cidr() *string
	SetInterfaceIpv6Cidr(val *string)
	InterfaceIpv6CidrInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenPort() *float64
	SetListenPort(val *float64)
	ListenPortInput() *float64
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() VpnWireguardGatewayMaintenanceWindowOutputReference
	MaintenanceWindowInput() *VpnWireguardGatewayMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicKey() *string
	// Experimental.
	RawOverrides() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() VpnWireguardGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutConnections(value interface{})
	PutMaintenanceWindow(value *VpnWireguardGatewayMaintenanceWindow)
	PutTimeouts(value *VpnWireguardGatewayTimeouts)
	ResetDescription()
	ResetId()
	ResetInterfaceIpv4Cidr()
	ResetInterfaceIpv6Cidr()
	ResetListenPort()
	ResetLocation()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTier()
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

// The jsii proxy struct for VpnWireguardGateway
type jsiiProxy_VpnWireguardGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnWireguardGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Connections() VpnWireguardGatewayConnectionsList {
	var returns VpnWireguardGatewayConnectionsList
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) ConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) GatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) GatewayIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) InterfaceIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) InterfaceIpv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceIpv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) InterfaceIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) InterfaceIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) ListenPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"listenPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) ListenPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"listenPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) MaintenanceWindow() VpnWireguardGatewayMaintenanceWindowOutputReference {
	var returns VpnWireguardGatewayMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) MaintenanceWindowInput() *VpnWireguardGatewayMaintenanceWindow {
	var returns *VpnWireguardGatewayMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) Timeouts() VpnWireguardGatewayTimeoutsOutputReference {
	var returns VpnWireguardGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnWireguardGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/vpn_wireguard_gateway ionoscloud_vpn_wireguard_gateway} Resource.
func NewVpnWireguardGateway(scope constructs.Construct, id *string, config *VpnWireguardGatewayConfig) VpnWireguardGateway {
	_init_.Initialize()

	if err := validateNewVpnWireguardGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnWireguardGateway{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.13/docs/resources/vpn_wireguard_gateway ionoscloud_vpn_wireguard_gateway} Resource.
func NewVpnWireguardGateway_Override(v VpnWireguardGateway, scope constructs.Construct, id *string, config *VpnWireguardGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetGatewayIp(val *string) {
	if err := j.validateSetGatewayIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIp",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetInterfaceIpv4Cidr(val *string) {
	if err := j.validateSetInterfaceIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetInterfaceIpv6Cidr(val *string) {
	if err := j.validateSetInterfaceIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetListenPort(val *float64) {
	if err := j.validateSetListenPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenPort",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnWireguardGateway)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

// Generates CDKTF code for importing a VpnWireguardGateway resource upon running "cdktf plan <stack-name>".
func VpnWireguardGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpnWireguardGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
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
func VpnWireguardGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnWireguardGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnWireguardGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnWireguardGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnWireguardGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnWireguardGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnWireguardGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.vpnWireguardGateway.VpnWireguardGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnWireguardGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnWireguardGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnWireguardGateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnWireguardGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnWireguardGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnWireguardGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnWireguardGateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnWireguardGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnWireguardGateway) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnWireguardGateway) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) PutConnections(value interface{}) {
	if err := v.validatePutConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putConnections",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) PutMaintenanceWindow(value *VpnWireguardGatewayMaintenanceWindow) {
	if err := v.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) PutTimeouts(value *VpnWireguardGatewayTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetInterfaceIpv4Cidr() {
	_jsii_.InvokeVoid(
		v,
		"resetInterfaceIpv4Cidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetInterfaceIpv6Cidr() {
	_jsii_.InvokeVoid(
		v,
		"resetInterfaceIpv6Cidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetListenPort() {
	_jsii_.InvokeVoid(
		v,
		"resetListenPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetLocation() {
	_jsii_.InvokeVoid(
		v,
		"resetLocation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		v,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetTier() {
	_jsii_.InvokeVoid(
		v,
		"resetTier",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnWireguardGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnWireguardGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

