// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectModel struct {
	ID        types.String                                    `tfsdk:"id" json:"id,computed"`
	Name      types.String                                    `tfsdk:"name" json:"name,required"`
	Tags      *[]types.String                                 `tfsdk:"tags" json:"tags,optional"`
	CreatedAt timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UserID    types.String                                    `tfsdk:"user_id" json:"user_id,computed"`
	Resources customfield.NestedObject[ProjectResourcesModel] `tfsdk:"resources" json:"resources,computed"`
}

func (m ProjectModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProjectModel) MarshalJSONForUpdate(state ProjectModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ProjectResourcesModel struct {
	Blockchain customfield.NestedObject[ProjectResourcesBlockchainModel] `tfsdk:"blockchain" json:"blockchain,computed"`
	Cloud      customfield.NestedObject[ProjectResourcesCloudModel]      `tfsdk:"cloud" json:"cloud,computed"`
}

type ProjectResourcesBlockchainModel struct {
	RPCNodesDedicated types.Int64 `tfsdk:"rpc_nodes_dedicated" json:"rpc_nodes_dedicated,computed"`
	RPCNodesFlex      types.Int64 `tfsdk:"rpc_nodes_flex" json:"rpc_nodes_flex,computed"`
}

type ProjectResourcesCloudModel struct {
	ConnectConnections types.Int64 `tfsdk:"connect_connections" json:"connect_connections,computed"`
	VMs                types.Int64 `tfsdk:"vms" json:"vms,computed"`
	Volumes            types.Int64 `tfsdk:"volumes" json:"volumes,computed"`
	VPCs               types.Int64 `tfsdk:"vpcs" json:"vpcs,computed"`
}
