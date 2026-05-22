// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_kubernetes_version

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSClusterKubernetesVersionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NKSClusterKubernetesVersionsItemsDataSourceModel] `json:"items,computed"`
}

type NKSClusterKubernetesVersionsDataSourceModel struct {
	MaxItems types.Int64                                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[NKSClusterKubernetesVersionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NKSClusterKubernetesVersionsDataSourceModel) toListParams(_ context.Context) (params nks.ClusterKubernetesVersionListParams, diags diag.Diagnostics) {
	params = nks.ClusterKubernetesVersionListParams{}

	return
}

type NKSClusterKubernetesVersionsItemsDataSourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DisplayName types.String      `tfsdk:"display_name" json:"display_name,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
}
