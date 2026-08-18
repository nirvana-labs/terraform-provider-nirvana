// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingVPCDataSourceModel struct {
	ID              types.String                                                 `tfsdk:"id" path:"vpc_id,computed"`
	VPCID           types.String                                                 `tfsdk:"vpc_id" path:"vpc_id,optional"`
	CreatedAt       timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name            types.String                                                 `tfsdk:"name" json:"name,computed"`
	ProjectID       types.String                                                 `tfsdk:"project_id" json:"project_id,computed"`
	Region          types.String                                                 `tfsdk:"region" json:"region,computed"`
	Status          types.String                                                 `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                            `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	FirewallRuleIDs customfield.List[types.String]                               `tfsdk:"firewall_rule_ids" json:"firewall_rule_ids,computed"`
	Tags            customfield.List[types.String]                               `tfsdk:"tags" json:"tags,computed"`
	Subnet          customfield.NestedObject[NetworkingVPCSubnetDataSourceModel] `tfsdk:"subnet" json:"subnet,computed"`
	FindOneBy       *NetworkingVPCFindOneByDataSourceModel                       `tfsdk:"find_one_by"`
}

func (m *NetworkingVPCDataSourceModel) toListParams(_ context.Context) (params networking.VPCListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = networking.VPCListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
		Tags:      mFindOneByTags,
	}

	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Region.IsNull() {
		params.Region = param.NewOpt(m.FindOneBy.Region.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = networking.VPCListParamsStatus(m.FindOneBy.Status.ValueString())
	}

	return
}

type NetworkingVPCSubnetDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CIDR      types.String      `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type NetworkingVPCFindOneByDataSourceModel struct {
	ProjectID types.String    `tfsdk:"project_id" query:"project_id,required"`
	Name      types.String    `tfsdk:"name" query:"name,optional"`
	Region    types.String    `tfsdk:"region" query:"region,optional"`
	Sort      types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status    types.String    `tfsdk:"status" query:"status,optional"`
	Tags      *[]types.String `tfsdk:"tags" query:"tags,optional"`
}
