// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_kubeconfig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NKSClusterKubeconfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				Required: true,
			},
			"kubeconfig": schema.StringAttribute{
				Description: "Kubeconfig content for the Cluster.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (d *NKSClusterKubeconfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NKSClusterKubeconfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
