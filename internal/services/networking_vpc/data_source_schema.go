// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
				Description: "When the VPC was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for the VPC.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the VPC.",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sea-1\", \"us-sva-1\", \"us-chi-1\", \"us-wdc-1\", \"eu-frk-1\", \"ap-sin-1\", \"ap-seo-1\", \"ap-tyo-1\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"us-sea-1",
						"us-sva-1",
						"us-chi-1",
						"us-wdc-1",
						"eu-frk-1",
						"ap-sin-1",
						"ap-seo-1",
						"ap-tyo-1",
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
			"updated_at": schema.StringAttribute{
				Description: "When the VPC was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"firewall_rule_ids": schema.ListAttribute{
				Description: "IDs of the firewall rules associated with the VPC.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"subnet": schema.SingleNestedAttribute{
				Description: "Subnet of the VPC.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NetworkingVPCSubnetDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique identifier for the subnet.",
						Computed:    true,
					},
					"cidr": schema.StringAttribute{
						Description: "CIDR block for the subnet.",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "When the subnet was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"name": schema.StringAttribute{
						Description: "Name of the subnet.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Description: "When the subnet was updated.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
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
