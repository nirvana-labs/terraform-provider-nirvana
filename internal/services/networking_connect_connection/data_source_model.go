// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingConnectConnectionDataSourceModel struct {
	ID               types.String                                                            `tfsdk:"id" path:"connection_id,computed"`
	ConnectionID     types.String                                                            `tfsdk:"connection_id" path:"connection_id,required"`
	ASN              types.Int64                                                             `tfsdk:"asn" json:"asn,computed"`
	BandwidthMbps    types.Int64                                                             `tfsdk:"bandwidth_mbps" json:"bandwidth_mbps,computed"`
	CreatedAt        timetypes.RFC3339                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name             types.String                                                            `tfsdk:"name" json:"name,computed"`
	ProviderASN      types.Int64                                                             `tfsdk:"provider_asn" json:"provider_asn,computed"`
	ProviderRouterIP types.String                                                            `tfsdk:"provider_router_ip" json:"provider_router_ip,computed"`
	Region           types.String                                                            `tfsdk:"region" json:"region,computed"`
	RouterIP         types.String                                                            `tfsdk:"router_ip" json:"router_ip,computed"`
	Status           types.String                                                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt        timetypes.RFC3339                                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	CIDRs            customfield.List[types.String]                                          `tfsdk:"cidrs" json:"cidrs,computed"`
	ProviderCIDRs    customfield.List[types.String]                                          `tfsdk:"provider_cidrs" json:"provider_cidrs,computed"`
	Tags             customfield.List[types.String]                                          `tfsdk:"tags" json:"tags,computed"`
	AWS              customfield.NestedObject[NetworkingConnectConnectionAWSDataSourceModel] `tfsdk:"aws" json:"aws,computed"`
}

type NetworkingConnectConnectionAWSDataSourceModel struct {
	Region types.String `tfsdk:"region" json:"region,computed"`
}
