// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/storage"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type StorageBucketsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[StorageBucketsItemsDataSourceModel] `json:"items,computed"`
}

type StorageBucketsDataSourceModel struct {
	ProjectID types.String                                                     `tfsdk:"project_id" query:"project_id,required"`
	Name      types.String                                                     `tfsdk:"name" query:"name,optional"`
	Region    types.String                                                     `tfsdk:"region" query:"region,optional"`
	Status    types.String                                                     `tfsdk:"status" query:"status,optional"`
	Tags      *[]types.String                                                  `tfsdk:"tags" query:"tags,optional"`
	Sort      types.String                                                     `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[StorageBucketsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *StorageBucketsDataSourceModel) toListParams(_ context.Context) (params storage.BucketListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = storage.BucketListParams{
		ProjectID: m.ProjectID.ValueString(),
		Tags:      mTags,
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
	if !m.Status.IsNull() {
		params.Status = storage.BucketListParamsStatus(m.Status.ValueString())
	}

	return
}

type StorageBucketsItemsDataSourceModel struct {
	ID            types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint      types.String                   `tfsdk:"endpoint" json:"endpoint,computed"`
	Name          types.String                   `tfsdk:"name" json:"name,computed"`
	ObjectCount   types.Int64                    `tfsdk:"object_count" json:"object_count,computed"`
	ProjectID     types.String                   `tfsdk:"project_id" json:"project_id,computed"`
	Region        types.String                   `tfsdk:"region" json:"region,computed"`
	SigningRegion types.String                   `tfsdk:"signing_region" json:"signing_region,computed"`
	SizeBytes     types.Int64                    `tfsdk:"size_bytes" json:"size_bytes,computed"`
	Status        types.String                   `tfsdk:"status" json:"status,computed"`
	Tags          customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt     timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	URL           types.String                   `tfsdk:"url" json:"url,computed"`
}
