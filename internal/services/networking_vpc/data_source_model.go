// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCDataSourceModel struct {
	VPCID           types.String                                                 `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt       timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID              types.String                                                 `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                 `tfsdk:"name" json:"name,computed"`
	Region          types.String                                                 `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                 `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                            `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	FirewallRuleIDs customfield.List[types.String]                               `tfsdk:"firewall_rule_ids" json:"firewall_rule_ids,computed"`
	Subnet          customfield.NestedObject[NetworkingVPCSubnetDataSourceModel] `tfsdk:"subnet" json:"subnet,computed"`
}

type NetworkingVPCSubnetDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CIDR      types.String      `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
