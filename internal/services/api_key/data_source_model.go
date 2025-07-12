// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIKeyDataSourceModel struct {
	APIKeyID  types.String      `tfsdk:"api_key_id" path:"api_key_id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Key       types.String      `tfsdk:"key" json:"key,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartsAt  timetypes.RFC3339 `tfsdk:"starts_at" json:"starts_at,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UserID    types.String      `tfsdk:"user_id" json:"user_id,computed"`
}
