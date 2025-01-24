// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volume

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VolumeDataSourceModel struct {
	VolumeID  types.String `tfsdk:"volume_id" path:"volume_id,required"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Size      types.Int64  `tfsdk:"size" json:"size,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
