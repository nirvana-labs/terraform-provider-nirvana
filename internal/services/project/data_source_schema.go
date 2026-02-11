// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ProjectDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Project was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Project name.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "When the Project was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"user_id": schema.StringAttribute{
				Description: "User ID that owns the project.",
				Computed:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags attached to the Project.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"resources": schema.SingleNestedAttribute{
				Description: "Resource counts for the project.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ProjectResourcesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"blockchain": schema.SingleNestedAttribute{
						Description: "Blockchain resources.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ProjectResourcesBlockchainDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"rpc_nodes_dedicated": schema.Int64Attribute{
								Description: "Number of dedicated RPC nodes in the project.",
								Computed:    true,
							},
							"rpc_nodes_flex": schema.Int64Attribute{
								Description: "Number of flex RPC nodes in the project.",
								Computed:    true,
							},
						},
					},
					"cloud": schema.SingleNestedAttribute{
						Description: "Cloud infrastructure resources.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ProjectResourcesCloudDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"connect_connections": schema.Int64Attribute{
								Description: "Number of Connect connections in the project.",
								Computed:    true,
							},
							"vms": schema.Int64Attribute{
								Description: "Number of VMs in the project.",
								Computed:    true,
							},
							"volumes": schema.Int64Attribute{
								Description: "Number of volumes in the project.",
								Computed:    true,
							},
							"vpcs": schema.Int64Attribute{
								Description: "Number of VPCs in the project.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *ProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ProjectDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
