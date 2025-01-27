// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCDataSourceModel struct {
	VPCID         types.String                                                            `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt     types.String                                                            `tfsdk:"created_at" json:"created_at,computed"`
	ID            types.String                                                            `tfsdk:"id" json:"id,computed"`
	Name          types.String                                                            `tfsdk:"name" json:"name,computed"`
	Region        types.String                                                            `tfsdk:"region" json:"region,computed"`
	Status        types.String                                                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt     types.String                                                            `tfsdk:"updated_at" json:"updated_at,computed"`
	FirewallRules customfield.NestedObjectList[NetworkingVPCFirewallRulesDataSourceModel] `tfsdk:"firewall_rules" json:"firewall_rules,computed"`
	Subnet        customfield.NestedObject[NetworkingVPCSubnetDataSourceModel]            `tfsdk:"subnet" json:"subnet,computed"`
}

type NetworkingVPCFirewallRulesDataSourceModel struct {
	ID          types.String                                                                   `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String                                                                   `tfsdk:"created_at" json:"created_at,computed"`
	Destination customfield.NestedObject[NetworkingVPCFirewallRulesDestinationDataSourceModel] `tfsdk:"destination" json:"destination,computed"`
	Name        types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Protocol    types.String                                                                   `tfsdk:"protocol" json:"protocol,computed"`
	Source      customfield.NestedObject[NetworkingVPCFirewallRulesSourceDataSourceModel]      `tfsdk:"source" json:"source,computed"`
	Status      types.String                                                                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                                                   `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID       types.String                                                                   `tfsdk:"vpc_id" json:"vpc_id,computed"`
}

type NetworkingVPCFirewallRulesDestinationDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type NetworkingVPCFirewallRulesSourceDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type NetworkingVPCSubnetDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
