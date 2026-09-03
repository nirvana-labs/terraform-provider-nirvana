// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_node_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSNodePoolDataSourceModel struct {
	ID         types.String                                                   `tfsdk:"id" path:"pool_id,computed"`
	PoolID     types.String                                                   `tfsdk:"pool_id" path:"pool_id,optional"`
	ClusterID  types.String                                                   `tfsdk:"cluster_id" path:"cluster_id,required"`
	CreatedAt  timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name       types.String                                                   `tfsdk:"name" json:"name,computed"`
	NodeCount  types.Int64                                                    `tfsdk:"node_count" json:"node_count,computed"`
	Status     types.String                                                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt  timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags       customfield.List[types.String]                                 `tfsdk:"tags" json:"tags,computed"`
	NodeConfig customfield.NestedObject[NKSNodePoolNodeConfigDataSourceModel] `tfsdk:"node_config" json:"node_config,computed"`
	FindOneBy  *NKSNodePoolFindOneByDataSourceModel                           `tfsdk:"find_one_by"`
}

func (m *NKSNodePoolDataSourceModel) toListParams(_ context.Context) (params nks.ClusterPoolListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = nks.ClusterPoolListParams{
		Tags: mFindOneByTags,
	}

	if !m.FindOneBy.InstanceType.IsNull() {
		params.InstanceType = param.NewOpt(m.FindOneBy.InstanceType.ValueString())
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.NodeCountMax.IsNull() {
		params.NodeCountMax = param.NewOpt(m.FindOneBy.NodeCountMax.ValueInt64())
	}
	if !m.FindOneBy.NodeCountMin.IsNull() {
		params.NodeCountMin = param.NewOpt(m.FindOneBy.NodeCountMin.ValueInt64())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = nks.ClusterPoolListParamsStatus(m.FindOneBy.Status.ValueString())
	}

	return
}

type NKSNodePoolNodeConfigDataSourceModel struct {
	BootVolume   customfield.NestedObject[NKSNodePoolNodeConfigBootVolumeDataSourceModel] `tfsdk:"boot_volume" json:"boot_volume,computed"`
	InstanceType types.String                                                             `tfsdk:"instance_type" json:"instance_type,computed"`
	Labels       customfield.List[types.String]                                           `tfsdk:"labels" json:"labels,computed"`
	Taints       customfield.List[types.String]                                           `tfsdk:"taints" json:"taints,computed"`
}

type NKSNodePoolNodeConfigBootVolumeDataSourceModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type NKSNodePoolFindOneByDataSourceModel struct {
	InstanceType types.String    `tfsdk:"instance_type" query:"instance_type,optional"`
	Name         types.String    `tfsdk:"name" query:"name,optional"`
	NodeCountMax types.Int64     `tfsdk:"node_count_max" query:"node_count_max,optional"`
	NodeCountMin types.Int64     `tfsdk:"node_count_min" query:"node_count_min,optional"`
	Sort         types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status       types.String    `tfsdk:"status" query:"status,optional"`
	Tags         *[]types.String `tfsdk:"tags" query:"tags,optional"`
}
