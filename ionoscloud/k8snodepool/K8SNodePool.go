// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package k8snodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/k8snodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/k8s_node_pool ionoscloud_k8s_node_pool}.
type K8SNodePool interface {
	cdktf.TerraformResource
	AllowReplace() interface{}
	SetAllowReplace(val interface{})
	AllowReplaceInput() interface{}
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AutoScaling() K8SNodePoolAutoScalingOutputReference
	AutoScalingInput() *K8SNodePoolAutoScaling
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoresCount() *float64
	SetCoresCount(val *float64)
	CoresCountInput() *float64
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
	K8SClusterId() *string
	SetK8SClusterId(val *string)
	K8SClusterIdInput() *string
	K8SVersion() *string
	SetK8SVersion(val *string)
	K8SVersionInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Lans() K8SNodePoolLansList
	LansInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindow() K8SNodePoolMaintenanceWindowOutputReference
	MaintenanceWindowInput() *K8SNodePoolMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIps() *[]*string
	SetPublicIps(val *[]*string)
	PublicIpsInput() *[]*string
	RamSize() *float64
	SetRamSize(val *float64)
	RamSizeInput() *float64
	// Experimental.
	RawOverrides() interface{}
	StorageSize() *float64
	SetStorageSize(val *float64)
	StorageSizeInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() K8SNodePoolTimeoutsOutputReference
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
	PutAutoScaling(value *K8SNodePoolAutoScaling)
	PutLans(value interface{})
	PutMaintenanceWindow(value *K8SNodePoolMaintenanceWindow)
	PutTimeouts(value *K8SNodePoolTimeouts)
	ResetAllowReplace()
	ResetAnnotations()
	ResetAutoScaling()
	ResetId()
	ResetLabels()
	ResetLans()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicIps()
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

// The jsii proxy struct for K8SNodePool
type jsiiProxy_K8SNodePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_K8SNodePool) AllowReplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AllowReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AutoScaling() K8SNodePoolAutoScalingOutputReference {
	var returns K8SNodePoolAutoScalingOutputReference
	_jsii_.Get(
		j,
		"autoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AutoScalingInput() *K8SNodePoolAutoScaling {
	var returns *K8SNodePoolAutoScaling
	_jsii_.Get(
		j,
		"autoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) CoresCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) CoresCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) CpuFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) CpuFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) K8SClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) K8SClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) K8SVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) K8SVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"k8SVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Lans() K8SNodePoolLansList {
	var returns K8SNodePoolLansList
	_jsii_.Get(
		j,
		"lans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) LansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) MaintenanceWindow() K8SNodePoolMaintenanceWindowOutputReference {
	var returns K8SNodePoolMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) MaintenanceWindowInput() *K8SNodePoolMaintenanceWindow {
	var returns *K8SNodePoolMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) PublicIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) PublicIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) RamSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) RamSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) StorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) StorageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) Timeouts() K8SNodePoolTimeoutsOutputReference {
	var returns K8SNodePoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_K8SNodePool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/k8s_node_pool ionoscloud_k8s_node_pool} Resource.
func NewK8SNodePool(scope constructs.Construct, id *string, config *K8SNodePoolConfig) K8SNodePool {
	_init_.Initialize()

	if err := validateNewK8SNodePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_K8SNodePool{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/resources/k8s_node_pool ionoscloud_k8s_node_pool} Resource.
func NewK8SNodePool_Override(k K8SNodePool, scope constructs.Construct, id *string, config *K8SNodePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_K8SNodePool)SetAllowReplace(val interface{}) {
	if err := j.validateSetAllowReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplace",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetCoresCount(val *float64) {
	if err := j.validateSetCoresCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coresCount",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetCpuFamily(val *string) {
	if err := j.validateSetCpuFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuFamily",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetK8SClusterId(val *string) {
	if err := j.validateSetK8SClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k8SClusterId",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetK8SVersion(val *string) {
	if err := j.validateSetK8SVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k8SVersion",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetPublicIps(val *[]*string) {
	if err := j.validateSetPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIps",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetRamSize(val *float64) {
	if err := j.validateSetRamSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramSize",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetStorageSize(val *float64) {
	if err := j.validateSetStorageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSize",
		val,
	)
}

func (j *jsiiProxy_K8SNodePool)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

// Generates CDKTF code for importing a K8SNodePool resource upon running "cdktf plan <stack-name>".
func K8SNodePool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateK8SNodePool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
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
func K8SNodePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SNodePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func K8SNodePool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SNodePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func K8SNodePool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateK8SNodePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func K8SNodePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.k8SNodePool.K8SNodePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_K8SNodePool) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_K8SNodePool) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_K8SNodePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_K8SNodePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_K8SNodePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_K8SNodePool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_K8SNodePool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_K8SNodePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_K8SNodePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_K8SNodePool) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_K8SNodePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_K8SNodePool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_K8SNodePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_K8SNodePool) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_K8SNodePool) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_K8SNodePool) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_K8SNodePool) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_K8SNodePool) PutAutoScaling(value *K8SNodePoolAutoScaling) {
	if err := k.validatePutAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAutoScaling",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SNodePool) PutLans(value interface{}) {
	if err := k.validatePutLansParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLans",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SNodePool) PutMaintenanceWindow(value *K8SNodePoolMaintenanceWindow) {
	if err := k.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SNodePool) PutTimeouts(value *K8SNodePoolTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_K8SNodePool) ResetAllowReplace() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowReplace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetAnnotations() {
	_jsii_.InvokeVoid(
		k,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetAutoScaling() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoScaling",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetLans() {
	_jsii_.InvokeVoid(
		k,
		"resetLans",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetPublicIps() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicIps",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_K8SNodePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_K8SNodePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

