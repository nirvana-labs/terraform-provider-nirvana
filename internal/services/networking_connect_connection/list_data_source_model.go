// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingConnectConnectionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NetworkingConnectConnectionsItemsDataSourceModel] `json:"items,computed"`
}

type NetworkingConnectConnectionsDataSourceModel struct {
	MaxItems types.Int64                                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[NetworkingConnectConnectionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NetworkingConnectConnectionsDataSourceModel) toListParams(_ context.Context) (params networking.ConnectConnectionListParams, diags diag.Diagnostics) {
	params = networking.ConnectConnectionListParams{}

	return
}

type NetworkingConnectConnectionsItemsDataSourceModel struct {
	ID               types.String                                                             `tfsdk:"id" json:"id,computed"`
	ASN              types.Int64                                                              `tfsdk:"asn" json:"asn,computed"`
	AWS              customfield.NestedObject[NetworkingConnectConnectionsAWSDataSourceModel] `tfsdk:"aws" json:"aws,computed"`
	BandwidthMbps    types.Int64                                                              `tfsdk:"bandwidth_mbps" json:"bandwidth_mbps,computed"`
	CIDRs            customfield.List[types.String]                                           `tfsdk:"cidrs" json:"cidrs,computed"`
	CreatedAt        timetypes.RFC3339                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name             types.String                                                             `tfsdk:"name" json:"name,computed"`
	ProviderASN      types.Int64                                                              `tfsdk:"provider_asn" json:"provider_asn,computed"`
	ProviderCIDRs    customfield.List[types.String]                                           `tfsdk:"provider_cidrs" json:"provider_cidrs,computed"`
	ProviderRouterIP types.String                                                             `tfsdk:"provider_router_ip" json:"provider_router_ip,computed"`
	Region           types.String                                                             `tfsdk:"region" json:"region,computed"`
	RouterIP         types.String                                                             `tfsdk:"router_ip" json:"router_ip,computed"`
	Status           types.String                                                             `tfsdk:"status" json:"status,computed"`
	Tags             customfield.List[types.String]                                           `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt        timetypes.RFC3339                                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type NetworkingConnectConnectionsAWSDataSourceModel struct {
	Region types.String `tfsdk:"region" json:"region,computed"`
}
