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
