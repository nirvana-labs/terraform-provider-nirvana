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
				Description: `Available values: "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-lon-1", "eu-ams-1", "eu-frk-1", "ap-mum-1", "ap-sin-1", "ap-tyo-1".`,
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
			"subnet_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"boot_volume": schema.SingleNestedAttribute{
				Description: "Boot volume create request.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Required: true,
						Validators: []validator.Int64{
							int64validator.Between(64, 512),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
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
			"data_volumes": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[ComputeVMDataVolumesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(32, 10240),
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU configuration details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"vcpu": schema.Int64Attribute{
						Description: "virtual CPUs",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 52),
						},
					},
				},
			},
			"memory_config": schema.SingleNestedAttribute{
				Description: "Memory configuration details.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "memory size",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 480),
						},
					},
				},
			},
			"boot_volume_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Description: `Available values: "vm", "volume", "vpc", "firewall_rule".`,
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
			"private_ip": schema.StringAttribute{
				Computed: true,
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
			},
			"resource_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "pending", "creating", "updating", "ready", "deleting", "deleted", "failed".`,
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
				Description: `Available values: "create", "update", "delete".`,
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
			"vpc_id": schema.StringAttribute{
				Computed: true,
			},
			"data_volume_ids": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
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
