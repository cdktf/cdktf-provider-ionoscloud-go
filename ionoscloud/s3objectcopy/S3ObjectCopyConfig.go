// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3objectcopy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ObjectCopyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#bucket S3ObjectCopy#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The key of the object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#key S3ObjectCopy#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The key of the source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#source S3ObjectCopy#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Can be used to specify caching behavior along the request/reply chain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#cache_control S3ObjectCopy#cache_control}
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Specifies presentational information for the object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#content_disposition S3ObjectCopy#content_disposition}
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Specifies what content encodings have been applied to the object copy and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#content_encoding S3ObjectCopy#content_encoding}
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// The natural language or languages of the intended audience for the object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#content_language S3ObjectCopy#content_language}
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// A standard MIME type describing the format of the contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#content_type S3ObjectCopy#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Copies the object if its entity tag (ETag) matches the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#copy_if_match S3ObjectCopy#copy_if_match}
	CopyIfMatch *string `field:"optional" json:"copyIfMatch" yaml:"copyIfMatch"`
	// Copies the object if it has been modified since the specified time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#copy_if_modified_since S3ObjectCopy#copy_if_modified_since}
	CopyIfModifiedSince *string `field:"optional" json:"copyIfModifiedSince" yaml:"copyIfModifiedSince"`
	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#copy_if_none_match S3ObjectCopy#copy_if_none_match}
	CopyIfNoneMatch *string `field:"optional" json:"copyIfNoneMatch" yaml:"copyIfNoneMatch"`
	// Copies the object if it hasn't been modified since the specified time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#copy_if_unmodified_since S3ObjectCopy#copy_if_unmodified_since}
	CopyIfUnmodifiedSince *string `field:"optional" json:"copyIfUnmodifiedSince" yaml:"copyIfUnmodifiedSince"`
	// The date and time at which the object copy is no longer cacheable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#expires S3ObjectCopy#expires}
	Expires *string `field:"optional" json:"expires" yaml:"expires"`
	// Specifies whether to delete the object copy even if it has a governance-type Object Copy Lock in place.
	//
	// You must explicitly pass a value of true for this parameter to delete the object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#force_destroy S3ObjectCopy#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// A map of metadata to store with the object copy in IONOS Object Storage Object Copy Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#metadata S3ObjectCopy#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#metadata_directive S3ObjectCopy#metadata_directive}
	MetadataDirective *string `field:"optional" json:"metadataDirective" yaml:"metadataDirective"`
	// Specifies whether a legal hold will be applied to this object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#object_lock_legal_hold S3ObjectCopy#object_lock_legal_hold}
	ObjectLockLegalHold *string `field:"optional" json:"objectLockLegalHold" yaml:"objectLockLegalHold"`
	// Confirms that the requester knows that they will be charged for the request.
	//
	// Bucket owners need not specify this parameter in their requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#object_lock_mode S3ObjectCopy#object_lock_mode}
	ObjectLockMode *string `field:"optional" json:"objectLockMode" yaml:"objectLockMode"`
	// The date and time when you want this object copy's Object Copy Lock to expire.
	//
	// Must be formatted as a timestamp parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#object_lock_retain_until_date S3ObjectCopy#object_lock_retain_until_date}
	ObjectLockRetainUntilDate *string `field:"optional" json:"objectLockRetainUntilDate" yaml:"objectLockRetainUntilDate"`
	// The server-side encryption algorithm used when storing this object copy in IONOS Object Storage Object Copy Storage (AES256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#server_side_encryption S3ObjectCopy#server_side_encryption}
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Specifies the algorithm to use to when encrypting the object copy (e.g., AES256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#server_side_encryption_customer_algorithm S3ObjectCopy#server_side_encryption_customer_algorithm}
	ServerSideEncryptionCustomerAlgorithm *string `field:"optional" json:"serverSideEncryptionCustomerAlgorithm" yaml:"serverSideEncryptionCustomerAlgorithm"`
	// Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#server_side_encryption_customer_key S3ObjectCopy#server_side_encryption_customer_key}
	ServerSideEncryptionCustomerKey *string `field:"optional" json:"serverSideEncryptionCustomerKey" yaml:"serverSideEncryptionCustomerKey"`
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	//
	// IONOS Object Storage Object Copy Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#server_side_encryption_customer_key_md5 S3ObjectCopy#server_side_encryption_customer_key_md5}
	ServerSideEncryptionCustomerKeyMd5 *string `field:"optional" json:"serverSideEncryptionCustomerKeyMd5" yaml:"serverSideEncryptionCustomerKeyMd5"`
	// Specifies the algorithm to use to when decrypting the source object (e.g., AES256).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#source_customer_algorithm S3ObjectCopy#source_customer_algorithm}
	SourceCustomerAlgorithm *string `field:"optional" json:"sourceCustomerAlgorithm" yaml:"sourceCustomerAlgorithm"`
	// Specifies the 256-bit, base64-encoded encryption key to use to decrypt the source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#source_customer_key S3ObjectCopy#source_customer_key}
	SourceCustomerKey *string `field:"optional" json:"sourceCustomerKey" yaml:"sourceCustomerKey"`
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	//
	// IONOS Object Storage Object Copy Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#source_customer_key_md5 S3ObjectCopy#source_customer_key_md5}
	SourceCustomerKeyMd5 *string `field:"optional" json:"sourceCustomerKeyMd5" yaml:"sourceCustomerKeyMd5"`
	// The storage class of the object copy. Valid value is 'STANDARD'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#storage_class S3ObjectCopy#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Specifies whether the object copy tag-set is copied from the source object or replaced with tag-set provided in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#tagging_directive S3ObjectCopy#tagging_directive}
	TaggingDirective *string `field:"optional" json:"taggingDirective" yaml:"taggingDirective"`
	// The tag-set for the object copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#tags S3ObjectCopy#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// If the bucket is configured as a website, redirects requests for this object copy to another object copy in the same bucket or to an external URL.
	//
	// IONOS Object Storage Object Copy Storage stores the value of this header in the object copy metadata
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.4/docs/resources/s3_object_copy#website_redirect S3ObjectCopy#website_redirect}
	WebsiteRedirect *string `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
}

