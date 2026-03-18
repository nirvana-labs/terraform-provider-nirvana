// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*APIKeyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"api_key_id": schema.StringAttribute{
				Required: true,
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
			"key": schema.StringAttribute{
				Description: "API Key. Only returned on creation.",
				Computed:    true,
				Sensitive:   true,
			},
			"name": schema.StringAttribute{
				Description: "API Key name.",
				Computed:    true,
			},
			"starts_at": schema.StringAttribute{
				Description: "When the API Key starts to be valid.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"updated_at": schema.StringAttribute{
				Description: "When the API Key was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the API Key.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"source_ip_rule": schema.SingleNestedAttribute{
				Description: "IP filter rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[APIKeySourceIPRuleDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"allowed": schema.ListAttribute{
						Description: "List of IPv4 CIDR addresses to allow.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"blocked": schema.ListAttribute{
						Description: "List of IPv4 CIDR addresses to deny.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *APIKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *APIKeyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
