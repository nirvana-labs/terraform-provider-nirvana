// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type APIKeyModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,required" format:"date-time"`
	StartsAt  timetypes.RFC3339 `tfsdk:"starts_at" json:"starts_at,optional" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Key       types.String      `tfsdk:"key" json:"key,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m APIKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIKeyModel) MarshalJSONForUpdate(state APIKeyModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
