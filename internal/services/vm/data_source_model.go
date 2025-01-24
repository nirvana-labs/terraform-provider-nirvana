// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type VMDataSourceModel struct {
	VMID        types.String                                               `tfsdk:"vm_id" path:"vm_id,required"`
	CreatedAt   types.String                                               `tfsdk:"created_at" json:"created_at,computed"`
	ID          types.String                                               `tfsdk:"id" json:"id,computed"`
	Name        types.String                                               `tfsdk:"name" json:"name,computed"`
	PublicIP    types.String                                               `tfsdk:"public_ip" json:"public_ip,computed"`
	Region      types.String                                               `tfsdk:"region" json:"region,computed"`
	Status      types.String                                               `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                               `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID       types.String                                               `tfsdk:"vpc_id" json:"vpc_id,computed"`
	BootVolume  customfield.NestedObject[VMBootVolumeDataSourceModel]      `tfsdk:"boot_volume" json:"boot_volume,computed"`
	CPUConfig   customfield.NestedObject[VMCPUConfigDataSourceModel]       `tfsdk:"cpu_config" json:"cpu_config,computed"`
	DataVolumes customfield.NestedObjectList[VMDataVolumesDataSourceModel] `tfsdk:"data_volumes" json:"data_volumes,computed"`
	MemConfig   customfield.NestedObject[VMMemConfigDataSourceModel]       `tfsdk:"mem_config" json:"mem_config,computed"`
}

type VMBootVolumeDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind      types.String `tfsdk:"kind" json:"kind,computed"`
	Size      types.Int64  `tfsdk:"size" json:"size,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

type VMCPUConfigDataSourceModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,computed"`
}

type VMDataVolumesDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind      types.String `tfsdk:"kind" json:"kind,computed"`
	Size      types.Int64  `tfsdk:"size" json:"size,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

type VMMemConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
