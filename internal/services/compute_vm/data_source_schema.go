// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeVMDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vm_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
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
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"vpc_id": schema.StringAttribute{
				Computed: true,
			},
			"boot_volume": schema.SingleNestedAttribute{
				Description: "Volume details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMBootVolumeDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"kind": schema.StringAttribute{
						Description: "Volume kind.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("boot", "data"),
						},
					},
					"size": schema.Int64Attribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Description: "Storage type.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("nvme"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"cpu_config": schema.SingleNestedAttribute{
				Description: "CPU details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMCPUConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"data_volumes": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ComputeVMDataVolumesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"kind": schema.StringAttribute{
							Description: "Volume kind.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("boot", "data"),
							},
						},
						"size": schema.Int64Attribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: "Storage type.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("nvme"),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"mem_config": schema.SingleNestedAttribute{
				Description: "RAM details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeVMMemConfigDataSourceModel](ctx),
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

func (d *ComputeVMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeVMDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
