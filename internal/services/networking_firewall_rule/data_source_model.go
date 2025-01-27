// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingFirewallRuleDataSourceModel struct {
	FirewallRuleID types.String                                                               `tfsdk:"firewall_rule_id" path:"firewall_rule_id,required"`
	VPCID          types.String                                                               `tfsdk:"vpc_id" path:"vpc_id,computed"`
	CreatedAt      types.String                                                               `tfsdk:"created_at" json:"created_at,computed"`
	ID             types.String                                                               `tfsdk:"id" json:"id,computed"`
	Name           types.String                                                               `tfsdk:"name" json:"name,computed"`
	Protocol       types.String                                                               `tfsdk:"protocol" json:"protocol,computed"`
	Status         types.String                                                               `tfsdk:"status" json:"status,computed"`
	UpdatedAt      types.String                                                               `tfsdk:"updated_at" json:"updated_at,computed"`
	Destination    customfield.NestedObject[NetworkingFirewallRuleDestinationDataSourceModel] `tfsdk:"destination" json:"destination,computed"`
	Source         customfield.NestedObject[NetworkingFirewallRuleSourceDataSourceModel]      `tfsdk:"source" json:"source,computed"`
}

type NetworkingFirewallRuleDestinationDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type NetworkingFirewallRuleSourceDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}
