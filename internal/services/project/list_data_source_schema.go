// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

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

var _ datasource.DataSourceWithConfigValidators = (*ProjectsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
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
				CustomType:  customfield.NewNestedObjectListType[ProjectsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Project ID.",
							Computed:    true,
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
						"resources": schema.SingleNestedAttribute{
							Description: "Resource counts for the project.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ProjectsResourcesDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"blockchain": schema.SingleNestedAttribute{
									Description: "Blockchain resources.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ProjectsResourcesBlockchainDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[ProjectsResourcesCloudDataSourceModel](ctx),
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
						"tags": schema.ListAttribute{
							Description: "Tags attached to the Project.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
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
					},
				},
			},
		},
	}
}

func (d *ProjectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ProjectsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
