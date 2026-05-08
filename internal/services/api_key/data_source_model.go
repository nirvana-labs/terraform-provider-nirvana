// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type APIKeyDataSourceModel struct {
	ID           types.String                                                   `tfsdk:"id" path:"api_key_id,computed"`
	APIKeyID     types.String                                                   `tfsdk:"api_key_id" path:"api_key_id,required"`
	CreatedAt    timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt    timetypes.RFC3339                                              `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Key          types.String                                                   `tfsdk:"key" json:"key,computed"`
	Managed      types.Bool                                                     `tfsdk:"managed" json:"managed,computed"`
	Name         types.String                                                   `tfsdk:"name" json:"name,computed"`
	StartsAt     timetypes.RFC3339                                              `tfsdk:"starts_at" json:"starts_at,computed" format:"date-time"`
	Status       types.String                                                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt    timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ProjectIDs   customfield.List[types.String]                                 `tfsdk:"project_ids" json:"project_ids,computed"`
	Tags         customfield.List[types.String]                                 `tfsdk:"tags" json:"tags,computed"`
	Permissions  customfield.NestedObjectList[APIKeyPermissionsDataSourceModel] `tfsdk:"permissions" json:"permissions,computed"`
	SourceIPRule customfield.NestedObject[APIKeySourceIPRuleDataSourceModel]    `tfsdk:"source_ip_rule" json:"source_ip_rule,computed"`
}

type APIKeyPermissionsDataSourceModel struct {
	Permission   types.String `tfsdk:"permission" json:"permission,computed"`
	ResourceType types.String `tfsdk:"resource_type" json:"resource_type,computed"`
}

type APIKeySourceIPRuleDataSourceModel struct {
	Allowed customfield.List[types.String] `tfsdk:"allowed" json:"allowed,computed"`
	Blocked customfield.List[types.String] `tfsdk:"blocked" json:"blocked,computed"`
}
