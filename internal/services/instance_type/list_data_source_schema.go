// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*InstanceTypesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"chipset": schema.StringAttribute{
				Description: "Filter by chipset",
				Optional:    true,
			},
			"family": schema.StringAttribute{
				Description: "Filter by family",
				Optional:    true,
			},
			"memory_gb_max": schema.Int64Attribute{
				Description: "Only Instance Types with at most this much memory, in GB",
				Optional:    true,
			},
			"memory_gb_min": schema.Int64Attribute{
				Description: "Only Instance Types with at least this much memory, in GB",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by a case-insensitive substring of the Instance Type name",
				Optional:    true,
			},
			"network_bandwidth_gbps_max": schema.Float64Attribute{
				Description: "Only Instance Types with at most this much network bandwidth, in Gbps",
				Optional:    true,
			},
			"network_bandwidth_gbps_min": schema.Float64Attribute{
				Description: "Only Instance Types with at least this much network bandwidth, in Gbps",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "Filter by region",
				Optional:    true,
			},
			"series": schema.StringAttribute{
				Description: "Filter by series",
				Optional:    true,
			},
			"vcpu_max": schema.Int64Attribute{
				Description: "Only Instance Types with at most this many vCPUs",
				Optional:    true,
			},
			"vcpu_min": schema.Int64Attribute{
				Description: "Only Instance Types with at least this many vCPUs",
				Optional:    true,
			},
			"sort": schema.StringAttribute{
				Description: "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: series, family, name, vcpu, memory_gb, network_bandwidth_gbps",
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
				CustomType:  customfield.NewNestedObjectListType[InstanceTypesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"chipset": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Description: "When the Instance Type was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"family": schema.StringAttribute{
							Computed: true,
						},
						"memory_gb": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"network_bandwidth_gbps": schema.Float64Attribute{
							Description: "Network bandwidth in Gbps.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Region the resource is in.\nAvailable values: \"us-sva-2\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("us-sva-2"),
							},
						},
						"series": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the Instance Type was updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"vcpu": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *InstanceTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *InstanceTypesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
