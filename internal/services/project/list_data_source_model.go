// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/projects"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ProjectsItemsDataSourceModel] `json:"items,computed"`
}

type ProjectsDataSourceModel struct {
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ProjectsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ProjectsDataSourceModel) toListParams(_ context.Context) (params projects.ProjectListParams, diags diag.Diagnostics) {
	params = projects.ProjectListParams{}

	return
}

type ProjectsItemsDataSourceModel struct {
	ID        types.String                                               `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                                               `tfsdk:"name" json:"name,computed"`
	Resources customfield.NestedObject[ProjectsResourcesDataSourceModel] `tfsdk:"resources" json:"resources,computed"`
	Tags      customfield.List[types.String]                             `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt timetypes.RFC3339                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UserID    types.String                                               `tfsdk:"user_id" json:"user_id,computed"`
}

type ProjectsResourcesDataSourceModel struct {
	Blockchain customfield.NestedObject[ProjectsResourcesBlockchainDataSourceModel] `tfsdk:"blockchain" json:"blockchain,computed"`
	Cloud      customfield.NestedObject[ProjectsResourcesCloudDataSourceModel]      `tfsdk:"cloud" json:"cloud,computed"`
}

type ProjectsResourcesBlockchainDataSourceModel struct {
	RPCNodesDedicated types.Int64 `tfsdk:"rpc_nodes_dedicated" json:"rpc_nodes_dedicated,computed"`
	RPCNodesFlex      types.Int64 `tfsdk:"rpc_nodes_flex" json:"rpc_nodes_flex,computed"`
}

type ProjectsResourcesCloudDataSourceModel struct {
	ConnectConnections types.Int64 `tfsdk:"connect_connections" json:"connect_connections,computed"`
	VMs                types.Int64 `tfsdk:"vms" json:"vms,computed"`
	Volumes            types.Int64 `tfsdk:"volumes" json:"volumes,computed"`
	VPCs               types.Int64 `tfsdk:"vpcs" json:"vpcs,computed"`
}
