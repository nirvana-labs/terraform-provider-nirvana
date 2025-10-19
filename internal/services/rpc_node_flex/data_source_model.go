// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type RPCNodeFlexDataSourceModel struct {
	ID         types.String                   `tfsdk:"id" path:"node_id,computed"`
	NodeID     types.String                   `tfsdk:"node_id" path:"node_id,required"`
	Blockchain types.String                   `tfsdk:"blockchain" json:"blockchain,computed"`
	CreatedAt  timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint   types.String                   `tfsdk:"endpoint" json:"endpoint,computed"`
	Name       types.String                   `tfsdk:"name" json:"name,computed"`
	Network    types.String                   `tfsdk:"network" json:"network,computed"`
	UpdatedAt  timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags       customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
}
