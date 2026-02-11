// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ProjectDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" path:"project_id,computed"`
	ProjectID types.String                   `tfsdk:"project_id" path:"project_id,required"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags      customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
}
