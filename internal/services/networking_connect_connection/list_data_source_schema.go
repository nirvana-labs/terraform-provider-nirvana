// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

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

var _ datasource.DataSourceWithConfigValidators = (*NetworkingConnectConnectionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
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
				CustomType:  customfield.NewNestedObjectListType[NetworkingConnectConnectionsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the Connect Connection",
							Computed:    true,
						},
						"asn": schema.Int64Attribute{
							Description: "ASN",
							Computed:    true,
						},
						"aws": schema.SingleNestedAttribute{
							Description: "AWS provider configuration",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[NetworkingConnectConnectionsAWSDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"region": schema.StringAttribute{
									Description: "AWS region where the connection is established",
									Computed:    true,
								},
							},
						},
						"bandwidth_mbps": schema.Int64Attribute{
							Description: "Connect Connection speed in Mbps\nAvailable values: 50, 200, 500, 1000, 2000.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.OneOf(
									50,
									200,
									500,
									1000,
									2000,
								),
							},
						},
						"cidrs": schema.ListAttribute{
							Description: "CIDRs for the Connect Connection",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"created_at": schema.StringAttribute{
							Description: "When the Connect Connection was created",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "Name of the Connect Connection",
							Computed:    true,
						},
						"provider_asn": schema.Int64Attribute{
							Description: "Provider ASN",
							Computed:    true,
						},
						"provider_cidrs": schema.ListAttribute{
							Description: "Provider CIDRs for the Connect Connection",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"provider_router_ip": schema.StringAttribute{
							Description: "Provider Router IP for the Connect Connection",
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
						"router_ip": schema.StringAttribute{
							Description: "Router IP",
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
							Description: "Tags to attach to the Connect Connection",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "When the Connect Connection was updated",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *NetworkingConnectConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *NetworkingConnectConnectionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
