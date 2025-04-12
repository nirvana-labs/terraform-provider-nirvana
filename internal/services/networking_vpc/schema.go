// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

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
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*NetworkingVPCResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"region": schema.StringAttribute{
				Description: "Region of the VPC.\nAvailable values: \"us-sea-1\", \"us-sva-1\", \"us-chi-1\", \"us-wdc-1\", \"eu-lon-1\", \"eu-ams-1\", \"eu-frk-1\", \"ap-sin-1\", \"ap-seo-1\", \"ap-tyo-1\".",
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
						"ap-sin-1",
						"ap-seo-1",
						"ap-tyo-1",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the VPC.",
				Required:    true,
			},
			"subnet_name": schema.StringAttribute{
				Description: "Name of the subnet to create.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Time the VPC was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Status of the VPC.\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"error\".",
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
				Description: "Time the VPC was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"firewall_rule_ids": schema.ListAttribute{
				Description: "IDs of the firewall rules associated with the VPC.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"subnet": schema.SingleNestedAttribute{
				Description: "Subnet of the VPC.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NetworkingVPCSubnetModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique identifier for the subnet.",
						Computed:    true,
					},
					"cidr": schema.StringAttribute{
						Description: "CIDR block for the subnet.",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "Time the subnet was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"name": schema.StringAttribute{
						Description: "Name of the subnet.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Description: "Time the subnet was updated.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
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
