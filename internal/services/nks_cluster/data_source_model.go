// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSClusterDataSourceModel struct {
	ID        types.String                        `tfsdk:"id" path:"cluster_id,computed"`
	ClusterID types.String                        `tfsdk:"cluster_id" path:"cluster_id,optional"`
	CreatedAt timetypes.RFC3339                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                        `tfsdk:"name" json:"name,computed"`
	PrivateIP types.String                        `tfsdk:"private_ip" json:"private_ip,computed"`
	ProjectID types.String                        `tfsdk:"project_id" json:"project_id,computed"`
	PublicIP  types.String                        `tfsdk:"public_ip" json:"public_ip,computed"`
	Region    types.String                        `tfsdk:"region" json:"region,computed"`
	Status    types.String                        `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID     types.String                        `tfsdk:"vpc_id" json:"vpc_id,computed"`
	PoolIDs   customfield.List[types.String]      `tfsdk:"pool_ids" json:"pool_ids,computed"`
	Tags      customfield.List[types.String]      `tfsdk:"tags" json:"tags,computed"`
	FindOneBy *NKSClusterFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *NKSClusterDataSourceModel) toListParams(_ context.Context) (params nks.ClusterListParams, diags diag.Diagnostics) {
	params = nks.ClusterListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
	}

	return
}

type NKSClusterFindOneByDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" query:"project_id,required"`
}
