// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_kubeconfig

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/logging"
)

type NKSClusterKubeconfigDataSource struct {
	client *nirvana.Client
}

var _ datasource.DataSourceWithConfigure = (*NKSClusterKubeconfigDataSource)(nil)

func NewNKSClusterKubeconfigDataSource() datasource.DataSource {
	return &NKSClusterKubeconfigDataSource{}
}

func (d *NKSClusterKubeconfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nks_cluster_kubeconfig"
}

func (d *NKSClusterKubeconfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NKSClusterKubeconfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NKSClusterKubeconfigDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.NKS.Clusters.Kubeconfig.Get(
		ctx,
		data.ClusterID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
