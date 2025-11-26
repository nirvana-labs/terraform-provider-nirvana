// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*NetworkingConnectConnectionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bandwidth_mbps": schema.Int64Attribute{
				Description: "Connect Connection speed in Mbps\nAvailable values: 50, 200, 500, 1000, 2000.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.OneOf(
						50,
						200,
						500,
						1000,
						2000,
					),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sea-1\", \"us-sva-1\", \"us-sva-2\", \"us-chi-1\", \"us-wdc-1\", \"eu-frk-1\", \"ap-sin-1\", \"ap-seo-1\", \"ap-tyo-1\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sea-1",
						"us-sva-1",
						"us-sva-2",
						"us-chi-1",
						"us-wdc-1",
						"eu-frk-1",
						"ap-sin-1",
						"ap-seo-1",
						"ap-tyo-1",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cidrs": schema.ListAttribute{
				Description:   "CIDRs for the Connect Connection. Must be in network-aligned/canonical form.",
				Required:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"provider_cidrs": schema.ListAttribute{
				Description:   "Provider CIDRs. Must be in network-aligned/canonical form.",
				Required:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"aws": schema.SingleNestedAttribute{
				Description: "AWS provider configuration",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "AWS account id",
						Required:    true,
					},
					"region": schema.StringAttribute{
						Description: "AWS region where the connection will be established",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Connect Connection.",
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Connect Connection",
				Optional:    true,
				ElementType: types.StringType,
			},
			"asn": schema.Int64Attribute{
				Description: "ASN",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Connect Connection was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"provider_asn": schema.Int64Attribute{
				Description: "Provider ASN",
				Computed:    true,
			},
			"provider_router_ip": schema.StringAttribute{
				Description: "Provider Router IP for the Connect Connection",
				Computed:    true,
			},
			"router_ip": schema.StringAttribute{
				Description: "Router IP",
				Computed:    true,
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
				Description: "When the Connect Connection was updated",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *NetworkingConnectConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkingConnectConnectionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
