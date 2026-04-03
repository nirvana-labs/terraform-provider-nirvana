// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/instance_types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type InstanceTypesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[InstanceTypesItemsDataSourceModel] `json:"items,computed"`
}

type InstanceTypesDataSourceModel struct {
	MaxItems types.Int64                                                     `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[InstanceTypesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *InstanceTypesDataSourceModel) toListParams(_ context.Context) (params instance_types.InstanceTypeListParams, diags diag.Diagnostics) {
	params = instance_types.InstanceTypeListParams{}

	return
}

type InstanceTypesItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"name,computed"`
	Chipset   types.String      `tfsdk:"chipset" json:"chipset,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Family    types.String      `tfsdk:"family" json:"family,computed"`
	MemoryGB  types.Int64       `tfsdk:"memory_gb" json:"memory_gb,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Region    types.String      `tfsdk:"region" json:"region,computed"`
	Series    types.String      `tfsdk:"series" json:"series,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Vcpu      types.Int64       `tfsdk:"vcpu" json:"vcpu,computed"`
}
