// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*RPCNodeFlexesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "Project ID of resources to request",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[RPCNodeFlexesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the RPC Node Flex.",
							Computed:    true,
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
						"project_id": schema.StringAttribute{
							Description: "Project identifier associated with the RPC Node Flex.",
							Computed:    true,
						},
						"tags": schema.ListAttribute{
							Description: "Tags to attach to the RPC Node Flex.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the RPC Node Flex was updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *RPCNodeFlexesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RPCNodeFlexesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
