// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_kubeconfig

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NKSClusterKubeconfigDataSourceModel struct {
	ClusterID  types.String `tfsdk:"cluster_id" path:"cluster_id,required"`
	Kubeconfig types.String `tfsdk:"kubeconfig" json:"kubeconfig,computed"`
}
