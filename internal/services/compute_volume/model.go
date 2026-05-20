// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVolumeModel struct {
	ID        types.String                                        `tfsdk:"id" json:"id,computed"`
	ProjectID types.String                                        `tfsdk:"project_id" json:"project_id,required"`
	Region    types.String                                        `tfsdk:"region" json:"region,required"`
	Type      types.String                                        `tfsdk:"type" json:"type,required"`
	VMID      types.String                                        `tfsdk:"vm_id" json:"vm_id,optional"`
	Name      types.String                                        `tfsdk:"name" json:"name,required"`
	Size      types.Int64                                         `tfsdk:"size" json:"size,required"`
	Tags      *[]types.String                                     `tfsdk:"tags" json:"tags,optional"`
	CreatedAt timetypes.RFC3339                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind      types.String                                        `tfsdk:"kind" json:"kind,computed"`
	Status    types.String                                        `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VMName    types.String                                        `tfsdk:"vm_name" json:"vm_name,computed"`
	Details   customfield.NestedObject[ComputeVolumeDetailsModel] `tfsdk:"details" json:"details,computed,no_refresh"`
}

func (m ComputeVolumeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeVolumeModel) MarshalJSONForUpdate(state ComputeVolumeModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ComputeVolumeDetailsModel struct {
	Changes customfield.NestedObjectMap[ComputeVolumeDetailsChangesModel] `tfsdk:"changes" json:"changes,computed"`
}

type ComputeVolumeDetailsChangesModel struct {
	From types.String `tfsdk:"from" json:"from,computed"`
	To   types.String `tfsdk:"to" json:"to,computed"`
}
