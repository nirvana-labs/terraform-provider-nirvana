// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

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

var _ datasource.DataSourceWithConfigValidators = (*NetworkingFirewallRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vpc_id": schema.StringAttribute{
				Required: true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkingFirewallRulesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the Firewall Rule.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "When the Firewall Rule was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"destination_address": schema.StringAttribute{
							Description: "Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC.",
							Computed:    true,
						},
						"destination_ports": schema.ListAttribute{
							Description: "Destination ports of the Firewall Rule.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "Name of the Firewall Rule.",
							Computed:    true,
						},
						"protocol": schema.StringAttribute{
							Description: "Protocol of the Firewall Rule.\nAvailable values: \"tcp\", \"udp\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
							},
						},
						"source_address": schema.StringAttribute{
							Description: "Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask of 0.",
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
						"tags": schema.ListAttribute{
							Description: "Tags to attach to the Firewall Rule.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the Firewall Rule was updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"vpc_id": schema.StringAttribute{
							Description: "ID of the VPC the Firewall Rule belongs to.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkingFirewallRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *NetworkingFirewallRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
