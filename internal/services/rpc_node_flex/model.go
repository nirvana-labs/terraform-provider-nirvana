// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
)

type RPCNodeFlexModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	Blockchain types.String      `tfsdk:"blockchain" json:"blockchain,required"`
	Network    types.String      `tfsdk:"network" json:"network,required"`
	Name       types.String      `tfsdk:"name" json:"name,required"`
	Tags       *[]types.String   `tfsdk:"tags" json:"tags,optional"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint   types.String      `tfsdk:"endpoint" json:"endpoint,computed"`
	UpdatedAt  timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m RPCNodeFlexModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RPCNodeFlexModel) MarshalJSONForUpdate(state RPCNodeFlexModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
