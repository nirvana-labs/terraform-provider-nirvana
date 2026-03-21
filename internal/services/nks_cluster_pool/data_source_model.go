// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSClusterPoolDataSourceModel struct {
	ID         types.String                                                      `tfsdk:"id" path:"pool_id,computed"`
	PoolID     types.String                                                      `tfsdk:"pool_id" path:"pool_id,required"`
	ClusterID  types.String                                                      `tfsdk:"cluster_id" path:"cluster_id,required"`
	CreatedAt  timetypes.RFC3339                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name       types.String                                                      `tfsdk:"name" json:"name,computed"`
	NodeCount  types.Int64                                                       `tfsdk:"node_count" json:"node_count,computed"`
	Status     types.String                                                      `tfsdk:"status" json:"status,computed"`
	UpdatedAt  timetypes.RFC3339                                                 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags       customfield.List[types.String]                                    `tfsdk:"tags" json:"tags,computed"`
	NodeConfig customfield.NestedObject[NKSClusterPoolNodeConfigDataSourceModel] `tfsdk:"node_config" json:"node_config,computed"`
}

type NKSClusterPoolNodeConfigDataSourceModel struct {
	BootVolume   customfield.NestedObject[NKSClusterPoolNodeConfigBootVolumeDataSourceModel]   `tfsdk:"boot_volume" json:"boot_volume,computed"`
	CPUConfig    customfield.NestedObject[NKSClusterPoolNodeConfigCPUConfigDataSourceModel]    `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemoryConfig customfield.NestedObject[NKSClusterPoolNodeConfigMemoryConfigDataSourceModel] `tfsdk:"memory_config" json:"memory_config,computed"`
}

type NKSClusterPoolNodeConfigBootVolumeDataSourceModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type NKSClusterPoolNodeConfigCPUConfigDataSourceModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,computed"`
}

type NKSClusterPoolNodeConfigMemoryConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}
