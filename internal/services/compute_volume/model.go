// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type ComputeVolumeModel struct {
	ID         types.String `tfsdk:"id" json:"id,computed"`
	VMID       types.String `tfsdk:"vm_id" json:"vm_id,required"`
	Size       types.Int64  `tfsdk:"size" json:"size,required"`
	CreatedAt  types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind       types.String `tfsdk:"kind" json:"kind,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
	ResourceID types.String `tfsdk:"resource_id" json:"resource_id,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	Type       types.String `tfsdk:"type" json:"type,computed"`
	UpdatedAt  types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

func (m ComputeVolumeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeVolumeModel) MarshalJSONForUpdate(state ComputeVolumeModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
