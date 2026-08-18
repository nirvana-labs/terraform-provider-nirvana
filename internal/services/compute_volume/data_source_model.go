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

type ComputeVolumeDataSourceModel struct {
	ID        types.String                           `tfsdk:"id" path:"volume_id,computed"`
	VolumeID  types.String                           `tfsdk:"volume_id" path:"volume_id,optional"`
	CreatedAt timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind      types.String                           `tfsdk:"kind" json:"kind,computed"`
	Name      types.String                           `tfsdk:"name" json:"name,computed"`
	ProjectID types.String                           `tfsdk:"project_id" json:"project_id,computed"`
	Region    types.String                           `tfsdk:"region" json:"region,computed"`
	Size      types.Int64                            `tfsdk:"size" json:"size,computed"`
	Status    types.String                           `tfsdk:"status" json:"status,computed"`
	Type      types.String                           `tfsdk:"type" json:"type,computed"`
	UpdatedAt timetypes.RFC3339                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VMID      types.String                           `tfsdk:"vm_id" json:"vm_id,computed"`
	VMName    types.String                           `tfsdk:"vm_name" json:"vm_name,computed"`
	Tags      customfield.List[types.String]         `tfsdk:"tags" json:"tags,computed"`
	FindOneBy *ComputeVolumeFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ComputeVolumeDataSourceModel) toListParams(_ context.Context) (params compute.VolumeListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = compute.VolumeListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
		Tags:      mFindOneByTags,
	}

	if !m.FindOneBy.Attached.IsNull() {
		params.Attached = param.NewOpt(m.FindOneBy.Attached.ValueBool())
	}
	if !m.FindOneBy.Kind.IsNull() {
		params.Kind = compute.VolumeListParamsKind(m.FindOneBy.Kind.ValueString())
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Region.IsNull() {
		params.Region = param.NewOpt(m.FindOneBy.Region.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = compute.VolumeListParamsStatus(m.FindOneBy.Status.ValueString())
	}
	if !m.FindOneBy.Type.IsNull() {
		params.Type = compute.VolumeListParamsType(m.FindOneBy.Type.ValueString())
	}
	if !m.FindOneBy.VMID.IsNull() {
		params.VMID = param.NewOpt(m.FindOneBy.VMID.ValueString())
	}

	return
}

type ComputeVolumeFindOneByDataSourceModel struct {
	ProjectID types.String    `tfsdk:"project_id" query:"project_id,required"`
	Attached  types.Bool      `tfsdk:"attached" query:"attached,optional"`
	Kind      types.String    `tfsdk:"kind" query:"kind,optional"`
	Name      types.String    `tfsdk:"name" query:"name,optional"`
	Region    types.String    `tfsdk:"region" query:"region,optional"`
	Sort      types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status    types.String    `tfsdk:"status" query:"status,optional"`
	Tags      *[]types.String `tfsdk:"tags" query:"tags,optional"`
	Type      types.String    `tfsdk:"type" query:"type,optional"`
	VMID      types.String    `tfsdk:"vm_id" query:"vm_id,optional"`
}
