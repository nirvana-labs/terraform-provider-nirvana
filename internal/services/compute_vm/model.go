// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMModel struct {
	ID              types.String                                            `tfsdk:"id" json:"id,computed"`
	Name            types.String                                            `tfsdk:"name" json:"name,required"`
	OSImageName     types.String                                            `tfsdk:"os_image_name" json:"os_image_name,required"`
	PublicIPEnabled types.Bool                                              `tfsdk:"public_ip_enabled" json:"public_ip_enabled,required"`
	Region          types.String                                            `tfsdk:"region" json:"region,required"`
	BootVolume      *ComputeVMBootVolumeModel                               `tfsdk:"boot_volume" json:"boot_volume,required"`
	SSHKey          *ComputeVMSSHKeyModel                                   `tfsdk:"ssh_key" json:"ssh_key,required"`
	SubnetID        types.String                                            `tfsdk:"subnet_id" json:"subnet_id,optional"`
	DataVolumes     customfield.NestedObjectList[ComputeVMDataVolumesModel] `tfsdk:"data_volumes" json:"data_volumes,computed_optional"`
	CPUConfig       *ComputeVMCPUConfigModel                                `tfsdk:"cpu_config" json:"cpu_config,required"`
	MemoryConfig    *ComputeVMMemoryConfigModel                             `tfsdk:"memory_config" json:"memory_config,required"`
	BootVolumeID    types.String                                            `tfsdk:"boot_volume_id" json:"boot_volume_id,computed"`
	CreatedAt       types.String                                            `tfsdk:"created_at" json:"created_at,computed"`
	Kind            types.String                                            `tfsdk:"kind" json:"kind,computed"`
	PrivateIP       types.String                                            `tfsdk:"private_ip" json:"private_ip,computed"`
	PublicIP        types.String                                            `tfsdk:"public_ip" json:"public_ip,computed"`
	ResourceID      types.String                                            `tfsdk:"resource_id" json:"resource_id,computed"`
	Status          types.String                                            `tfsdk:"status" json:"status,computed"`
	Type            types.String                                            `tfsdk:"type" json:"type,computed"`
	UpdatedAt       types.String                                            `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID           types.String                                            `tfsdk:"vpc_id" json:"vpc_id,computed"`
	DataVolumeIDs   customfield.List[types.String]                          `tfsdk:"data_volume_ids" json:"data_volume_ids,computed"`
}

func (m ComputeVMModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeVMModel) MarshalJSONForUpdate(state ComputeVMModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ComputeVMBootVolumeModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type ComputeVMSSHKeyModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type ComputeVMDataVolumesModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type ComputeVMCPUConfigModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,required"`
}

type ComputeVMMemoryConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}
