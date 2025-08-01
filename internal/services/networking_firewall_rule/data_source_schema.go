// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule

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

var _ datasource.DataSourceWithConfigValidators = (*NetworkingFirewallRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"firewall_rule_id": schema.StringAttribute{
				Required: true,
			},
			"vpc_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the firewall rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"destination_address": schema.StringAttribute{
				Description: "Destination address of the firewall rule. Either VPC CIDR or VM in VPC.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for the firewall rule.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the firewall rule.",
				Computed:    true,
			},
			"protocol": schema.StringAttribute{
				Description: "Protocol of the firewall rule.\nAvailable values: \"tcp\", \"udp\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
				},
			},
			"source_address": schema.StringAttribute{
				Description: "Source address of the firewall rule. Address of 0.0.0.0 requires a CIDR mask of 0.",
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
				Description: "When the firewall rule was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"destination_ports": schema.ListAttribute{
				Description: "Destination ports of the firewall rule.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *NetworkingFirewallRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkingFirewallRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
