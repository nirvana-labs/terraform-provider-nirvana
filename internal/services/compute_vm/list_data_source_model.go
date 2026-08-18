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

type ComputeVMsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ComputeVMsItemsDataSourceModel] `json:"items,computed"`
}

type ComputeVMsDataSourceModel struct {
	ProjectID       types.String                                                 `tfsdk:"project_id" query:"project_id,required"`
	Name            types.String                                                 `tfsdk:"name" query:"name,optional"`
	PublicIPEnabled types.Bool                                                   `tfsdk:"public_ip_enabled" query:"public_ip_enabled,optional"`
	Region          types.String                                                 `tfsdk:"region" query:"region,optional"`
	Status          types.String                                                 `tfsdk:"status" query:"status,optional"`
	SubnetID        types.String                                                 `tfsdk:"subnet_id" query:"subnet_id,optional"`
	VPCID           types.String                                                 `tfsdk:"vpc_id" query:"vpc_id,optional"`
	Tags            *[]types.String                                              `tfsdk:"tags" query:"tags,optional"`
	Sort            types.String                                                 `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems        types.Int64                                                  `tfsdk:"max_items"`
	Items           customfield.NestedObjectList[ComputeVMsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ComputeVMsDataSourceModel) toListParams(_ context.Context) (params compute.VMListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = compute.VMListParams{
		ProjectID: m.ProjectID.ValueString(),
		Tags:      mTags,
	}

	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.PublicIPEnabled.IsNull() {
		params.PublicIPEnabled = param.NewOpt(m.PublicIPEnabled.ValueBool())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = compute.VMListParamsStatus(m.Status.ValueString())
	}
	if !m.SubnetID.IsNull() {
		params.SubnetID = param.NewOpt(m.SubnetID.ValueString())
	}
	if !m.VPCID.IsNull() {
		params.VPCID = param.NewOpt(m.VPCID.ValueString())
	}

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
	ProjectID       types.String                                                    `tfsdk:"project_id" json:"project_id,computed"`
	PublicIP        types.String                                                    `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPEnabled types.Bool                                                      `tfsdk:"public_ip_enabled" json:"public_ip_enabled,computed"`
	Region          types.String                                                    `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                    `tfsdk:"status" json:"status,computed"`
	SubnetID        types.String                                                    `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Tags            customfield.List[types.String]                                  `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt       timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID           types.String                                                    `tfsdk:"vpc_id" json:"vpc_id,computed"`
	VPCName         types.String                                                    `tfsdk:"vpc_name" json:"vpc_name,computed"`
	InstanceType    types.String                                                    `tfsdk:"instance_type" json:"instance_type,computed"`
}

type ComputeVMsCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type ComputeVMsMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
