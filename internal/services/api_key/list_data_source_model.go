// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/api_keys"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type APIKeysItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[APIKeysItemsDataSourceModel] `json:"items,computed"`
}

type APIKeysDataSourceModel struct {
	MaxItems types.Int64                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[APIKeysItemsDataSourceModel] `tfsdk:"items"`
}

func (m *APIKeysDataSourceModel) toListParams(_ context.Context) (params api_keys.APIKeyListParams, diags diag.Diagnostics) {
	params = api_keys.APIKeyListParams{}

	return
}

type APIKeysItemsDataSourceModel struct {
	ID           types.String                                                 `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt    timetypes.RFC3339                                            `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Name         types.String                                                 `tfsdk:"name" json:"name,computed"`
	SourceIPRule customfield.NestedObject[APIKeysSourceIPRuleDataSourceModel] `tfsdk:"source_ip_rule" json:"source_ip_rule,computed"`
	Status       types.String                                                 `tfsdk:"status" json:"status,computed"`
	Tags         customfield.List[types.String]                               `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt    timetypes.RFC3339                                            `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Key          types.String                                                 `tfsdk:"key" json:"key,computed"`
	StartsAt     timetypes.RFC3339                                            `tfsdk:"starts_at" json:"starts_at,computed" format:"date-time"`
}

type APIKeysSourceIPRuleDataSourceModel struct {
	Allowed customfield.List[types.String] `tfsdk:"allowed" json:"allowed,computed"`
	Blocked customfield.List[types.String] `tfsdk:"blocked" json:"blocked,computed"`
}
