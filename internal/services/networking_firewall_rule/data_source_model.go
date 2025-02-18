// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingFirewallRuleDataSourceModel struct {
	FirewallRuleID     types.String                   `tfsdk:"firewall_rule_id" path:"firewall_rule_id,required"`
	VPCID              types.String                   `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt          types.String                   `tfsdk:"created_at" json:"created_at,computed"`
	DestinationAddress types.String                   `tfsdk:"destination_address" json:"destination_address,computed"`
	ID                 types.String                   `tfsdk:"id" json:"id,computed"`
	Name               types.String                   `tfsdk:"name" json:"name,computed"`
	Protocol           types.String                   `tfsdk:"protocol" json:"protocol,computed"`
	SourceAddress      types.String                   `tfsdk:"source_address" json:"source_address,computed"`
	Status             types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt          types.String                   `tfsdk:"updated_at" json:"updated_at,computed"`
	DestinationPorts   customfield.List[types.String] `tfsdk:"destination_ports" json:"destination_ports,computed"`
	SourcePorts        customfield.List[types.String] `tfsdk:"source_ports" json:"source_ports,computed"`
}
