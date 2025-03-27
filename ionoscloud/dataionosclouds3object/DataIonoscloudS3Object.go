// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataionosclouds3object

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/dataionosclouds3object/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/s3_object ionoscloud_s3_object}.
type DataIonoscloudS3Object interface {
	cdktf.TerraformDataSource
	Body() *string
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	CacheControl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentDisposition() *string
	ContentEncoding() *string
	ContentLanguage() *string
	ContentLength() *float64
	ContentType() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	Expires() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() cdktf.StringMap
	// The tree node.
	Node() constructs.Node
	ObjectLockLegalHold() *string
	ObjectLockMode() *string
	ObjectLockRetainUntilDate() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Range() *string
	SetRange(val *string)
	RangeInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequestPayer() *string
	ServerSideEncryption() *string
	ServerSideEncryptionContext() *string
	ServerSideEncryptionCustomerAlgorithm() *string
	ServerSideEncryptionCustomerKey() *string
	ServerSideEncryptionCustomerKeyMd5() *string
	StorageClass() *string
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
	WebsiteRedirect() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRange()
	ResetVersionId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataIonoscloudS3Object
type jsiiProxy_DataIonoscloudS3Object struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataIonoscloudS3Object) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ContentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Expires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ObjectLockLegalHold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ServerSideEncryptionContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ServerSideEncryptionCustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionCustomerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ServerSideEncryptionCustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionCustomerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) ServerSideEncryptionCustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionCustomerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIonoscloudS3Object) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/s3_object ionoscloud_s3_object} Data Source.
func NewDataIonoscloudS3Object(scope constructs.Construct, id *string, config *DataIonoscloudS3ObjectConfig) DataIonoscloudS3Object {
	_init_.Initialize()

	if err := validateNewDataIonoscloudS3ObjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIonoscloudS3Object{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.5/docs/data-sources/s3_object ionoscloud_s3_object} Data Source.
func NewDataIonoscloudS3Object_Override(d DataIonoscloudS3Object, scope constructs.Construct, id *string, config *DataIonoscloudS3ObjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetRange(val *string) {
	if err := j.validateSetRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"range",
		val,
	)
}

func (j *jsiiProxy_DataIonoscloudS3Object)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTF code for importing a DataIonoscloudS3Object resource upon running "cdktf plan <stack-name>".
func DataIonoscloudS3Object_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataIonoscloudS3Object_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
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
func DataIonoscloudS3Object_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataIonoscloudS3Object_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataIonoscloudS3Object_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataIonoscloudS3Object_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataIonoscloudS3Object_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataIonoscloudS3Object_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataIonoscloudS3Object_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.dataIonoscloudS3Object.DataIonoscloudS3Object",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataIonoscloudS3Object) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataIonoscloudS3Object) ResetRange() {
	_jsii_.InvokeVoid(
		d,
		"resetRange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataIonoscloudS3Object) ResetVersionId() {
	_jsii_.InvokeVoid(
		d,
		"resetVersionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataIonoscloudS3Object) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIonoscloudS3Object) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

