// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMDataSourceModel struct {
	VMID          types.String                                                   `tfsdk:"vm_id" path:"vm_id,required"`
	BootVolumeID  types.String                                                   `tfsdk:"boot_volume_id" json:"boot_volume_id,computed"`
	CreatedAt     types.String                                                   `tfsdk:"created_at" json:"created_at,computed"`
	ID            types.String                                                   `tfsdk:"id" json:"id,computed"`
	Name          types.String                                                   `tfsdk:"name" json:"name,computed"`
	PrivateIP     types.String                                                   `tfsdk:"private_ip" json:"private_ip,computed"`
	PublicIP      types.String                                                   `tfsdk:"public_ip" json:"public_ip,computed"`
	Region        types.String                                                   `tfsdk:"region" json:"region,computed"`
	Status        types.String                                                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt     types.String                                                   `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID         types.String                                                   `tfsdk:"vpc_id" json:"vpc_id,computed"`
	DataVolumeIDs customfield.List[types.String]                                 `tfsdk:"data_volume_ids" json:"data_volume_ids,computed"`
	CPUConfig     customfield.NestedObject[ComputeVMCPUConfigDataSourceModel]    `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemoryConfig  customfield.NestedObject[ComputeVMMemoryConfigDataSourceModel] `tfsdk:"memory_config" json:"memory_config,computed"`
}

type ComputeVMCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type ComputeVMMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
