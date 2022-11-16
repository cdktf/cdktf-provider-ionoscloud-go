package cubeserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v4/jsii"

	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v4/cubeserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CubeServerVolumeOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupUnitId() *string
	SetBackupUnitId(val *string)
	BackupUnitIdInput() *string
	BootServer() *string
	Bus() *string
	SetBus(val *string)
	BusInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CpuHotPlug() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceNumber() *float64
	DiscVirtioHotPlug() cdktf.IResolvable
	DiscVirtioHotUnplug() cdktf.IResolvable
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	ImagePassword() *string
	SetImagePassword(val *string)
	ImagePasswordInput() *string
	InternalValue() *CubeServerVolume
	SetInternalValue(val *CubeServerVolume)
	LicenceType() *string
	SetLicenceType(val *string)
	LicenceTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NicHotPlug() cdktf.IResolvable
	NicHotUnplug() cdktf.IResolvable
	PciSlot() *float64
	RamHotPlug() cdktf.IResolvable
	SshKeyPath() *[]*string
	SetSshKeyPath(val *[]*string)
	SshKeyPathInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAvailabilityZone()
	ResetBackupUnitId()
	ResetBus()
	ResetImagePassword()
	ResetLicenceType()
	ResetName()
	ResetSshKeyPath()
	ResetUserData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CubeServerVolumeOutputReference
type jsiiProxy_CubeServerVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) BackupUnitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) BackupUnitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupUnitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) BootServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) Bus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) BusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) CpuHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) DeviceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) DiscVirtioHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) DiscVirtioHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discVirtioHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) ImagePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) ImagePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) InternalValue() *CubeServerVolume {
	var returns *CubeServerVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) LicenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) LicenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) NicHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) NicHotUnplug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nicHotUnplug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) PciSlot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pciSlot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) RamHotPlug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ramHotPlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) SshKeyPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) SshKeyPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CubeServerVolumeOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewCubeServerVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CubeServerVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewCubeServerVolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CubeServerVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCubeServerVolumeOutputReference_Override(c CubeServerVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.cubeServer.CubeServerVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetBackupUnitId(val *string) {
	if err := j.validateSetBackupUnitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupUnitId",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetBus(val *string) {
	if err := j.validateSetBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bus",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetImagePassword(val *string) {
	if err := j.validateSetImagePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePassword",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetInternalValue(val *CubeServerVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetLicenceType(val *string) {
	if err := j.validateSetLicenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenceType",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetSshKeyPath(val *[]*string) {
	if err := j.validateSetSshKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeyPath",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CubeServerVolumeOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CubeServerVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetBackupUnitId() {
	_jsii_.InvokeVoid(
		c,
		"resetBackupUnitId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetBus() {
	_jsii_.InvokeVoid(
		c,
		"resetBus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetImagePassword() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetLicenceType() {
	_jsii_.InvokeVoid(
		c,
		"resetLicenceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetSshKeyPath() {
	_jsii_.InvokeVoid(
		c,
		"resetSshKeyPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		c,
		"resetUserData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CubeServerVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

