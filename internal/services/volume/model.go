// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volume

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/apijson"
)

type VolumeModel struct {
	ID         types.String `tfsdk:"id" json:"id,computed"`
	VMID       types.String `tfsdk:"vm_id" path:"vm_id,required"`
	Size       types.Int64  `tfsdk:"size" json:"size,required"`
	Kind       types.String `tfsdk:"kind" json:"kind,computed"`
	ResourceID types.String `tfsdk:"resource_id" json:"resource_id,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	Type       types.String `tfsdk:"type" json:"type,computed"`
}

func (m VolumeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m VolumeModel) MarshalJSONForUpdate(state VolumeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
