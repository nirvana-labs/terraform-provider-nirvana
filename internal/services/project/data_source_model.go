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

type ProjectDataSourceModel struct {
	ID        types.String                     `tfsdk:"id" path:"project_id,computed"`
	ProjectID types.String                     `tfsdk:"project_id" path:"project_id,optional"`
	CreatedAt timetypes.RFC3339                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                     `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339                `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags      customfield.List[types.String]   `tfsdk:"tags" json:"tags,computed"`
	FindOneBy *ProjectFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ProjectDataSourceModel) toListParams(_ context.Context) (params projects.ProjectListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = projects.ProjectListParams{
		Tags: mFindOneByTags,
	}

	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}

	return
}

type ProjectFindOneByDataSourceModel struct {
	Name types.String    `tfsdk:"name" query:"name,optional"`
	Sort types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Tags *[]types.String `tfsdk:"tags" query:"tags,optional"`
}
