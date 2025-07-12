// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeVolumeDataSourceModel struct {
	VolumeID  types.String      `tfsdk:"volume_id" path:"volume_id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Kind      types.String      `tfsdk:"kind" json:"kind,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Size      types.Int64       `tfsdk:"size" json:"size,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VMID      types.String      `tfsdk:"vm_id" json:"vm_id,computed"`
	VMName    types.String      `tfsdk:"vm_name" json:"vm_name,computed"`
}
