// Prebuilt ionoscloud Provider for Terraform CDK (cdktf)
package ionoscloud


type BackupUnitTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/backup_unit#create BackupUnit#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/backup_unit#default BackupUnit#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/backup_unit#delete BackupUnit#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/backup_unit#update BackupUnit#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

