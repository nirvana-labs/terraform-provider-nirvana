// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*NetworkingFirewallRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vpc_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"destination_address": schema.StringAttribute{
				Description: "Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC. Must be in network-aligned/canonical form.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Firewall Rule.",
				Optional:    true,
			},
			"protocol": schema.StringAttribute{
				Description: "Protocol of the Firewall Rule.\nAvailable values: \"tcp\", \"udp\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
				},
			},
			"source_address": schema.StringAttribute{
				Description: "Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask of 0. Must be in network-aligned/canonical form.",
				Optional:    true,
			},
			"destination_ports": schema.ListAttribute{
				Description: "Destination ports of the Firewall Rule.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Firewall Rule.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Firewall Rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Status of the resource.\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"creating",
						"updating",
						"ready",
						"deleting",
						"deleted",
						"error",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the Firewall Rule was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *NetworkingFirewallRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkingFirewallRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
