// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type StorageKeyModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	ProjectID types.String      `tfsdk:"project_id" json:"project_id,required"`
	Region    types.String      `tfsdk:"region" json:"region,required"`
	Name      types.String      `tfsdk:"name" json:"name,required"`
	AccessKey types.String      `tfsdk:"access_key" json:"access_key,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Secret    types.String      `tfsdk:"secret" json:"secret,computed,no_refresh"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m StorageKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StorageKeyModel) MarshalJSONForUpdate(state StorageKeyModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
