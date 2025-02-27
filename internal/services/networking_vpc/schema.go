// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*NetworkingVPCResource)(nil)

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
				Description: "available values: \"us-sea-1\", \"us-sva-1\", \"us-chi-1\", \"us-wdc-1\", \"eu-lon-1\", \"eu-ams-1\", \"eu-frk-1\", \"ap-mum-1\", \"ap-sin-1\", \"ap-tyo-1\"",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sea-1",
						"us-sva-1",
						"us-chi-1",
						"us-wdc-1",
						"eu-lon-1",
						"eu-ams-1",
						"eu-frk-1",
						"ap-mum-1",
						"ap-sin-1",
						"ap-tyo-1",
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
				Description: "available values: \"vm\", \"volume\", \"vpc\", \"firewall_rule\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"vm",
						"volume",
						"vpc",
						"firewall_rule",
					),
				},
			},
			"resource_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: "available values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"failed\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"creating",
						"updating",
						"ready",
						"deleting",
						"deleted",
						"failed",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "available values: \"create\", \"update\", \"delete\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"create",
						"update",
						"delete",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"firewall_rule_ids": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"subnet": schema.SingleNestedAttribute{
				Description: "Subnet details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NetworkingVPCSubnetModel](ctx),
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

func (r *NetworkingVPCResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkingVPCResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
