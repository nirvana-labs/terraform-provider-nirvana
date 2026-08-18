// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/instance_types"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type InstanceTypesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[InstanceTypesItemsDataSourceModel] `json:"items,computed"`
}

type InstanceTypesDataSourceModel struct {
	Chipset                 types.String                                                    `tfsdk:"chipset" query:"chipset,optional"`
	Family                  types.String                                                    `tfsdk:"family" query:"family,optional"`
	MemoryGBMax             types.Int64                                                     `tfsdk:"memory_gb_max" query:"memory_gb_max,optional"`
	MemoryGBMin             types.Int64                                                     `tfsdk:"memory_gb_min" query:"memory_gb_min,optional"`
	Name                    types.String                                                    `tfsdk:"name" query:"name,optional"`
	NetworkBandwidthGbpsMax types.Float64                                                   `tfsdk:"network_bandwidth_gbps_max" query:"network_bandwidth_gbps_max,optional"`
	NetworkBandwidthGbpsMin types.Float64                                                   `tfsdk:"network_bandwidth_gbps_min" query:"network_bandwidth_gbps_min,optional"`
	Region                  types.String                                                    `tfsdk:"region" query:"region,optional"`
	Series                  types.String                                                    `tfsdk:"series" query:"series,optional"`
	VcpuMax                 types.Int64                                                     `tfsdk:"vcpu_max" query:"vcpu_max,optional"`
	VcpuMin                 types.Int64                                                     `tfsdk:"vcpu_min" query:"vcpu_min,optional"`
	Sort                    types.String                                                    `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems                types.Int64                                                     `tfsdk:"max_items"`
	Items                   customfield.NestedObjectList[InstanceTypesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *InstanceTypesDataSourceModel) toListParams(_ context.Context) (params instance_types.InstanceTypeListParams, diags diag.Diagnostics) {
	params = instance_types.InstanceTypeListParams{}

	if !m.Chipset.IsNull() {
		params.Chipset = param.NewOpt(m.Chipset.ValueString())
	}
	if !m.Family.IsNull() {
		params.Family = param.NewOpt(m.Family.ValueString())
	}
	if !m.MemoryGBMax.IsNull() {
		params.MemoryGBMax = param.NewOpt(m.MemoryGBMax.ValueInt64())
	}
	if !m.MemoryGBMin.IsNull() {
		params.MemoryGBMin = param.NewOpt(m.MemoryGBMin.ValueInt64())
	}
	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.NetworkBandwidthGbpsMax.IsNull() {
		params.NetworkBandwidthGbpsMax = param.NewOpt(m.NetworkBandwidthGbpsMax.ValueFloat64())
	}
	if !m.NetworkBandwidthGbpsMin.IsNull() {
		params.NetworkBandwidthGbpsMin = param.NewOpt(m.NetworkBandwidthGbpsMin.ValueFloat64())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Series.IsNull() {
		params.Series = param.NewOpt(m.Series.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.VcpuMax.IsNull() {
		params.VcpuMax = param.NewOpt(m.VcpuMax.ValueInt64())
	}
	if !m.VcpuMin.IsNull() {
		params.VcpuMin = param.NewOpt(m.VcpuMin.ValueInt64())
	}

	return
}

type InstanceTypesItemsDataSourceModel struct {
	ID                   types.String      `tfsdk:"id" json:"name,computed"`
	Chipset              types.String      `tfsdk:"chipset" json:"chipset,computed"`
	CreatedAt            timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Family               types.String      `tfsdk:"family" json:"family,computed"`
	MemoryGB             types.Int64       `tfsdk:"memory_gb" json:"memory_gb,computed"`
	Name                 types.String      `tfsdk:"name" json:"name,computed"`
	NetworkBandwidthGbps types.Float64     `tfsdk:"network_bandwidth_gbps" json:"network_bandwidth_gbps,computed"`
	Region               types.String      `tfsdk:"region" json:"region,computed"`
	Series               types.String      `tfsdk:"series" json:"series,computed"`
	UpdatedAt            timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Vcpu                 types.Int64       `tfsdk:"vcpu" json:"vcpu,computed"`
}
