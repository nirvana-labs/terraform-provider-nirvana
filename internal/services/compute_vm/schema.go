// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ComputeVMResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"os_image_name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"public_ip_enabled": schema.BoolAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
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
			"source_address": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ports": schema.ListAttribute{
				Required:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"ssh_key": schema.SingleNestedAttribute{
				Description: "SSH key details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"public_key": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"subnet_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"boot_volume": schema.SingleNestedAttribute{
				Description: "Boot volume create request.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.Between(50, 500),
						},
					},
				},
			},
			"cpu": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"ram": schema.SingleNestedAttribute{
				Description: "RAM details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "RAM size",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 128),
						},
					},
				},
			},
			"data_volumes": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[ComputeVMDataVolumesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(50, 1400),
							},
						},
						"type": schema.StringAttribute{
							Description: "Storage type.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("nvme"),
							},
						},
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
						"vm",
						"volume",
						"vpc",
						"firewall_rule",
					),
				},
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
			},
			"resource_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
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
				Computed: true,
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
			"vpc_id": schema.StringAttribute{
				Computed: true,
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMCPUConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"mem_config": schema.SingleNestedAttribute{
				Description: "RAM details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMMemConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "RAM size",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 128),
						},
					},
				},
			},
		},
	}
}

func (r *ComputeVMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeVMResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
