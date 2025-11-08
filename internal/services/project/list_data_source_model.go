// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/projects"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ProjectsItemsDataSourceModel] `json:"items,computed"`
}

type ProjectsDataSourceModel struct {
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ProjectsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ProjectsDataSourceModel) toListParams(_ context.Context) (params projects.ProjectListParams, diags diag.Diagnostics) {
	params = projects.ProjectListParams{}

	return
}

type ProjectsItemsDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	Tags      customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UserID    types.String                   `tfsdk:"user_id" json:"user_id,computed"`
}
