// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/storage"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type StorageKeysItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[StorageKeysItemsDataSourceModel] `json:"items,computed"`
}

type StorageKeysDataSourceModel struct {
	ProjectID types.String                                                  `tfsdk:"project_id" query:"project_id,required"`
	Name      types.String                                                  `tfsdk:"name" query:"name,optional"`
	Region    types.String                                                  `tfsdk:"region" query:"region,optional"`
	Sort      types.String                                                  `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems  types.Int64                                                   `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[StorageKeysItemsDataSourceModel] `tfsdk:"items"`
}

func (m *StorageKeysDataSourceModel) toListParams(_ context.Context) (params storage.KeyListParams, diags diag.Diagnostics) {
	params = storage.KeyListParams{
		ProjectID: m.ProjectID.ValueString(),
	}

	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}

	return
}

type StorageKeysItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccessKey types.String      `tfsdk:"access_key" json:"access_key,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	ProjectID types.String      `tfsdk:"project_id" json:"project_id,computed"`
	Region    types.String      `tfsdk:"region" json:"region,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
