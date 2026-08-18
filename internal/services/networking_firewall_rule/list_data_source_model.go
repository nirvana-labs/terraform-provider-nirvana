// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingFirewallRulesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NetworkingFirewallRulesItemsDataSourceModel] `json:"items,computed"`
}

type NetworkingFirewallRulesDataSourceModel struct {
	VPCID    types.String                                                              `tfsdk:"vpc_id" path:"vpc_id,required"`
	Name     types.String                                                              `tfsdk:"name" query:"name,optional"`
	Protocol types.String                                                              `tfsdk:"protocol" query:"protocol,optional"`
	Status   types.String                                                              `tfsdk:"status" query:"status,optional"`
	Tags     *[]types.String                                                           `tfsdk:"tags" query:"tags,optional"`
	Sort     types.String                                                              `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems types.Int64                                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[NetworkingFirewallRulesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NetworkingFirewallRulesDataSourceModel) toListParams(_ context.Context) (params networking.FirewallRuleListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = networking.FirewallRuleListParams{
		Tags: mTags,
	}

	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Protocol.IsNull() {
		params.Protocol = networking.FirewallRuleListParamsProtocol(m.Protocol.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = networking.FirewallRuleListParamsStatus(m.Status.ValueString())
	}

	return
}

type NetworkingFirewallRulesItemsDataSourceModel struct {
	ID                 types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt          timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DestinationAddress types.String                   `tfsdk:"destination_address" json:"destination_address,computed"`
	DestinationPorts   customfield.List[types.String] `tfsdk:"destination_ports" json:"destination_ports,computed"`
	Name               types.String                   `tfsdk:"name" json:"name,computed"`
	Protocol           types.String                   `tfsdk:"protocol" json:"protocol,computed"`
	SourceAddress      types.String                   `tfsdk:"source_address" json:"source_address,computed"`
	Status             types.String                   `tfsdk:"status" json:"status,computed"`
	Tags               customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt          timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID              types.String                   `tfsdk:"vpc_id" json:"vpc_id,computed"`
}
