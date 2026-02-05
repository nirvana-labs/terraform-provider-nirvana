// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/rpc_nodes"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type RPCNodeFlexesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[RPCNodeFlexesItemsDataSourceModel] `json:"items,computed"`
}

type RPCNodeFlexesDataSourceModel struct {
	ProjectID types.String                                                    `tfsdk:"project_id" query:"project_id,required"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[RPCNodeFlexesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *RPCNodeFlexesDataSourceModel) toListParams(_ context.Context) (params rpc_nodes.FlexListParams, diags diag.Diagnostics) {
	params = rpc_nodes.FlexListParams{
		ProjectID: m.ProjectID.ValueString(),
	}

	return
}

type RPCNodeFlexesItemsDataSourceModel struct {
	ID         types.String                   `tfsdk:"id" json:"id,computed"`
	Blockchain types.String                   `tfsdk:"blockchain" json:"blockchain,computed"`
	CreatedAt  timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint   types.String                   `tfsdk:"endpoint" json:"endpoint,computed"`
	Name       types.String                   `tfsdk:"name" json:"name,computed"`
	Network    types.String                   `tfsdk:"network" json:"network,computed"`
	ProjectID  types.String                   `tfsdk:"project_id" json:"project_id,computed"`
	Tags       customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt  timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
