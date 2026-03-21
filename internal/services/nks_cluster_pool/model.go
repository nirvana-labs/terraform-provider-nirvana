// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type NKSClusterPoolModel struct {
	ID         types.String                   `tfsdk:"id" json:"id,computed"`
	ClusterID  types.String                   `tfsdk:"cluster_id" path:"cluster_id,required"`
	NodeCount  types.Int64                    `tfsdk:"node_count" json:"node_count,required"`
	NodeConfig *NKSClusterPoolNodeConfigModel `tfsdk:"node_config" json:"node_config,required"`
	Name       types.String                   `tfsdk:"name" json:"name,required"`
	Tags       *[]types.String                `tfsdk:"tags" json:"tags,optional"`
	CreatedAt  timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ProjectID  types.String                   `tfsdk:"project_id" json:"project_id,computed,no_refresh"`
	Status     types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt  timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m NKSClusterPoolModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NKSClusterPoolModel) MarshalJSONForUpdate(state NKSClusterPoolModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NKSClusterPoolNodeConfigModel struct {
	BootVolume   *NKSClusterPoolNodeConfigBootVolumeModel   `tfsdk:"boot_volume" json:"boot_volume,required"`
	CPUConfig    *NKSClusterPoolNodeConfigCPUConfigModel    `tfsdk:"cpu_config" json:"cpu_config,required"`
	MemoryConfig *NKSClusterPoolNodeConfigMemoryConfigModel `tfsdk:"memory_config" json:"memory_config,required"`
}

type NKSClusterPoolNodeConfigBootVolumeModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,required"`
	Type types.String `tfsdk:"type" json:"type,required"`
}

type NKSClusterPoolNodeConfigCPUConfigModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,required"`
}

type NKSClusterPoolNodeConfigMemoryConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}
