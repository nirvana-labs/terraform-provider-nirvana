// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_node_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*NKSNodePoolDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"pool_id": schema.StringAttribute{
				Required: true,
			},
			"cluster_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the node pool was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the node pool.",
				Computed:    true,
			},
			"node_count": schema.Int64Attribute{
				Description: "Number of nodes.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the resource.\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"creating",
						"updating",
						"ready",
						"deleting",
						"deleted",
						"error",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the node pool was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"tags": schema.ListAttribute{
				Description: "Tags attached to the node pool.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"node_config": schema.SingleNestedAttribute{
				Description: "Node configuration.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NKSNodePoolNodeConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"boot_volume": schema.SingleNestedAttribute{
						Description: "Boot volume configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[NKSNodePoolNodeConfigBootVolumeDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"size": schema.Int64Attribute{
								Description: "Size of the boot volume in GB.",
								Computed:    true,
								Validators: []validator.Int64{
									int64validator.Between(64, 512),
								},
							},
							"type": schema.StringAttribute{
								Description: "Type of the Volume.\nAvailable values: \"nvme\", \"abs\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("nvme", "abs"),
								},
							},
						},
					},
					"instance_type": schema.StringAttribute{
						Description: "Instance type name.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *NKSNodePoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NKSNodePoolDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
