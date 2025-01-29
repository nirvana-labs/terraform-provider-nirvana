// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCDataSourceModel struct {
	VPCID           types.String                                                 `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt       types.String                                                 `tfsdk:"created_at" json:"created_at,computed"`
	ID              types.String                                                 `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                 `tfsdk:"name" json:"name,computed"`
	Region          types.String                                                 `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                 `tfsdk:"status" json:"status,computed"`
	UpdatedAt       types.String                                                 `tfsdk:"updated_at" json:"updated_at,computed"`
	FirewallRuleIDs customfield.List[types.String]                               `tfsdk:"firewall_rule_ids" json:"firewall_rule_ids,computed"`
	Subnet          customfield.NestedObject[NetworkingVPCSubnetDataSourceModel] `tfsdk:"subnet" json:"subnet,computed"`
}

type NetworkingVPCSubnetDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
