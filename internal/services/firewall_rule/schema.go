// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*FirewallRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vpc_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"protocol": schema.StringAttribute{
				Description: "Supported protocols.",
				Required:    true,
			},
			"destination": schema.SingleNestedAttribute{
				Description: "Firewall rule endpoint.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Optional: true,
					},
					"ports": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Description: "Firewall rule endpoint.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Optional: true,
					},
					"ports": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
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
						"VOLUME",
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
		},
	}
}

func (r *FirewallRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *FirewallRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
