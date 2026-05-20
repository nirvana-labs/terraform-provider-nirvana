// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

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

var _ resource.ResourceWithConfigValidators = (*NKSClusterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"project_id": schema.StringAttribute{
				Description:   "Project ID to create the Cluster in.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sva-2\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("us-sva-2"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"vpc_id": schema.StringAttribute{
				Description:   "ID of the VPC to use for the Cluster.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"autoscaling": schema.BoolAttribute{
				Description: "Whether to enable autoscaling for the Cluster.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Cluster.",
				Required:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Cluster.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Cluster was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"private_ip": schema.StringAttribute{
				Description: "Private IP (VIP) of the Cluster.",
				Computed:    true,
			},
			"public_ip": schema.StringAttribute{
				Description: "Public IP of the Cluster.",
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
				Description: "When the Cluster was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"pool_ids": schema.ListAttribute{
				Description: "IDs of pools belonging to this Cluster.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"details": schema.SingleNestedAttribute{
				Description: "Structured details about what an operation is changing.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NKSClusterDetailsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"changes": schema.MapNestedAttribute{
						Description: "Map of changed field names to their from/to diffs. Keys depend on the parent operation's kind+type.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectMapType[NKSClusterDetailsChangesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description: "Previous value.",
									Computed:    true,
								},
								"to": schema.StringAttribute{
									Description: "New value.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NKSClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NKSClusterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
