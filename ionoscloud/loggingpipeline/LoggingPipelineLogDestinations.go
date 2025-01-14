// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingpipeline


type LoggingPipelineLogDestinations struct {
	// Defines the number of days a log record should be kept in loki.
	//
	// Works with loki destination type only. Possible values are: 7, 14, 30.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/logging_pipeline#retention_in_days LoggingPipeline#retention_in_days}
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.6.9/docs/resources/logging_pipeline#type LoggingPipeline#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

