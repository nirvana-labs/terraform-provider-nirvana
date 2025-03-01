// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkingVPCDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vpc_id": schema.StringAttribute{
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
			"region": schema.StringAttribute{
				Description: `Available values: "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-lon-1", "eu-ams-1", "eu-frk-1", "ap-mum-1", "ap-sin-1", "ap-tyo-1".`,
				Computed:    true,
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
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"firewall_rule_ids": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"subnet": schema.SingleNestedAttribute{
				Description: "Subnet details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NetworkingVPCSubnetDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"cidr": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworkingVPCDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkingVPCDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
