// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectDataSourceModel struct {
	ID        types.String                                              `tfsdk:"id" path:"project_id,computed"`
	ProjectID types.String                                              `tfsdk:"project_id" path:"project_id,required"`
	CreatedAt timetypes.RFC3339                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                                              `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UserID    types.String                                              `tfsdk:"user_id" json:"user_id,computed"`
	Tags      customfield.List[types.String]                            `tfsdk:"tags" json:"tags,computed"`
	Resources customfield.NestedObject[ProjectResourcesDataSourceModel] `tfsdk:"resources" json:"resources,computed"`
}

type ProjectResourcesDataSourceModel struct {
	Blockchain customfield.NestedObject[ProjectResourcesBlockchainDataSourceModel] `tfsdk:"blockchain" json:"blockchain,computed"`
	Cloud      customfield.NestedObject[ProjectResourcesCloudDataSourceModel]      `tfsdk:"cloud" json:"cloud,computed"`
}

type ProjectResourcesBlockchainDataSourceModel struct {
	RPCNodesDedicated types.Int64 `tfsdk:"rpc_nodes_dedicated" json:"rpc_nodes_dedicated,computed"`
	RPCNodesFlex      types.Int64 `tfsdk:"rpc_nodes_flex" json:"rpc_nodes_flex,computed"`
}

type ProjectResourcesCloudDataSourceModel struct {
	ConnectConnections types.Int64 `tfsdk:"connect_connections" json:"connect_connections,computed"`
	VMs                types.Int64 `tfsdk:"vms" json:"vms,computed"`
	Volumes            types.Int64 `tfsdk:"volumes" json:"volumes,computed"`
	VPCs               types.Int64 `tfsdk:"vpcs" json:"vpcs,computed"`
}
