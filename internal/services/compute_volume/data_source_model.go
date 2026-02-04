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
	params = compute.VolumeListParams{}

	if !m.FindOneBy.ProjectID.IsNull() {
		params.ProjectID = param.NewOpt(m.FindOneBy.ProjectID.ValueString())
	}

	return
}

type ComputeVolumeFindOneByDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" query:"project_id,optional"`
}
