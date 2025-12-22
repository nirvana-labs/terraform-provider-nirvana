// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

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

var _ datasource.DataSourceWithConfigValidators = (*APIKeysDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[APIKeysItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "API Key ID.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "When the API Key was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"expires_at": schema.StringAttribute{
							Description: "When the API Key expires and is no longer valid.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "API Key name.",
							Computed:    true,
						},
						"source_ip_rule": schema.SingleNestedAttribute{
							Description: "IP filter configuration for the API Key.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[APIKeysSourceIPRuleDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"in": schema.ListAttribute{
									Description: "List of IPv4/IPv6 CIDR addresses to allow.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"not_in": schema.ListAttribute{
									Description: "List of IPv4/IPv6 CIDR addresses to deny.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"status": schema.StringAttribute{
							Description: "Status of the API Key.\nAvailable values: \"active\", \"inactive\", \"expired\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"active",
									"inactive",
									"expired",
								),
							},
						},
						"tags": schema.ListAttribute{
							Description: "Tags to attach to the API Key.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the API Key was updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"key": schema.StringAttribute{
							Description: "API Key. Only returned on creation.",
							Computed:    true,
							Sensitive:   true,
						},
						"starts_at": schema.StringAttribute{
							Description: "When the API Key starts to be valid.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *APIKeysDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *APIKeysDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
