// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"os_image_name": schema.StringAttribute{
				Description:   "Name of the OS Image to use for the VM.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"subnet_id": schema.StringAttribute{
				Description:   "ID of the subnet to use for the VM.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"boot_volume": schema.SingleNestedAttribute{
				Description: "Boot volume for the VM.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "Size of the Volume in GB.",
						Required:    true,
					},
					"tags": schema.ListAttribute{
						Description: "Tags to attach to the Volume.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"type": schema.StringAttribute{
						Description: "Type of the Volume.\nAvailable values: \"nvme\", \"abs\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("nvme", "abs"),
						},
						Default: stringdefault.StaticString("nvme"),
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"ssh_key": schema.SingleNestedAttribute{
				Description: "Public SSH key configuration for the VM.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"public_key": schema.StringAttribute{
						Description: "Public key to and and use to access the VM.",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"data_volumes": schema.ListNestedAttribute{
				Description: "Data volumes for the VM.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ComputeVMDataVolumesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
						"type": schema.StringAttribute{
							Description: "Type of the Volume.\nAvailable values: \"nvme\", \"abs\".",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("nvme", "abs"),
							},
							Default: stringdefault.StaticString("nvme"),
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplaceIfConfigured()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the VM.",
				Optional:    true,
			},
			"public_ip_enabled": schema.BoolAttribute{
				Description: "Whether to enable public IP for the VM.",
				Optional:    true,
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the VM.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU configuration for the VM.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"vcpu": schema.Int64Attribute{
						Description: "Number of virtual CPUs.",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 52),
						},
					},
				},
			},
			"memory_config": schema.SingleNestedAttribute{
				Description: "Memory configuration for the VM.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "Size of the memory in GB.",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 480),
						},
					},
				},
			},
			"boot_volume_id": schema.StringAttribute{
				Description: "ID of the boot volume attached to the VM.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the VM was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"private_ip": schema.StringAttribute{
				Description: "Private IP of the VM.",
				Computed:    true,
			},
			"public_ip": schema.StringAttribute{
				Description: "Public IP of the VM.",
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
				Description: "When the VM was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"vpc_id": schema.StringAttribute{
				Description: "ID of the VPC the VM is in.",
				Computed:    true,
			},
			"vpc_name": schema.StringAttribute{
				Description: "Name of the VPC the VM is in.",
				Computed:    true,
			},
			"data_volume_ids": schema.ListAttribute{
				Description: "IDs of the data volumes attached to the VM.",
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
