// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package region

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type RegionDataSourceModel struct {
	Name         types.String                                              `tfsdk:"name" path:"name,required"`
	Availability types.String                                              `tfsdk:"availability" json:"availability,computed"`
	Compute      customfield.NestedObject[RegionComputeDataSourceModel]    `tfsdk:"compute" json:"compute,computed"`
	Networking   customfield.NestedObject[RegionNetworkingDataSourceModel] `tfsdk:"networking" json:"networking,computed"`
	Storage      customfield.NestedObject[RegionStorageDataSourceModel]    `tfsdk:"storage" json:"storage,computed"`
}

type RegionComputeDataSourceModel struct {
	VMs types.Bool `tfsdk:"vms" json:"vms,computed"`
}

type RegionNetworkingDataSourceModel struct {
	Connect types.Bool `tfsdk:"connect" json:"connect,computed"`
	VPCs    types.Bool `tfsdk:"vpcs" json:"vpcs,computed"`
}

type RegionStorageDataSourceModel struct {
	ABS       types.Bool `tfsdk:"abs" json:"abs,computed"`
	LocalNvme types.Bool `tfsdk:"local_nvme" json:"local_nvme,computed"`
}
