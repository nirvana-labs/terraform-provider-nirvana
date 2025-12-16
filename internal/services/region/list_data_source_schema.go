// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package region

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*RegionsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[RegionsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Name of the region.",
							Computed:    true,
						},
						"availability": schema.StringAttribute{
							Description: "Availability status of the region.\nAvailable values: \"live\", \"preview\", \"maintenance\", \"sunset\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"live",
									"preview",
									"maintenance",
									"sunset",
								),
							},
						},
						"compute": schema.SingleNestedAttribute{
							Description: "Compute products available in this region.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RegionsComputeDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"vms": schema.BoolAttribute{
									Description: "VMs indicates if Virtual Machines are available.",
									Computed:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Name of the region.",
							Computed:    true,
						},
						"networking": schema.SingleNestedAttribute{
							Description: "Networking products available in this region.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RegionsNetworkingDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connect": schema.BoolAttribute{
									Description: "Connect indicates if Nirvana Connect is available.",
									Computed:    true,
								},
								"vpcs": schema.BoolAttribute{
									Description: "VPCs indicates if Virtual Private Clouds are available.",
									Computed:    true,
								},
							},
						},
						"storage": schema.SingleNestedAttribute{
							Description: "Storage products available in this region.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RegionsStorageDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"abs": schema.BoolAttribute{
									Description: "ABS indicates if Accelerated Block Storage is available.",
									Computed:    true,
								},
								"local_nvme": schema.BoolAttribute{
									Description: "LocalNVMe indicates if locally-attached NVMe storage is available.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *RegionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RegionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
