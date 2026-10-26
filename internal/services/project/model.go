// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectModel struct {
	ID        types.String                                    `tfsdk:"id" json:"id,computed"`
	Name      types.String                                    `tfsdk:"name" json:"name,required"`
	Tags      *[]types.String                                 `tfsdk:"tags" json:"tags,optional"`
	CreatedAt timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ManagedBy customfield.NestedObject[ProjectManagedByModel] `tfsdk:"managed_by" json:"managed_by,computed"`
}

func (m ProjectModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProjectModel) MarshalJSONForUpdate(state ProjectModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ProjectManagedByModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
