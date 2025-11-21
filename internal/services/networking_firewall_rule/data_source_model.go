// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingFirewallRuleDataSourceModel struct {
	ID                 types.String                   `tfsdk:"id" path:"firewall_rule_id,computed"`
	FirewallRuleID     types.String                   `tfsdk:"firewall_rule_id" path:"firewall_rule_id,required"`
	VPCID              types.String                   `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt          timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DestinationAddress types.String                   `tfsdk:"destination_address" json:"destination_address,computed"`
	Name               types.String                   `tfsdk:"name" json:"name,computed"`
	Protocol           types.String                   `tfsdk:"protocol" json:"protocol,computed"`
	SourceAddress      types.String                   `tfsdk:"source_address" json:"source_address,computed"`
	Status             types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt          timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	DestinationPorts   customfield.List[types.String] `tfsdk:"destination_ports" json:"destination_ports,computed"`
	Tags               customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
}
