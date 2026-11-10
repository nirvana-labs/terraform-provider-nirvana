// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket

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

var _ datasource.DataSourceWithConfigValidators = (*StorageBucketsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "Project ID of resources to request",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by a case-insensitive substring of the Bucket name",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "Filter by region",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Filter by Bucket status\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"error\".",
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
				Description: "Filter by tags. Repeat the parameter to require several tags; a Bucket must carry all of them.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"sort": schema.StringAttribute{
				Description: "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status, size_bytes, object_count",
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
				CustomType:  customfield.NewNestedObjectListType[StorageBucketsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the Bucket.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "When the Bucket was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"endpoint": schema.StringAttribute{
							Description: "S3 endpoint of the Bucket's region. Point S3 clients at it, with an Object\nStorage Key of the same project and region.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the Bucket, used to address it over S3.",
							Computed:    true,
						},
						"object_count": schema.Int64Attribute{
							Description: "Number of objects in the Bucket. Updated periodically, so it can lag behind\nrecent writes.",
							Computed:    true,
						},
						"project_id": schema.StringAttribute{
							Description: "Project ID the Bucket belongs to.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Region the resource is in.\nAvailable values: \"us-sva-2\", \"us-wdc-1\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("us-sva-2", "us-wdc-1"),
							},
						},
						"signing_region": schema.StringAttribute{
							Description: "Region S3 clients must sign their requests for (the SigV4 region). Configure\nit as the region of any S3 client or SDK used with the Bucket.",
							Computed:    true,
						},
						"size_bytes": schema.Int64Attribute{
							Description: "Total size of the objects in the Bucket, in bytes. Updated periodically,\nso it can lag behind recent writes.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the resource.\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"error\", \"suspended\".",
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
									"suspended",
								),
							},
						},
						"tags": schema.ListAttribute{
							Description: "Tags attached to the Bucket.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the Bucket was updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"url": schema.StringAttribute{
							Description: "Path-style URL of the Bucket.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *StorageBucketsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *StorageBucketsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
