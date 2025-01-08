// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*VPCResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"subnet_name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"firewall_rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[VPCFirewallRulesModel](ctx),
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
							CustomType:  customfield.NewNestedObjectType[VPCFirewallRulesDestModel](ctx),
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
							CustomType:  customfield.NewNestedObjectType[VPCFirewallRulesSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[VPCSubnetModel](ctx),
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

func (r *VPCResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *VPCResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
