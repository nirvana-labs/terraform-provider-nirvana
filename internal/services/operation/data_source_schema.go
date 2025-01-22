// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package operation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*OperationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"operation_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"VM",
						"VOLUME",
						"VPC",
						"FIREWALL_RULE",
					),
				},
			},
			"resource_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"RUNNING",
						"DONE",
						"FAILED",
						"UNKNOWN",
					),
				},
			},
			"type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"CREATE",
						"UPDATE",
						"DELETE",
					),
				},
			},
		},
	}
}

func (d *OperationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OperationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
