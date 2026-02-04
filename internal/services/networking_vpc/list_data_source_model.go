// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NetworkingVPCsItemsDataSourceModel] `json:"items,computed"`
}

type NetworkingVPCsDataSourceModel struct {
	ProjectID types.String                                                     `tfsdk:"project_id" query:"project_id,optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[NetworkingVPCsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NetworkingVPCsDataSourceModel) toListParams(_ context.Context) (params networking.VPCListParams, diags diag.Diagnostics) {
	params = networking.VPCListParams{}

	if !m.ProjectID.IsNull() {
		params.ProjectID = param.NewOpt(m.ProjectID.ValueString())
	}

	return
}

type NetworkingVPCsItemsDataSourceModel struct {
	ID              types.String                                                  `tfsdk:"id" json:"id,computed"`
	CreatedAt       timetypes.RFC3339                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	FirewallRuleIDs customfield.List[types.String]                                `tfsdk:"firewall_rule_ids" json:"firewall_rule_ids,computed"`
	Name            types.String                                                  `tfsdk:"name" json:"name,computed"`
	ProjectID       types.String                                                  `tfsdk:"project_id" json:"project_id,computed"`
	Region          types.String                                                  `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                  `tfsdk:"status" json:"status,computed"`
	Subnet          customfield.NestedObject[NetworkingVPCsSubnetDataSourceModel] `tfsdk:"subnet" json:"subnet,computed"`
	Tags            customfield.List[types.String]                                `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt       timetypes.RFC3339                                             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type NetworkingVPCsSubnetDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CIDR      types.String      `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
