// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*VPCDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vpc_id": schema.StringAttribute{
				Required: true,
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
			"region": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"amsterdam",
						"chicago",
						"frankfurt",
						"hongkong",
						"london",
						"mumbai",
						"saopaulo",
						"seattle",
						"siliconvalley",
						"singapore",
						"stockholm",
						"sydney",
						"tokyo",
						"washingtondc",
						"staging",
					),
				},
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
			"firewall_rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[VPCFirewallRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"dest": schema.SingleNestedAttribute{
							Description: "Firewall rule endpoint.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[VPCFirewallRulesDestDataSourceModel](ctx),
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
						"name": schema.StringAttribute{
							Computed: true,
						},
						"protocol": schema.StringAttribute{
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Description: "Firewall rule endpoint.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[VPCFirewallRulesSourceDataSourceModel](ctx),
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
						"vpc_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"subnet": schema.SingleNestedAttribute{
				Description: "Subnet details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VPCSubnetDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"cidr": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *VPCDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *VPCDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
