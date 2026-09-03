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

type NKSNodePoolsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NKSNodePoolsItemsDataSourceModel] `json:"items,computed"`
}

type NKSNodePoolsDataSourceModel struct {
	ClusterID    types.String                                                   `tfsdk:"cluster_id" path:"cluster_id,required"`
	InstanceType types.String                                                   `tfsdk:"instance_type" query:"instance_type,optional"`
	Name         types.String                                                   `tfsdk:"name" query:"name,optional"`
	NodeCountMax types.Int64                                                    `tfsdk:"node_count_max" query:"node_count_max,optional"`
	NodeCountMin types.Int64                                                    `tfsdk:"node_count_min" query:"node_count_min,optional"`
	Status       types.String                                                   `tfsdk:"status" query:"status,optional"`
	Tags         *[]types.String                                                `tfsdk:"tags" query:"tags,optional"`
	Sort         types.String                                                   `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems     types.Int64                                                    `tfsdk:"max_items"`
	Items        customfield.NestedObjectList[NKSNodePoolsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NKSNodePoolsDataSourceModel) toListParams(_ context.Context) (params nks.ClusterPoolListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = nks.ClusterPoolListParams{
		Tags: mTags,
	}

	if !m.InstanceType.IsNull() {
		params.InstanceType = param.NewOpt(m.InstanceType.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.NodeCountMax.IsNull() {
		params.NodeCountMax = param.NewOpt(m.NodeCountMax.ValueInt64())
	}
	if !m.NodeCountMin.IsNull() {
		params.NodeCountMin = param.NewOpt(m.NodeCountMin.ValueInt64())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = nks.ClusterPoolListParamsStatus(m.Status.ValueString())
	}

	return
}

type NKSNodePoolsItemsDataSourceModel struct {
	ID         types.String                                                    `tfsdk:"id" json:"id,computed"`
	ClusterID  types.String                                                    `tfsdk:"cluster_id" json:"cluster_id,computed"`
	CreatedAt  timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name       types.String                                                    `tfsdk:"name" json:"name,computed"`
	NodeConfig customfield.NestedObject[NKSNodePoolsNodeConfigDataSourceModel] `tfsdk:"node_config" json:"node_config,computed"`
	NodeCount  types.Int64                                                     `tfsdk:"node_count" json:"node_count,computed"`
	Status     types.String                                                    `tfsdk:"status" json:"status,computed"`
	Tags       customfield.List[types.String]                                  `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt  timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type NKSNodePoolsNodeConfigDataSourceModel struct {
	BootVolume   customfield.NestedObject[NKSNodePoolsNodeConfigBootVolumeDataSourceModel] `tfsdk:"boot_volume" json:"boot_volume,computed"`
	InstanceType types.String                                                              `tfsdk:"instance_type" json:"instance_type,computed"`
	Labels       customfield.List[types.String]                                            `tfsdk:"labels" json:"labels,computed"`
	Taints       customfield.List[types.String]                                            `tfsdk:"taints" json:"taints,computed"`
}

type NKSNodePoolsNodeConfigBootVolumeDataSourceModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
