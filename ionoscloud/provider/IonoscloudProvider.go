// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ionoscloud-go/ionoscloud/v11/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs ionoscloud}.
type IonoscloudProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContractNumber() *string
	SetContractNumber(val *string)
	ContractNumberInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	S3AccessKey() *string
	SetS3AccessKey(val *string)
	S3AccessKeyInput() *string
	S3Region() *string
	SetS3Region(val *string)
	S3RegionInput() *string
	S3SecretKey() *string
	SetS3SecretKey(val *string)
	S3SecretKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetContractNumber()
	ResetEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetRetries()
	ResetS3AccessKey()
	ResetS3Region()
	ResetS3SecretKey()
	ResetToken()
	ResetUsername()
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

// The jsii proxy struct for IonoscloudProvider
type jsiiProxy_IonoscloudProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_IonoscloudProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) ContractNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contractNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) ContractNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contractNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) S3SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IonoscloudProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs ionoscloud} Resource.
func NewIonoscloudProvider(scope constructs.Construct, id *string, config *IonoscloudProviderConfig) IonoscloudProvider {
	_init_.Initialize()

	if err := validateNewIonoscloudProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IonoscloudProvider{}

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.5.2/docs ionoscloud} Resource.
func NewIonoscloudProvider_Override(i IonoscloudProvider, scope constructs.Construct, id *string, config *IonoscloudProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetContractNumber(val *string) {
	_jsii_.Set(
		j,
		"contractNumber",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetRetries(val *float64) {
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetS3AccessKey(val *string) {
	_jsii_.Set(
		j,
		"s3AccessKey",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetS3Region(val *string) {
	_jsii_.Set(
		j,
		"s3Region",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetS3SecretKey(val *string) {
	_jsii_.Set(
		j,
		"s3SecretKey",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_IonoscloudProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a IonoscloudProvider resource upon running "cdktf plan <stack-name>".
func IonoscloudProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIonoscloudProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
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
func IonoscloudProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIonoscloudProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IonoscloudProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIonoscloudProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IonoscloudProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIonoscloudProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IonoscloudProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ionoscloud.provider.IonoscloudProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IonoscloudProvider) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IonoscloudProvider) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		i,
		"resetAlias",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetContractNumber() {
	_jsii_.InvokeVoid(
		i,
		"resetContractNumber",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetEndpoint() {
	_jsii_.InvokeVoid(
		i,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetRetries() {
	_jsii_.InvokeVoid(
		i,
		"resetRetries",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetS3AccessKey() {
	_jsii_.InvokeVoid(
		i,
		"resetS3AccessKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetS3Region() {
	_jsii_.InvokeVoid(
		i,
		"resetS3Region",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetS3SecretKey() {
	_jsii_.InvokeVoid(
		i,
		"resetS3SecretKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetToken() {
	_jsii_.InvokeVoid(
		i,
		"resetToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		i,
		"resetUsername",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IonoscloudProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IonoscloudProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IonoscloudProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IonoscloudProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IonoscloudProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IonoscloudProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

