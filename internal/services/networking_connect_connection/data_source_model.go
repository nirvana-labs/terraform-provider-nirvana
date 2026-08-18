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

type NetworkingConnectConnectionDataSourceModel struct {
	ID               types.String                                                            `tfsdk:"id" path:"connection_id,computed"`
	ConnectionID     types.String                                                            `tfsdk:"connection_id" path:"connection_id,optional"`
	ASN              types.Int64                                                             `tfsdk:"asn" json:"asn,computed"`
	BandwidthMbps    types.Int64                                                             `tfsdk:"bandwidth_mbps" json:"bandwidth_mbps,computed"`
	CreatedAt        timetypes.RFC3339                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name             types.String                                                            `tfsdk:"name" json:"name,computed"`
	ProjectID        types.String                                                            `tfsdk:"project_id" json:"project_id,computed"`
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
	FindOneBy        *NetworkingConnectConnectionFindOneByDataSourceModel                    `tfsdk:"find_one_by"`
}

func (m *NetworkingConnectConnectionDataSourceModel) toListParams(_ context.Context) (params networking.ConnectConnectionListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = networking.ConnectConnectionListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
		Tags:      mFindOneByTags,
	}

	if !m.FindOneBy.BandwidthMbps.IsNull() {
		params.BandwidthMbps = int64(m.FindOneBy.BandwidthMbps.ValueInt64())
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.NetworkingConnectConnectionProvider.IsNull() {
		params.Provider = param.NewOpt(m.FindOneBy.NetworkingConnectConnectionProvider.ValueString())
	}
	if !m.FindOneBy.ProviderRegion.IsNull() {
		params.ProviderRegion = param.NewOpt(m.FindOneBy.ProviderRegion.ValueString())
	}
	if !m.FindOneBy.Region.IsNull() {
		params.Region = param.NewOpt(m.FindOneBy.Region.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = networking.ConnectConnectionListParamsStatus(m.FindOneBy.Status.ValueString())
	}

	return
}

type NetworkingConnectConnectionAWSDataSourceModel struct {
	Region types.String `tfsdk:"region" json:"region,computed"`
}

type NetworkingConnectConnectionFindOneByDataSourceModel struct {
	ProjectID                           types.String    `tfsdk:"project_id" query:"project_id,required"`
	BandwidthMbps                       types.Int64     `tfsdk:"bandwidth_mbps" query:"bandwidth_mbps,optional"`
	Name                                types.String    `tfsdk:"name" query:"name,optional"`
	NetworkingConnectConnectionProvider types.String    `tfsdk:"networking_connect_connection_provider" query:"provider,optional"`
	ProviderRegion                      types.String    `tfsdk:"provider_region" query:"provider_region,optional"`
	Region                              types.String    `tfsdk:"region" query:"region,optional"`
	Sort                                types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status                              types.String    `tfsdk:"status" query:"status,optional"`
	Tags                                *[]types.String `tfsdk:"tags" query:"tags,optional"`
}
