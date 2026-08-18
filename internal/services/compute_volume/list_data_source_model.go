// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVolumesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ComputeVolumesItemsDataSourceModel] `json:"items,computed"`
}

type ComputeVolumesDataSourceModel struct {
	ProjectID types.String                                                     `tfsdk:"project_id" query:"project_id,required"`
	Attached  types.Bool                                                       `tfsdk:"attached" query:"attached,optional"`
	Kind      types.String                                                     `tfsdk:"kind" query:"kind,optional"`
	Name      types.String                                                     `tfsdk:"name" query:"name,optional"`
	Region    types.String                                                     `tfsdk:"region" query:"region,optional"`
	Status    types.String                                                     `tfsdk:"status" query:"status,optional"`
	Type      types.String                                                     `tfsdk:"type" query:"type,optional"`
	VMID      types.String                                                     `tfsdk:"vm_id" query:"vm_id,optional"`
	Tags      *[]types.String                                                  `tfsdk:"tags" query:"tags,optional"`
	Sort      types.String                                                     `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[ComputeVolumesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ComputeVolumesDataSourceModel) toListParams(_ context.Context) (params compute.VolumeListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = compute.VolumeListParams{
		ProjectID: m.ProjectID.ValueString(),
		Tags:      mTags,
	}

	if !m.Attached.IsNull() {
		params.Attached = param.NewOpt(m.Attached.ValueBool())
	}
	if !m.Kind.IsNull() {
		params.Kind = compute.VolumeListParamsKind(m.Kind.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = compute.VolumeListParamsStatus(m.Status.ValueString())
	}
	if !m.Type.IsNull() {
		params.Type = compute.VolumeListParamsType(m.Type.ValueString())
	}
	if !m.VMID.IsNull() {
		params.VMID = param.NewOpt(m.VMID.ValueString())
	}

	return
}

type ComputeVolumesItemsDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind      types.String                   `tfsdk:"kind" json:"kind,computed"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	ProjectID types.String                   `tfsdk:"project_id" json:"project_id,computed"`
	Region    types.String                   `tfsdk:"region" json:"region,computed"`
	Size      types.Int64                    `tfsdk:"size" json:"size,computed"`
	Status    types.String                   `tfsdk:"status" json:"status,computed"`
	Tags      customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	Type      types.String                   `tfsdk:"type" json:"type,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VMID      types.String                   `tfsdk:"vm_id" json:"vm_id,computed"`
	VMName    types.String                   `tfsdk:"vm_name" json:"vm_name,computed"`
}
