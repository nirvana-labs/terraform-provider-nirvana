// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*FirewallRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"firewall_rule_id": schema.StringAttribute{
				Required: true,
			},
			"vpc_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"CREATING",
						"UPDATING",
						"READY",
						"DELETING",
						"DELETED",
						"FAILED",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"dest": schema.SingleNestedAttribute{
				Description: "Firewall rule endpoint.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[FirewallRuleDestDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Computed: true,
					},
					"ports": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Description: "Firewall rule endpoint.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[FirewallRuleSourceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Computed: true,
					},
					"ports": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *FirewallRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FirewallRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
