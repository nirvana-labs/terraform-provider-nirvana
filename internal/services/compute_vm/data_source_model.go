// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMDataSourceModel struct {
	ID              types.String                                                   `tfsdk:"id" path:"vm_id,computed"`
	VMID            types.String                                                   `tfsdk:"vm_id" path:"vm_id,optional"`
	BootVolumeID    types.String                                                   `tfsdk:"boot_volume_id" json:"boot_volume_id,computed"`
	CreatedAt       timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name            types.String                                                   `tfsdk:"name" json:"name,computed"`
	PrivateIP       types.String                                                   `tfsdk:"private_ip" json:"private_ip,computed"`
	ProjectID       types.String                                                   `tfsdk:"project_id" json:"project_id,computed"`
	PublicIP        types.String                                                   `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPEnabled types.Bool                                                     `tfsdk:"public_ip_enabled" json:"public_ip_enabled,computed"`
	Region          types.String                                                   `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                   `tfsdk:"status" json:"status,computed"`
	SubnetID        types.String                                                   `tfsdk:"subnet_id" json:"subnet_id,computed"`
	UpdatedAt       timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID           types.String                                                   `tfsdk:"vpc_id" json:"vpc_id,computed"`
	VPCName         types.String                                                   `tfsdk:"vpc_name" json:"vpc_name,computed"`
	DataVolumeIDs   customfield.List[types.String]                                 `tfsdk:"data_volume_ids" json:"data_volume_ids,computed"`
	Tags            customfield.List[types.String]                                 `tfsdk:"tags" json:"tags,computed"`
	CPUConfig       customfield.NestedObject[ComputeVMCPUConfigDataSourceModel]    `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemoryConfig    customfield.NestedObject[ComputeVMMemoryConfigDataSourceModel] `tfsdk:"memory_config" json:"memory_config,computed"`
	FindOneBy       *ComputeVMFindOneByDataSourceModel                             `tfsdk:"find_one_by"`
}

func (m *ComputeVMDataSourceModel) toListParams(_ context.Context) (params compute.VMListParams, diags diag.Diagnostics) {
	params = compute.VMListParams{}

	if !m.FindOneBy.ProjectID.IsNull() {
		params.ProjectID = param.NewOpt(m.FindOneBy.ProjectID.ValueString())
	}

	return
}

type ComputeVMCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type ComputeVMMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}

type ComputeVMFindOneByDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" query:"project_id,optional"`
}
