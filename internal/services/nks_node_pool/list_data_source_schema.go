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

var _ datasource.DataSourceWithConfigValidators = (*NKSNodePoolsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				Required: true,
			},
			"instance_type": schema.StringAttribute{
				Description: "Filter by the instance type the pool's nodes run",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by a case-insensitive substring of the node pool name",
				Optional:    true,
			},
			"node_count_max": schema.Int64Attribute{
				Description: "Only pools with at most this many nodes",
				Optional:    true,
			},
			"node_count_min": schema.Int64Attribute{
				Description: "Only pools with at least this many nodes",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Filter by node pool status\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"error\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"creating",
						"updating",
						"ready",
						"deleting",
						"error",
					),
				},
			},
			"tags": schema.ListAttribute{
				Description: "Filter by tags. Repeat the parameter to require several tags; a node pool must carry all of them.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"sort": schema.StringAttribute{
				Description: "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status, node_count",
				Computed:    true,
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
				CustomType:  customfield.NewNestedObjectListType[NKSNodePoolsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the node pool.",
							Computed:    true,
						},
						"cluster_id": schema.StringAttribute{
							Description: "ID of the Cluster this node pool belongs to.",
							Computed:    true,
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
						"node_config": schema.SingleNestedAttribute{
							Description: "Node configuration.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[NKSNodePoolsNodeConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"boot_volume": schema.SingleNestedAttribute{
									Description: "Boot volume configuration.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[NKSNodePoolsNodeConfigBootVolumeDataSourceModel](ctx),
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
								"labels": schema.ListAttribute{
									Description: `Kubernetes labels applied to each node in the pool. Each entry is "key=value".`,
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"taints": schema.ListAttribute{
									Description: `Kubernetes taints applied to each node in the pool. Each entry is "key=value:Effect".`,
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
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
						"tags": schema.ListAttribute{
							Description: "Tags attached to the node pool.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the node pool was last updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *NKSNodePoolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *NKSNodePoolsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
