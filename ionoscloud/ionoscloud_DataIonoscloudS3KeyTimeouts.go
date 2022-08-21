// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type DataIonoscloudS3KeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/s3_key#create DataIonoscloudS3Key#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/s3_key#default DataIonoscloudS3Key#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/s3_key#delete DataIonoscloudS3Key#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/d/s3_key#update DataIonoscloudS3Key#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
