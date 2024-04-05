// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingpipeline


type LoggingPipelineLog struct {
	// Protocol to use as intake. Possible values are: http, tcp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.14/docs/resources/logging_pipeline#protocol LoggingPipeline#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The source parser to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.14/docs/resources/logging_pipeline#source LoggingPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The tag is used to distinguish different pipelines. Must be unique amongst the pipeline's array items.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.14/docs/resources/logging_pipeline#tag LoggingPipeline#tag}
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.14/docs/resources/logging_pipeline#destinations LoggingPipeline#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

