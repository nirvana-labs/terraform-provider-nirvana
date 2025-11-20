// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*RPCNodeFlexResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the RPC Node Flex.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"blockchain": schema.StringAttribute{
				Description:   "Blockchain.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"network": schema.StringAttribute{
				Description:   "Network type (e.g., mainnet, testnet).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the RPC Node Flex.",
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the RPC Node Flex (optional, max 50).",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the RPC Node Flex was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"endpoint": schema.StringAttribute{
				Description: "RPC endpoint URL.",
				Computed:    true,
				Sensitive:   true,
			},
			"project_id": schema.StringAttribute{
				Description: "Project identifier associated with the RPC Node Flex.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "When the RPC Node Flex was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *RPCNodeFlexResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *RPCNodeFlexResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
