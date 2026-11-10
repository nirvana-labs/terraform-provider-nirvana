// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StorageKeyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"key_id": schema.StringAttribute{
				Optional: true,
			},
			"access_key": schema.StringAttribute{
				Description: "S3 access key ID.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Object Storage Key was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Object Storage Key.",
				Computed:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "Project ID the Object Storage Key belongs to.",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sva-2\", \"us-wdc-1\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("us-sva-2", "us-wdc-1"),
				},
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
			"updated_at": schema.StringAttribute{
				Description: "When the Object Storage Key was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"project_id": schema.StringAttribute{
						Description: "Project ID of resources to request",
						Required:    true,
					},
					"name": schema.StringAttribute{
						Description: "Filter by a case-insensitive substring of the Object Storage Key name",
						Optional:    true,
					},
					"region": schema.StringAttribute{
						Description: "Filter by region",
						Optional:    true,
					},
					"sort": schema.StringAttribute{
						Description: "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *StorageKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StorageKeyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("key_id"), path.MatchRoot("find_one_by")),
	}
}
