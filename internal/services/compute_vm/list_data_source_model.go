// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ComputeVMsItemsDataSourceModel] `json:"items,computed"`
}

type ComputeVMsDataSourceModel struct {
	MaxItems types.Int64                                                  `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ComputeVMsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ComputeVMsDataSourceModel) toListParams(_ context.Context) (params compute.VMListParams, diags diag.Diagnostics) {
	params = compute.VMListParams{}

	return
}

type ComputeVMsItemsDataSourceModel struct {
	ID              types.String                                                    `tfsdk:"id" json:"id,computed"`
	BootVolumeID    types.String                                                    `tfsdk:"boot_volume_id" json:"boot_volume_id,computed"`
	CPUConfig       customfield.NestedObject[ComputeVMsCPUConfigDataSourceModel]    `tfsdk:"cpu_config" json:"cpu_config,computed"`
	CreatedAt       timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DataVolumeIDs   customfield.List[types.String]                                  `tfsdk:"data_volume_ids" json:"data_volume_ids,computed"`
	MemoryConfig    customfield.NestedObject[ComputeVMsMemoryConfigDataSourceModel] `tfsdk:"memory_config" json:"memory_config,computed"`
	Name            types.String                                                    `tfsdk:"name" json:"name,computed"`
	PrivateIP       types.String                                                    `tfsdk:"private_ip" json:"private_ip,computed"`
	PublicIP        types.String                                                    `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPEnabled types.Bool                                                      `tfsdk:"public_ip_enabled" json:"public_ip_enabled,computed"`
	Region          types.String                                                    `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                    `tfsdk:"status" json:"status,computed"`
	SubnetID        types.String                                                    `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Tags            customfield.List[types.String]                                  `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt       timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID           types.String                                                    `tfsdk:"vpc_id" json:"vpc_id,computed"`
	VPCName         types.String                                                    `tfsdk:"vpc_name" json:"vpc_name,computed"`
}

type ComputeVMsCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type ComputeVMsMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
