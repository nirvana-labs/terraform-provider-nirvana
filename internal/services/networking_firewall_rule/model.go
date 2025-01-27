// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type NetworkingFirewallRuleModel struct {
	ID          types.String                            `tfsdk:"id" json:"id,computed"`
	VPCID       types.String                            `tfsdk:"vpc_id" path:"vpc_id,required"`
	Name        types.String                            `tfsdk:"name" json:"name,required"`
	Protocol    types.String                            `tfsdk:"protocol" json:"protocol,required"`
	Destination *NetworkingFirewallRuleDestinationModel `tfsdk:"destination" json:"destination,required"`
	Source      *NetworkingFirewallRuleSourceModel      `tfsdk:"source" json:"source,required"`
	CreatedAt   types.String                            `tfsdk:"created_at" json:"created_at,computed"`
	Kind        types.String                            `tfsdk:"kind" json:"kind,computed"`
	ResourceID  types.String                            `tfsdk:"resource_id" json:"resource_id,computed"`
	Status      types.String                            `tfsdk:"status" json:"status,computed"`
	Type        types.String                            `tfsdk:"type" json:"type,computed"`
	UpdatedAt   types.String                            `tfsdk:"updated_at" json:"updated_at,computed"`
}

func (m NetworkingFirewallRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkingFirewallRuleModel) MarshalJSONForUpdate(state NetworkingFirewallRuleModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NetworkingFirewallRuleDestinationModel struct {
	Address types.String    `tfsdk:"address" json:"address,optional"`
	Ports   *[]types.String `tfsdk:"ports" json:"ports,optional"`
}

type NetworkingFirewallRuleSourceModel struct {
	Address types.String    `tfsdk:"address" json:"address,optional"`
	Ports   *[]types.String `tfsdk:"ports" json:"ports,optional"`
}
