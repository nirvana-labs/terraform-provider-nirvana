// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type VPCModel struct {
	ID            types.String                                        `tfsdk:"id" json:"id,computed"`
	Name          types.String                                        `tfsdk:"name" json:"name,required"`
	Region        types.String                                        `tfsdk:"region" json:"region,required"`
	SubnetName    types.String                                        `tfsdk:"subnet_name" json:"subnet_name,required"`
	CreatedAt     types.String                                        `tfsdk:"created_at" json:"created_at,computed"`
	Kind          types.String                                        `tfsdk:"kind" json:"kind,computed"`
	ResourceID    types.String                                        `tfsdk:"resource_id" json:"resource_id,computed"`
	Status        types.String                                        `tfsdk:"status" json:"status,computed"`
	Type          types.String                                        `tfsdk:"type" json:"type,computed"`
	UpdatedAt     types.String                                        `tfsdk:"updated_at" json:"updated_at,computed"`
	FirewallRules customfield.NestedObjectList[VPCFirewallRulesModel] `tfsdk:"firewall_rules" json:"firewall_rules,computed"`
	Subnet        customfield.NestedObject[VPCSubnetModel]            `tfsdk:"subnet" json:"subnet,computed"`
}

func (m VPCModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m VPCModel) MarshalJSONForUpdate(state VPCModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type VPCFirewallRulesModel struct {
	ID          types.String                                               `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String                                               `tfsdk:"created_at" json:"created_at,computed"`
	Destination customfield.NestedObject[VPCFirewallRulesDestinationModel] `tfsdk:"destination" json:"destination,computed"`
	Name        types.String                                               `tfsdk:"name" json:"name,computed"`
	Protocol    types.String                                               `tfsdk:"protocol" json:"protocol,computed"`
	Source      customfield.NestedObject[VPCFirewallRulesSourceModel]      `tfsdk:"source" json:"source,computed"`
	Status      types.String                                               `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                               `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID       types.String                                               `tfsdk:"vpc_id" json:"vpc_id,computed"`
}

type VPCFirewallRulesDestinationModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VPCFirewallRulesSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VPCSubnetModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
