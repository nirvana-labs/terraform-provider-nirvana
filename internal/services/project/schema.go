// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ProjectResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Project ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description: "Project name.",
				Required:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Project.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Project was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"resources": schema.SingleNestedAttribute{
				Description: "Resource counts for the project.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ProjectResourcesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"blockchain": schema.SingleNestedAttribute{
						Description: "Blockchain resources.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ProjectResourcesBlockchainModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[ProjectResourcesCloudModel](ctx),
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

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ProjectResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
