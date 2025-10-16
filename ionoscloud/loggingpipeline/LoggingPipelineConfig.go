// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingpipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingPipelineConfig struct {
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
	// log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/logging_pipeline#log LoggingPipeline#log}
	Log interface{} `field:"required" json:"log" yaml:"log"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/logging_pipeline#name LoggingPipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/logging_pipeline#id LoggingPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of your logging pipeline. Default: de/txl. Supported locations: de/fra, de/fra/2, de/txl, es/vit, gb/lhr, fr/par.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/logging_pipeline#location LoggingPipeline#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.7.16/docs/resources/logging_pipeline#timeouts LoggingPipeline#timeouts}
	Timeouts *LoggingPipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

