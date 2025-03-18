// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCModel struct {
	ID              types.String                                       `tfsdk:"id" json:"id,computed"`
	Region          types.String                                       `tfsdk:"region" json:"region,required"`
	Name            types.String                                       `tfsdk:"name" json:"name,required"`
	SubnetName      types.String                                       `tfsdk:"subnet_name" json:"subnet_name,required"`
	CreatedAt       timetypes.RFC3339                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status          types.String                                       `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                  `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	FirewallRuleIDs customfield.List[types.String]                     `tfsdk:"firewall_rule_ids" json:"firewall_rule_ids,computed"`
	Subnet          customfield.NestedObject[NetworkingVPCSubnetModel] `tfsdk:"subnet" json:"subnet,computed"`
}

func (m NetworkingVPCModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkingVPCModel) MarshalJSONForUpdate(state NetworkingVPCModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NetworkingVPCSubnetModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Cidr      types.String      `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
