// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

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

var _ datasource.DataSourceWithConfigValidators = (*NKSClustersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "Project ID of resources to request",
				Required:    true,
			},
			"autoscaling": schema.BoolAttribute{
				Description: "Filter by whether autoscaling is enabled",
				Optional:    true,
			},
			"kubernetes_version": schema.StringAttribute{
				Description: "Filter by Kubernetes version, matched exactly",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by a case-insensitive substring of the Cluster name",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "Filter by region",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Filter by Cluster status\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"error\".",
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
			"vpc_id": schema.StringAttribute{
				Description: "Filter by the VPC the Cluster is in",
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Filter by tags. Repeat the parameter to require several tags; a Cluster must carry all of them.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"sort": schema.StringAttribute{
				Description: "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status",
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
				CustomType:  customfield.NewNestedObjectListType[NKSClustersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the Cluster.",
							Computed:    true,
						},
						"autoscaling": schema.BoolAttribute{
							Description: "Whether autoscaling is enabled for the Cluster.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "When the Cluster was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"kubernetes_version": schema.StringAttribute{
							Description: "Kubernetes version of the Cluster.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the Cluster.",
							Computed:    true,
						},
						"pool_ids": schema.ListAttribute{
							Description: "IDs of pools belonging to this Cluster.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"private_ip": schema.StringAttribute{
							Description: "Private IP (VIP) of the Cluster.",
							Computed:    true,
						},
						"project_id": schema.StringAttribute{
							Description: "Project ID the Cluster belongs to.",
							Computed:    true,
						},
						"public_ip": schema.StringAttribute{
							Description: "Public IP of the Cluster.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Region the resource is in.\nAvailable values: \"us-sva-2\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("us-sva-2"),
							},
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
							Description: "Tags attached to the Cluster.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the Cluster was last updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"vpc_id": schema.StringAttribute{
							Description: "ID of the VPC the Cluster is in.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *NKSClustersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *NKSClustersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
