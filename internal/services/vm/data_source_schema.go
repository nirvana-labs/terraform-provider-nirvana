// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*VMDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vm_id": schema.StringAttribute{
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
			"public_ip": schema.StringAttribute{
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
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VMCPUConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"mem_config": schema.SingleNestedAttribute{
				Description: "RAM details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VMMemConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "RAM size",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"unit": schema.StringAttribute{
						Description: "Unit (GB, MB, etc.)",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("GB"),
						},
					},
				},
			},
			"storage_config": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[VMStorageConfigDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Description: "Storage size",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
						},
						"type": schema.StringAttribute{
							Description: "Storage type.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("nvme"),
							},
						},
						"unit": schema.StringAttribute{
							Description: "Storage unit.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("GB"),
							},
						},
						"disk_name": schema.StringAttribute{
							Description: "Disk name, used later",
							Computed:    true,
						},
					},
				},
			},
			"vpc": schema.SingleNestedAttribute{
				Description: "VPC details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VMVPCDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"firewall_rules": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[VMVPCFirewallRulesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"created_at": schema.StringAttribute{
									Computed: true,
								},
								"destination": schema.SingleNestedAttribute{
									Description: "Firewall rule endpoint.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[VMVPCFirewallRulesDestinationDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[VMVPCFirewallRulesSourceDataSourceModel](ctx),
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
					"subnet": schema.SingleNestedAttribute{
						Description: "Subnet details.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[VMVPCSubnetDataSourceModel](ctx),
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
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *VMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *VMDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
