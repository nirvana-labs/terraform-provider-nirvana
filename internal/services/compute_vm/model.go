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
	SourceAddress   types.String                                            `tfsdk:"source_address" json:"source_address,required"`
	Ports           *[]types.String                                         `tfsdk:"ports" json:"ports,required"`
	SSHKey          *ComputeVMSSHKeyModel                                   `tfsdk:"ssh_key" json:"ssh_key,required"`
	SubnetID        types.String                                            `tfsdk:"subnet_id" json:"subnet_id,optional"`
	BootVolume      *ComputeVMBootVolumeModel                               `tfsdk:"boot_volume" json:"boot_volume,required"`
	CPU             *ComputeVMCPUModel                                      `tfsdk:"cpu" json:"cpu,required"`
	Ram             *ComputeVMRamModel                                      `tfsdk:"ram" json:"ram,required"`
	DataVolumes     customfield.NestedObjectList[ComputeVMDataVolumesModel] `tfsdk:"data_volumes" json:"data_volumes,computed_optional"`
	CreatedAt       types.String                                            `tfsdk:"created_at" json:"created_at,computed"`
	Kind            types.String                                            `tfsdk:"kind" json:"kind,computed"`
	PublicIP        types.String                                            `tfsdk:"public_ip" json:"public_ip,computed"`
	ResourceID      types.String                                            `tfsdk:"resource_id" json:"resource_id,computed"`
	Status          types.String                                            `tfsdk:"status" json:"status,computed"`
	Type            types.String                                            `tfsdk:"type" json:"type,computed"`
	UpdatedAt       types.String                                            `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID           types.String                                            `tfsdk:"vpc_id" json:"vpc_id,computed"`
	CPUConfig       customfield.NestedObject[ComputeVMCPUConfigModel]       `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemConfig       customfield.NestedObject[ComputeVMMemConfigModel]       `tfsdk:"mem_config" json:"mem_config,computed"`
}

func (m ComputeVMModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeVMModel) MarshalJSONForUpdate(state ComputeVMModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ComputeVMSSHKeyModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type ComputeVMBootVolumeModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type ComputeVMCPUModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,required"`
}

type ComputeVMRamModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type ComputeVMDataVolumesModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,required"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}

type ComputeVMCPUConfigModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,computed"`
}

type ComputeVMMemConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
