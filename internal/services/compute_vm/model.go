// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMModel struct {
	ID              types.String                                            `tfsdk:"id" json:"id,computed"`
	OSImageName     types.String                                            `tfsdk:"os_image_name" json:"os_image_name,required,no_refresh"`
	Region          types.String                                            `tfsdk:"region" json:"region,required"`
	SubnetID        types.String                                            `tfsdk:"subnet_id" json:"subnet_id,required"`
	BootVolume      *ComputeVMBootVolumeModel                               `tfsdk:"boot_volume" json:"boot_volume,required,no_refresh"`
	SSHKey          *ComputeVMSSHKeyModel                                   `tfsdk:"ssh_key" json:"ssh_key,required,no_refresh"`
	DataVolumes     customfield.NestedObjectList[ComputeVMDataVolumesModel] `tfsdk:"data_volumes" json:"data_volumes,computed_optional,no_refresh"`
	Name            types.String                                            `tfsdk:"name" json:"name,optional"`
	PublicIPEnabled types.Bool                                              `tfsdk:"public_ip_enabled" json:"public_ip_enabled,optional"`
	Tags            *[]types.String                                         `tfsdk:"tags" json:"tags,optional"`
	CPUConfig       *ComputeVMCPUConfigModel                                `tfsdk:"cpu_config" json:"cpu_config,optional"`
	MemoryConfig    *ComputeVMMemoryConfigModel                             `tfsdk:"memory_config" json:"memory_config,optional"`
	BootVolumeID    types.String                                            `tfsdk:"boot_volume_id" json:"boot_volume_id,computed"`
	CreatedAt       timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	PrivateIP       types.String                                            `tfsdk:"private_ip" json:"private_ip,computed"`
	PublicIP        types.String                                            `tfsdk:"public_ip" json:"public_ip,computed"`
	Status          types.String                                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID           types.String                                            `tfsdk:"vpc_id" json:"vpc_id,computed"`
	VPCName         types.String                                            `tfsdk:"vpc_name" json:"vpc_name,computed"`
	DataVolumeIDs   customfield.List[types.String]                          `tfsdk:"data_volume_ids" json:"data_volume_ids,computed"`
}

func (m ComputeVMModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeVMModel) MarshalJSONForUpdate(state ComputeVMModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ComputeVMBootVolumeModel struct {
	Size types.Int64     `tfsdk:"size" json:"size,required"`
	Type types.String    `tfsdk:"type" json:"type,computed_optional"`
	Tags *[]types.String `tfsdk:"tags" json:"tags,optional"`
}

type ComputeVMSSHKeyModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type ComputeVMDataVolumesModel struct {
	Name types.String    `tfsdk:"name" json:"name,required"`
	Size types.Int64     `tfsdk:"size" json:"size,required"`
	Type types.String    `tfsdk:"type" json:"type,computed_optional"`
	Tags *[]types.String `tfsdk:"tags" json:"tags,optional"`
}

type ComputeVMCPUConfigModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,required"`
}

type ComputeVMMemoryConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}
