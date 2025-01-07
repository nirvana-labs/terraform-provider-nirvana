// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

type VPCDataSourceModel struct {
	VPCID         types.String                                                  `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt     types.String                                                  `tfsdk:"created_at" json:"created_at,computed"`
	ID            types.String                                                  `tfsdk:"id" json:"id,computed"`
	Name          types.String                                                  `tfsdk:"name" json:"name,computed"`
	Region        types.String                                                  `tfsdk:"region" json:"region,computed"`
	Status        types.String                                                  `tfsdk:"status" json:"status,computed"`
	UpdatedAt     types.String                                                  `tfsdk:"updated_at" json:"updated_at,computed"`
	FirewallRules customfield.NestedObjectList[VPCFirewallRulesDataSourceModel] `tfsdk:"firewall_rules" json:"firewall_rules,computed"`
	Subnet        customfield.NestedObject[VPCSubnetDataSourceModel]            `tfsdk:"subnet" json:"subnet,computed"`
}

type VPCFirewallRulesDataSourceModel struct {
	ID        types.String                                                    `tfsdk:"id" json:"id,computed"`
	CreatedAt types.String                                                    `tfsdk:"created_at" json:"created_at,computed"`
	Dest      customfield.NestedObject[VPCFirewallRulesDestDataSourceModel]   `tfsdk:"dest" json:"dest,computed"`
	Name      types.String                                                    `tfsdk:"name" json:"name,computed"`
	Protocol  types.String                                                    `tfsdk:"protocol" json:"protocol,computed"`
	Source    customfield.NestedObject[VPCFirewallRulesSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
	Status    types.String                                                    `tfsdk:"status" json:"status,computed"`
	UpdatedAt types.String                                                    `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID     types.String                                                    `tfsdk:"vpc_id" json:"vpc_id,computed"`
}

type VPCFirewallRulesDestDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VPCFirewallRulesSourceDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VPCSubnetDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
