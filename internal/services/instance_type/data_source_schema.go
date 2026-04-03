// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*InstanceTypeDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Description: `Available values: "us-sva-1", "us-sva-2", "us-chi-1", "us-wdc-1".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sva-1",
						"us-sva-2",
						"us-chi-1",
						"us-wdc-1",
					),
				},
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
	}
}

func (d *InstanceTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *InstanceTypeDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
