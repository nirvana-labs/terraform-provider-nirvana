// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeVMDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vm_id": schema.StringAttribute{
				Required: true,
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
			"name": schema.StringAttribute{
				Description: "Name of the VM.",
				Computed:    true,
			},
			"private_ip": schema.StringAttribute{
				Description: "Private IP of the VM.",
				Computed:    true,
			},
			"public_ip": schema.StringAttribute{
				Description: "Public IP of the VM.",
				Computed:    true,
			},
			"public_ip_enabled": schema.BoolAttribute{
				Description: "Whether the public IP is enabled for the VM.",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sea-1\", \"us-sva-1\", \"us-sva-2\", \"us-chi-1\", \"us-wdc-1\", \"eu-frk-1\", \"ap-sin-1\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sea-1",
						"us-sva-1",
						"us-sva-2",
						"us-chi-1",
						"us-wdc-1",
						"eu-frk-1",
						"ap-sin-1",
					),
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
			"subnet_id": schema.StringAttribute{
				Description: "ID of the subnet the VM is in.",
				Computed:    true,
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
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the VM.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU configuration for the VM.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMCPUConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"vcpu": schema.Int64Attribute{
						Description: "Number of virtual CPUs.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 52),
						},
					},
				},
			},
			"memory_config": schema.SingleNestedAttribute{
				Description: "Memory configuration for the VM.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMMemoryConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"size": schema.Int64Attribute{
						Description: "Size of the memory in GB.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 480),
						},
					},
				},
			},
		},
	}
}

func (d *ComputeVMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeVMDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
