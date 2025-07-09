// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3object

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ObjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#bucket S3Object#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The key of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#key S3Object#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Can be used to specify caching behavior along the request/reply chain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#cache_control S3Object#cache_control}
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// The utf-8 content of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#content S3Object#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Specifies presentational information for the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#content_disposition S3Object#content_disposition}
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#content_encoding S3Object#content_encoding}
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// The natural language or languages of the intended audience for the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#content_language S3Object#content_language}
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// A standard MIME type describing the format of the contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#content_type S3Object#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The date and time at which the object is no longer cacheable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#expires S3Object#expires}
	Expires *string `field:"optional" json:"expires" yaml:"expires"`
	// Specifies whether to delete the object even if it has a governance-type Object Lock in place.
	//
	// You must explicitly pass a value of true for this parameter to delete the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#force_destroy S3Object#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// A map of metadata to store with the object in IONOS Object Storage Object Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#metadata S3Object#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
	//
	// Required to permanently delete a versioned object if versioning is configured with MFA Delete enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#mfa S3Object#mfa}
	Mfa *string `field:"optional" json:"mfa" yaml:"mfa"`
	// Specifies whether a legal hold will be applied to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#object_lock_legal_hold S3Object#object_lock_legal_hold}
	ObjectLockLegalHold *string `field:"optional" json:"objectLockLegalHold" yaml:"objectLockLegalHold"`
	// Confirms that the requester knows that they will be charged for the request.
	//
	// Bucket owners need not specify this parameter in their requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#object_lock_mode S3Object#object_lock_mode}
	ObjectLockMode *string `field:"optional" json:"objectLockMode" yaml:"objectLockMode"`
	// The date and time when you want this object's Object Lock to expire.
	//
	// Must be formatted as a timestamp parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#object_lock_retain_until_date S3Object#object_lock_retain_until_date}
	ObjectLockRetainUntilDate *string `field:"optional" json:"objectLockRetainUntilDate" yaml:"objectLockRetainUntilDate"`
	// Confirms that the requester knows that they will be charged for the request.
	//
	// Bucket owners need not specify this parameter in their requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#request_payer S3Object#request_payer}
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// The server-side encryption algorithm used when storing this object in IONOS Object Storage Object Storage (AES256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#server_side_encryption S3Object#server_side_encryption}
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Specifies the IONOS Object Storage Object Storage Encryption Context to use for object encryption.
	//
	// The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#server_side_encryption_context S3Object#server_side_encryption_context}
	ServerSideEncryptionContext *string `field:"optional" json:"serverSideEncryptionContext" yaml:"serverSideEncryptionContext"`
	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#server_side_encryption_customer_algorithm S3Object#server_side_encryption_customer_algorithm}
	ServerSideEncryptionCustomerAlgorithm *string `field:"optional" json:"serverSideEncryptionCustomerAlgorithm" yaml:"serverSideEncryptionCustomerAlgorithm"`
	// Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#server_side_encryption_customer_key S3Object#server_side_encryption_customer_key}
	ServerSideEncryptionCustomerKey *string `field:"optional" json:"serverSideEncryptionCustomerKey" yaml:"serverSideEncryptionCustomerKey"`
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	//
	// IONOS Object Storage Object Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#server_side_encryption_customer_key_md5 S3Object#server_side_encryption_customer_key_md5}
	ServerSideEncryptionCustomerKeyMd5 *string `field:"optional" json:"serverSideEncryptionCustomerKeyMd5" yaml:"serverSideEncryptionCustomerKeyMd5"`
	// The path to the file to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#source S3Object#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The storage class of the object. Valid value is 'STANDARD'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#storage_class S3Object#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// The tag-set for the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#tags S3Object#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.
	//
	// IONOS Object Storage Object Storage stores the value of this header in the object metadata
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.11/docs/resources/s3_object#website_redirect S3Object#website_redirect}
	WebsiteRedirect *string `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
}

