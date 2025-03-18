// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type NetworkingFirewallRuleModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	VPCID              types.String      `tfsdk:"vpc_id" path:"vpc_id,required"`
	DestinationAddress types.String      `tfsdk:"destination_address" json:"destination_address,required"`
	Name               types.String      `tfsdk:"name" json:"name,required"`
	Protocol           types.String      `tfsdk:"protocol" json:"protocol,required"`
	SourceAddress      types.String      `tfsdk:"source_address" json:"source_address,required"`
	DestinationPorts   *[]types.String   `tfsdk:"destination_ports" json:"destination_ports,required"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status             types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m NetworkingFirewallRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkingFirewallRuleModel) MarshalJSONForUpdate(state NetworkingFirewallRuleModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
