// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*RPCNodeFlexDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"node_id": schema.StringAttribute{
				Required: true,
			},
			"blockchain": schema.StringAttribute{
				Description: "Blockchain type.",
				Computed:    true,
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
			"name": schema.StringAttribute{
				Description: "Name of the RPC Node Flex.",
				Computed:    true,
			},
			"network": schema.StringAttribute{
				Description: "Network type (e.g., mainnet, testnet).",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "When the RPC Node Flex was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the RPC Node Flex.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *RPCNodeFlexDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RPCNodeFlexDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
