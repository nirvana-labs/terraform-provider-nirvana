// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

type FirewallRuleDataSourceModel struct {
	FirewallRuleID types.String                                                `tfsdk:"firewall_rule_id" path:"firewall_rule_id,required"`
	VPCID          types.String                                                `tfsdk:"vpc_id" path:"vpc_id,computed"`
	CreatedAt      types.String                                                `tfsdk:"created_at" json:"created_at,computed"`
	ID             types.String                                                `tfsdk:"id" json:"id,computed"`
	Name           types.String                                                `tfsdk:"name" json:"name,computed"`
	Protocol       types.String                                                `tfsdk:"protocol" json:"protocol,computed"`
	Status         types.String                                                `tfsdk:"status" json:"status,computed"`
	UpdatedAt      types.String                                                `tfsdk:"updated_at" json:"updated_at,computed"`
	Dest           customfield.NestedObject[FirewallRuleDestDataSourceModel]   `tfsdk:"dest" json:"dest,computed"`
	Source         customfield.NestedObject[FirewallRuleSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

type FirewallRuleDestDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type FirewallRuleSourceDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}
