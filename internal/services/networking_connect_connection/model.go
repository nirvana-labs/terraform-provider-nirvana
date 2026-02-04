// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type NetworkingConnectConnectionModel struct {
	ID               types.String                         `tfsdk:"id" json:"id,computed"`
	BandwidthMbps    types.Int64                          `tfsdk:"bandwidth_mbps" json:"bandwidth_mbps,required"`
	Region           types.String                         `tfsdk:"region" json:"region,required"`
	CIDRs            *[]types.String                      `tfsdk:"cidrs" json:"cidrs,required"`
	ProviderCIDRs    *[]types.String                      `tfsdk:"provider_cidrs" json:"provider_cidrs,required"`
	ProjectID        types.String                         `tfsdk:"project_id" json:"project_id,optional"`
	AWS              *NetworkingConnectConnectionAWSModel `tfsdk:"aws" json:"aws,optional"`
	Name             types.String                         `tfsdk:"name" json:"name,required"`
	Tags             *[]types.String                      `tfsdk:"tags" json:"tags,optional"`
	ASN              types.Int64                          `tfsdk:"asn" json:"asn,computed"`
	CreatedAt        timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ProviderASN      types.Int64                          `tfsdk:"provider_asn" json:"provider_asn,computed"`
	ProviderRouterIP types.String                         `tfsdk:"provider_router_ip" json:"provider_router_ip,computed"`
	RouterIP         types.String                         `tfsdk:"router_ip" json:"router_ip,computed"`
	Status           types.String                         `tfsdk:"status" json:"status,computed"`
	UpdatedAt        timetypes.RFC3339                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m NetworkingConnectConnectionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkingConnectConnectionModel) MarshalJSONForUpdate(state NetworkingConnectConnectionModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NetworkingConnectConnectionAWSModel struct {
	AccountID types.String `tfsdk:"account_id" json:"account_id,required,no_refresh"`
	Region    types.String `tfsdk:"region" json:"region,required"`
}
