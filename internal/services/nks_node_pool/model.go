// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_node_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type NKSNodePoolModel struct {
	ID         types.String                `tfsdk:"id" json:"id,computed"`
	ClusterID  types.String                `tfsdk:"cluster_id" path:"cluster_id,required"`
	NodeConfig *NKSNodePoolNodeConfigModel `tfsdk:"node_config" json:"node_config,required"`
	Name       types.String                `tfsdk:"name" json:"name,required"`
	NodeCount  types.Int64                 `tfsdk:"node_count" json:"node_count,required"`
	Tags       *[]types.String             `tfsdk:"tags" json:"tags,optional"`
	CreatedAt  timetypes.RFC3339           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ProjectID  types.String                `tfsdk:"project_id" json:"project_id,computed,no_refresh"`
	Status     types.String                `tfsdk:"status" json:"status,computed"`
	UpdatedAt  timetypes.RFC3339           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m NKSNodePoolModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NKSNodePoolModel) MarshalJSONForUpdate(state NKSNodePoolModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NKSNodePoolNodeConfigModel struct {
	BootVolume   *NKSNodePoolNodeConfigBootVolumeModel   `tfsdk:"boot_volume" json:"boot_volume,required"`
	CPUConfig    *NKSNodePoolNodeConfigCPUConfigModel    `tfsdk:"cpu_config" json:"cpu_config,required"`
	MemoryConfig *NKSNodePoolNodeConfigMemoryConfigModel `tfsdk:"memory_config" json:"memory_config,required"`
}

type NKSNodePoolNodeConfigBootVolumeModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,required"`
	Type types.String `tfsdk:"type" json:"type,required"`
}

type NKSNodePoolNodeConfigCPUConfigModel struct {
	Vcpu types.Int64 `tfsdk:"vcpu" json:"vcpu,required"`
}

type NKSNodePoolNodeConfigMemoryConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}
