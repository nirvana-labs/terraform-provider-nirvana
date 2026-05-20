// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_node_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*NKSNodePoolResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"cluster_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the node pool.",
				Required:    true,
			},
			"node_count": schema.Int64Attribute{
				Description: "Number of nodes. Must be between 1 and 100.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"node_config": schema.SingleNestedAttribute{
				Description: "Node configuration.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"boot_volume": schema.SingleNestedAttribute{
						Description: "Boot volume configuration.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"size": schema.Int64Attribute{
								Description: "Size of the boot volume in GB.",
								Required:    true,
								Validators: []validator.Int64{
									int64validator.Between(64, 512),
								},
							},
							"type": schema.StringAttribute{
								Description: "Type of the Volume.\nAvailable values: \"nvme\", \"abs\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("nvme", "abs"),
								},
							},
						},
					},
					"instance_type": schema.StringAttribute{
						Description: "Instance type name used for worker nodes.",
						Required:    true,
					},
					"labels": schema.ListAttribute{
						Description: "Kubernetes labels to apply to each node in the pool. Each entry is \"key=value\".\nKeys under kubernetes.io, k8s.io, and nirvanalabs.io prefixes are reserved.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the node pool.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the node pool was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"project_id": schema.StringAttribute{
				Description: "Project ID the Operation belongs to.",
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
				Description: "When the node pool was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *NKSNodePoolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NKSNodePoolResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
