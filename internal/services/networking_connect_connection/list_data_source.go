// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NetworkingConnectConnectionsDataSource struct {
	client *nirvana.Client
}

var _ datasource.DataSourceWithConfigure = (*NetworkingConnectConnectionsDataSource)(nil)

func NewNetworkingConnectConnectionsDataSource() datasource.DataSource {
	return &NetworkingConnectConnectionsDataSource{}
}

func (d *NetworkingConnectConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_connect_connections"
}

func (d *NetworkingConnectConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*nirvana.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *nirvana.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *NetworkingConnectConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NetworkingConnectConnectionsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := NetworkingConnectConnectionsItemsListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	page, err := d.client.Networking.Connect.Connections.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Items) > 0 {
		bytes := []byte(page.RawJSON())
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, env.Items.Elements()...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:min(len(acc), maxItems)]
	result, diags := customfield.NewObjectListFromAttributes[NetworkingConnectConnectionsItemsDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Items = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
