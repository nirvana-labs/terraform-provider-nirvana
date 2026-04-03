// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type InstanceTypeDataSourceModel struct {
	Name      types.String      `tfsdk:"name" path:"name,required"`
	Region    types.String      `tfsdk:"region" path:"region,required"`
	Chipset   types.String      `tfsdk:"chipset" json:"chipset,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Family    types.String      `tfsdk:"family" json:"family,computed"`
	MemoryGB  types.Int64       `tfsdk:"memory_gb" json:"memory_gb,computed"`
	Series    types.String      `tfsdk:"series" json:"series,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Vcpu      types.Int64       `tfsdk:"vcpu" json:"vcpu,computed"`
}
