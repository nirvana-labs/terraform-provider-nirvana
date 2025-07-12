// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeVolumeDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"volume_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the volume was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for the volume.",
				Computed:    true,
			},
			"kind": schema.StringAttribute{
				Description: "Volume kind.\nAvailable values: \"boot\", \"data\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("boot", "data"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the volume.",
				Computed:    true,
			},
			"size": schema.Int64Attribute{
				Description: "Size of the volume in GB.",
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
			"type": schema.StringAttribute{
				Description: "Storage type the volume is using.\nAvailable values: \"nvme\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("nvme"),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the volume was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"vm_id": schema.StringAttribute{
				Description: "ID of the VM the volume is attached to.",
				Computed:    true,
			},
			"vm_name": schema.StringAttribute{
				Description: "Name of the VM the volume is attached to.",
				Computed:    true,
			},
		},
	}
}

func (d *ComputeVolumeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeVolumeDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
