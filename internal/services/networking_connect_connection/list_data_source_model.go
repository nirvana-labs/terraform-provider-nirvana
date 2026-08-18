// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingConnectConnectionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NetworkingConnectConnectionsItemsDataSourceModel] `json:"items,computed"`
}

type NetworkingConnectConnectionsDataSourceModel struct {
	ProjectID                           types.String                                                                   `tfsdk:"project_id" query:"project_id,required"`
	BandwidthMbps                       types.Int64                                                                    `tfsdk:"bandwidth_mbps" query:"bandwidth_mbps,optional"`
	Name                                types.String                                                                   `tfsdk:"name" query:"name,optional"`
	NetworkingConnectConnectionProvider types.String                                                                   `tfsdk:"networking_connect_connection_provider" query:"provider,optional"`
	ProviderRegion                      types.String                                                                   `tfsdk:"provider_region" query:"provider_region,optional"`
	Region                              types.String                                                                   `tfsdk:"region" query:"region,optional"`
	Status                              types.String                                                                   `tfsdk:"status" query:"status,optional"`
	Tags                                *[]types.String                                                                `tfsdk:"tags" query:"tags,optional"`
	Sort                                types.String                                                                   `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems                            types.Int64                                                                    `tfsdk:"max_items"`
	Items                               customfield.NestedObjectList[NetworkingConnectConnectionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NetworkingConnectConnectionsDataSourceModel) toListParams(_ context.Context) (params networking.ConnectConnectionListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = networking.ConnectConnectionListParams{
		ProjectID: m.ProjectID.ValueString(),
		Tags:      mTags,
	}

	if !m.BandwidthMbps.IsNull() {
		params.BandwidthMbps = int64(m.BandwidthMbps.ValueInt64())
	}
	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.NetworkingConnectConnectionProvider.IsNull() {
		params.Provider = param.NewOpt(m.NetworkingConnectConnectionProvider.ValueString())
	}
	if !m.ProviderRegion.IsNull() {
		params.ProviderRegion = param.NewOpt(m.ProviderRegion.ValueString())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = networking.ConnectConnectionListParamsStatus(m.Status.ValueString())
	}

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
	ProjectID        types.String                                                             `tfsdk:"project_id" json:"project_id,computed"`
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
