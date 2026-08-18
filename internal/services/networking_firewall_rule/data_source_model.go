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

type NetworkingFirewallRuleDataSourceModel struct {
	ID                 types.String                                    `tfsdk:"id" path:"firewall_rule_id,computed"`
	FirewallRuleID     types.String                                    `tfsdk:"firewall_rule_id" path:"firewall_rule_id,optional"`
	VPCID              types.String                                    `tfsdk:"vpc_id" path:"vpc_id,required"`
	CreatedAt          timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DestinationAddress types.String                                    `tfsdk:"destination_address" json:"destination_address,computed"`
	Name               types.String                                    `tfsdk:"name" json:"name,computed"`
	Protocol           types.String                                    `tfsdk:"protocol" json:"protocol,computed"`
	SourceAddress      types.String                                    `tfsdk:"source_address" json:"source_address,computed"`
	Status             types.String                                    `tfsdk:"status" json:"status,computed"`
	UpdatedAt          timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	DestinationPorts   customfield.List[types.String]                  `tfsdk:"destination_ports" json:"destination_ports,computed"`
	Tags               customfield.List[types.String]                  `tfsdk:"tags" json:"tags,computed"`
	FindOneBy          *NetworkingFirewallRuleFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *NetworkingFirewallRuleDataSourceModel) toListParams(_ context.Context) (params networking.FirewallRuleListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = networking.FirewallRuleListParams{
		Tags: mFindOneByTags,
	}

	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Protocol.IsNull() {
		params.Protocol = networking.FirewallRuleListParamsProtocol(m.FindOneBy.Protocol.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = networking.FirewallRuleListParamsStatus(m.FindOneBy.Status.ValueString())
	}

	return
}

type NetworkingFirewallRuleFindOneByDataSourceModel struct {
	Name     types.String    `tfsdk:"name" query:"name,optional"`
	Protocol types.String    `tfsdk:"protocol" query:"protocol,optional"`
	Sort     types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status   types.String    `tfsdk:"status" query:"status,optional"`
	Tags     *[]types.String `tfsdk:"tags" query:"tags,optional"`
}
