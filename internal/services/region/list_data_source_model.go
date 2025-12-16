// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package region

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/regions"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type RegionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[RegionsItemsDataSourceModel] `json:"items,computed"`
}

type RegionsDataSourceModel struct {
	MaxItems types.Int64                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[RegionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *RegionsDataSourceModel) toListParams(_ context.Context) (params regions.RegionListParams, diags diag.Diagnostics) {
	params = regions.RegionListParams{}

	return
}

type RegionsItemsDataSourceModel struct {
	Availability types.String                                               `tfsdk:"availability" json:"availability,computed"`
	Compute      customfield.NestedObject[RegionsComputeDataSourceModel]    `tfsdk:"compute" json:"compute,computed"`
	Name         types.String                                               `tfsdk:"name" json:"name,computed"`
	Networking   customfield.NestedObject[RegionsNetworkingDataSourceModel] `tfsdk:"networking" json:"networking,computed"`
	Storage      customfield.NestedObject[RegionsStorageDataSourceModel]    `tfsdk:"storage" json:"storage,computed"`
}

type RegionsComputeDataSourceModel struct {
	VMs types.Bool `tfsdk:"vms" json:"vms,computed"`
}

type RegionsNetworkingDataSourceModel struct {
	Connect types.Bool `tfsdk:"connect" json:"connect,computed"`
	VPCs    types.Bool `tfsdk:"vpcs" json:"vpcs,computed"`
}

type RegionsStorageDataSourceModel struct {
	ABS       types.Bool `tfsdk:"abs" json:"abs,computed"`
	LocalNvme types.Bool `tfsdk:"local_nvme" json:"local_nvme,computed"`
}
