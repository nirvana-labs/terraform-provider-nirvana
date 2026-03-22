// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_node_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSNodePoolsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NKSNodePoolsItemsDataSourceModel] `json:"items,computed"`
}

type NKSNodePoolsDataSourceModel struct {
	ClusterID types.String                                                   `tfsdk:"cluster_id" path:"cluster_id,required"`
	MaxItems  types.Int64                                                    `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[NKSNodePoolsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NKSNodePoolsDataSourceModel) toListParams(_ context.Context) (params nks.ClusterPoolListParams, diags diag.Diagnostics) {
	params = nks.ClusterPoolListParams{}

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
	BootVolume   customfield.NestedObject[NKSNodePoolsNodeConfigBootVolumeDataSourceModel]   `tfsdk:"boot_volume" json:"boot_volume,computed"`
	CPUConfig    customfield.NestedObject[NKSNodePoolsNodeConfigCPUConfigDataSourceModel]    `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemoryConfig customfield.NestedObject[NKSNodePoolsNodeConfigMemoryConfigDataSourceModel] `tfsdk:"memory_config" json:"memory_config,computed"`
}

type NKSNodePoolsNodeConfigBootVolumeDataSourceModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type NKSNodePoolsNodeConfigCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type NKSNodePoolsNodeConfigMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
