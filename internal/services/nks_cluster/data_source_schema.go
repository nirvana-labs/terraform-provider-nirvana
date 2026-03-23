// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*NKSClusterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cluster_id": schema.StringAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Cluster was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Cluster.",
				Computed:    true,
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
				Description: "Region the resource is in.\nAvailable values: \"us-sva-1\", \"us-sva-2\", \"us-chi-1\", \"us-wdc-1\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sva-1",
						"us-sva-2",
						"us-chi-1",
						"us-wdc-1",
					),
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
			"updated_at": schema.StringAttribute{
				Description: "When the Cluster was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"vpc_id": schema.StringAttribute{
				Description: "ID of the VPC the Cluster is in.",
				Computed:    true,
			},
			"pool_ids": schema.ListAttribute{
				Description: "IDs of pools belonging to this Cluster.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"tags": schema.ListAttribute{
				Description: "Tags attached to the Cluster.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"project_id": schema.StringAttribute{
						Description: "Project ID of resources to request",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *NKSClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NKSClusterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("cluster_id"), path.MatchRoot("find_one_by")),
	}
}
