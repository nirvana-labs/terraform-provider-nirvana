// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/storage"
)

type StorageKeyDataSourceModel struct {
	ID        types.String                        `tfsdk:"id" path:"key_id,computed"`
	KeyID     types.String                        `tfsdk:"key_id" path:"key_id,optional"`
	AccessKey types.String                        `tfsdk:"access_key" json:"access_key,computed"`
	CreatedAt timetypes.RFC3339                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                        `tfsdk:"name" json:"name,computed"`
	ProjectID types.String                        `tfsdk:"project_id" json:"project_id,computed"`
	Region    types.String                        `tfsdk:"region" json:"region,computed"`
	Status    types.String                        `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	FindOneBy *StorageKeyFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *StorageKeyDataSourceModel) toListParams(_ context.Context) (params storage.KeyListParams, diags diag.Diagnostics) {
	params = storage.KeyListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
	}

	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Region.IsNull() {
		params.Region = param.NewOpt(m.FindOneBy.Region.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}

	return
}

type StorageKeyFindOneByDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" query:"project_id,required"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
	Region    types.String `tfsdk:"region" query:"region,optional"`
	Sort      types.String `tfsdk:"sort" query:"sort,computed_optional"`
}
