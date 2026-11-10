// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type StorageBucketModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	Name          types.String      `tfsdk:"name" json:"name,required"`
	ProjectID     types.String      `tfsdk:"project_id" json:"project_id,required"`
	Region        types.String      `tfsdk:"region" json:"region,required"`
	Tags          *[]types.String   `tfsdk:"tags" json:"tags,optional"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint      types.String      `tfsdk:"endpoint" json:"endpoint,computed"`
	ObjectCount   types.Int64       `tfsdk:"object_count" json:"object_count,computed"`
	SigningRegion types.String      `tfsdk:"signing_region" json:"signing_region,computed"`
	SizeBytes     types.Int64       `tfsdk:"size_bytes" json:"size_bytes,computed"`
	Status        types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	URL           types.String      `tfsdk:"url" json:"url,computed"`
}

func (m StorageBucketModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StorageBucketModel) MarshalJSONForUpdate(state StorageBucketModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
