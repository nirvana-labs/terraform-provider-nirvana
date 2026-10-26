// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

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

var _ resource.ResourceWithConfigValidators = (*ProjectResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Project ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description: "Project name.",
				Required:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Project.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Project was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the Project was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"managed_by": schema.SingleNestedAttribute{
				Description: "What manages the resource, or null when nothing does. A managed resource can be read but not changed.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ProjectManagedByModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "ID of the thing managing the resource.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "Kind of thing managing the resource.\nAvailable values: \"sandbox_environment\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("sandbox_environment"),
						},
					},
				},
			},
		},
	}
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ProjectResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
