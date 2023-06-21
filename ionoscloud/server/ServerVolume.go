package server


type ServerVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#disk_type Server#disk_type}.
	DiskType *string `field:"required" json:"diskType" yaml:"diskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#availability_zone Server#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The uuid of the Backup Unit that user has access to.
	//
	// The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#backup_unit_id Server#backup_unit_id}
	BackupUnitId *string `field:"optional" json:"backupUnitId" yaml:"backupUnitId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#bus Server#bus}.
	Bus *string `field:"optional" json:"bus" yaml:"bus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#image_password Server#image_password}.
	ImagePassword *string `field:"optional" json:"imagePassword" yaml:"imagePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#licence_type Server#licence_type}.
	LicenceType *string `field:"optional" json:"licenceType" yaml:"licenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#name Server#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The size of the volume in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#size Server#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key.
	//
	// This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#ssh_key_path Server#ssh_key_path}
	SshKeyPath *[]*string `field:"optional" json:"sshKeyPath" yaml:"sshKeyPath"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key.
	//
	// This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#ssh_keys Server#ssh_keys}
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// The cloud-init configuration for the volume as base64 encoded string.
	//
	// The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/ionos-cloud/ionoscloud/6.4.1/docs/resources/server#user_data Server#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

