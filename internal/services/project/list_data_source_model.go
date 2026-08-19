// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/projects"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ProjectsItemsDataSourceModel] `json:"items,computed"`
}

type ProjectsDataSourceModel struct {
	Name     types.String                                               `tfsdk:"name" query:"name,optional"`
	Tags     *[]types.String                                            `tfsdk:"tags" query:"tags,optional"`
	Sort     types.String                                               `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ProjectsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ProjectsDataSourceModel) toListParams(_ context.Context) (params projects.ProjectListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = projects.ProjectListParams{
		Tags: mTags,
	}

	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}

	return
}

type ProjectsItemsDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	Tags      customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
