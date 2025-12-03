// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

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

var _ resource.ResourceWithConfigValidators = (*ComputeVolumeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Description: "Type of the Volume.\nAvailable values: \"nvme\", \"abs\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("nvme", "abs"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"vm_id": schema.StringAttribute{
				Description:   "ID of the VM the Volume is attached to.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Volume.",
				Required:    true,
			},
			"size": schema.Int64Attribute{
				Description: "Size of the Volume in GB.",
				Required:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Volume.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Volume was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"kind": schema.StringAttribute{
				Description: "Volume kind.\nAvailable values: \"boot\", \"data\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("boot", "data"),
				},
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
				Description: "When the Volume was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"vm_name": schema.StringAttribute{
				Description: "Name of the VM the Volume is attached to.",
				Computed:    true,
			},
		},
	}
}

func (r *ComputeVolumeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeVolumeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
