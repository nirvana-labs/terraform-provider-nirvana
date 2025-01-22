// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*VMResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"need_public_ip": schema.BoolAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"os_image_id": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Required: true,
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source_address": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ports": schema.ListAttribute{
				Required:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"ssh_key": schema.SingleNestedAttribute{
				Description: "SSH key details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"public_key": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"subnet_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"boot_volume": schema.SingleNestedAttribute{
				Description: "Boot volume create request.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.Between(50, 500),
						},
					},
				},
			},
			"cpu": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"ram": schema.SingleNestedAttribute{
				Description: "RAM details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "RAM size",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 128),
						},
					},
				},
			},
			"data_volumes": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[VMDataVolumesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(50, 1400),
							},
						},
						"type": schema.StringAttribute{
							Description: "Storage type.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("nvme"),
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"VM",
						"VPC",
						"FIREWALL_RULE",
					),
				},
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
			},
			"resource_id": schema.StringAttribute{
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
			"type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"CREATE",
						"UPDATE",
						"DELETE",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VMCPUConfigModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[VMMemConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "RAM size",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 128),
						},
					},
				},
			},
			"vpc": schema.SingleNestedAttribute{
				Description: "VPC details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VMVPCModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"firewall_rules": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[VMVPCFirewallRulesModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[VMVPCFirewallRulesDestinationModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[VMVPCFirewallRulesSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[VMVPCSubnetModel](ctx),
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

func (r *VMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *VMResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
