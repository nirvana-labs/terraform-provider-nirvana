// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/api_keys"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type APIKeysItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[APIKeysItemsDataSourceModel] `json:"items,computed"`
}

type APIKeysDataSourceModel struct {
	Name     types.String                                              `tfsdk:"name" query:"name,optional"`
	Status   types.String                                              `tfsdk:"status" query:"status,optional"`
	Tags     *[]types.String                                           `tfsdk:"tags" query:"tags,optional"`
	Sort     types.String                                              `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems types.Int64                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[APIKeysItemsDataSourceModel] `tfsdk:"items"`
}

func (m *APIKeysDataSourceModel) toListParams(_ context.Context) (params api_keys.APIKeyListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = api_keys.APIKeyListParams{
		Tags: mTags,
	}

	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = api_keys.APIKeyListParamsStatus(m.Status.ValueString())
	}

	return
}

type APIKeysItemsDataSourceModel struct {
	ID           types.String                                                    `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt    timetypes.RFC3339                                               `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Managed      types.Bool                                                      `tfsdk:"managed" json:"managed,computed"`
	Name         types.String                                                    `tfsdk:"name" json:"name,computed"`
	Permissions  customfield.NestedObjectList[APIKeysPermissionsDataSourceModel] `tfsdk:"permissions" json:"permissions,computed"`
	ProjectIDs   customfield.List[types.String]                                  `tfsdk:"project_ids" json:"project_ids,computed"`
	SourceIPRule customfield.NestedObject[APIKeysSourceIPRuleDataSourceModel]    `tfsdk:"source_ip_rule" json:"source_ip_rule,computed"`
	Status       types.String                                                    `tfsdk:"status" json:"status,computed"`
	Tags         customfield.List[types.String]                                  `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt    timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Key          types.String                                                    `tfsdk:"key" json:"key,computed"`
	StartsAt     timetypes.RFC3339                                               `tfsdk:"starts_at" json:"starts_at,computed" format:"date-time"`
}

type APIKeysPermissionsDataSourceModel struct {
	Permission   types.String `tfsdk:"permission" json:"permission,computed"`
	ResourceType types.String `tfsdk:"resource_type" json:"resource_type,computed"`
}

type APIKeysSourceIPRuleDataSourceModel struct {
	Allowed customfield.List[types.String] `tfsdk:"allowed" json:"allowed,computed"`
	Blocked customfield.List[types.String] `tfsdk:"blocked" json:"blocked,computed"`
}
