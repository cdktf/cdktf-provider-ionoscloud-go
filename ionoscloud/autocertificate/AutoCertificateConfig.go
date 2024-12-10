// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autocertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoCertificateConfig struct {
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
	// The common name (DNS) of the certificate to issue.
	//
	// The common name needs to be part of a zone in IONOS Cloud DNS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#common_name AutoCertificate#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The key algorithm used to generate the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#key_algorithm AutoCertificate#key_algorithm}
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// The location of the auto-certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#location AutoCertificate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The certificate provider used to issue the certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#provider_id AutoCertificate#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#id AutoCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A certificate name used for management purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#name AutoCertificate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional additional names to be added to the issued certificate.
	//
	// The additional names needs to be part of a zone in IONOS Cloud DNS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#subject_alternative_names AutoCertificate#subject_alternative_names}
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.7/docs/resources/auto_certificate#timeouts AutoCertificate#timeouts}
	Timeouts *AutoCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

