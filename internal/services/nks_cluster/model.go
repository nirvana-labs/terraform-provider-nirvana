// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSClusterModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	ProjectID types.String                   `tfsdk:"project_id" json:"project_id,required"`
	Region    types.String                   `tfsdk:"region" json:"region,required"`
	VPCID     types.String                   `tfsdk:"vpc_id" json:"vpc_id,required"`
	Name      types.String                   `tfsdk:"name" json:"name,required"`
	Tags      *[]types.String                `tfsdk:"tags" json:"tags,optional"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	PrivateIP types.String                   `tfsdk:"private_ip" json:"private_ip,computed"`
	PublicIP  types.String                   `tfsdk:"public_ip" json:"public_ip,computed"`
	Status    types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	PoolIDs   customfield.List[types.String] `tfsdk:"pool_ids" json:"pool_ids,computed"`
}

func (m NKSClusterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NKSClusterModel) MarshalJSONForUpdate(state NKSClusterModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
