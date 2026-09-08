// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package region

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/regions"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type RegionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[RegionsItemsDataSourceModel] `json:"items,computed"`
}

type RegionsDataSourceModel struct {
	Availability     types.String                                              `tfsdk:"availability" query:"availability,optional"`
	ComputeVMs       types.Bool                                                `tfsdk:"compute_vms" query:"compute_vms,optional"`
	NetworkingVPCs   types.Bool                                                `tfsdk:"networking_vpcs" query:"networking_vpcs,optional"`
	NKSAutoscaling   types.Bool                                                `tfsdk:"nks_autoscaling" query:"nks_autoscaling,optional"`
	NKSClusters      types.Bool                                                `tfsdk:"nks_clusters" query:"nks_clusters,optional"`
	StorageABS       types.Bool                                                `tfsdk:"storage_abs" query:"storage_abs,optional"`
	StorageLocalNvme types.Bool                                                `tfsdk:"storage_local_nvme" query:"storage_local_nvme,optional"`
	Sort             types.String                                              `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems         types.Int64                                               `tfsdk:"max_items"`
	Items            customfield.NestedObjectList[RegionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *RegionsDataSourceModel) toListParams(_ context.Context) (params regions.RegionListParams, diags diag.Diagnostics) {
	params = regions.RegionListParams{}

	if !m.Availability.IsNull() {
		params.Availability = regions.RegionListParamsAvailability(m.Availability.ValueString())
	}
	if !m.ComputeVMs.IsNull() {
		params.ComputeVMs = param.NewOpt(m.ComputeVMs.ValueBool())
	}
	if !m.NetworkingVPCs.IsNull() {
		params.NetworkingVPCs = param.NewOpt(m.NetworkingVPCs.ValueBool())
	}
	if !m.NKSAutoscaling.IsNull() {
		params.NKSAutoscaling = param.NewOpt(m.NKSAutoscaling.ValueBool())
	}
	if !m.NKSClusters.IsNull() {
		params.NKSClusters = param.NewOpt(m.NKSClusters.ValueBool())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.StorageABS.IsNull() {
		params.StorageABS = param.NewOpt(m.StorageABS.ValueBool())
	}
	if !m.StorageLocalNvme.IsNull() {
		params.StorageLocalNvme = param.NewOpt(m.StorageLocalNvme.ValueBool())
	}

	return
}

type RegionsItemsDataSourceModel struct {
	ID           types.String                                               `tfsdk:"id" json:"name,computed"`
	Availability types.String                                               `tfsdk:"availability" json:"availability,computed"`
	Compute      customfield.NestedObject[RegionsComputeDataSourceModel]    `tfsdk:"compute" json:"compute,computed"`
	Name         types.String                                               `tfsdk:"name" json:"name,computed"`
	Networking   customfield.NestedObject[RegionsNetworkingDataSourceModel] `tfsdk:"networking" json:"networking,computed"`
	NKS          customfield.NestedObject[RegionsNKSDataSourceModel]        `tfsdk:"nks" json:"nks,computed"`
	Storage      customfield.NestedObject[RegionsStorageDataSourceModel]    `tfsdk:"storage" json:"storage,computed"`
}

type RegionsComputeDataSourceModel struct {
	VMs types.Bool `tfsdk:"vms" json:"vms,computed"`
}

type RegionsNetworkingDataSourceModel struct {
	VPCs types.Bool `tfsdk:"vpcs" json:"vpcs,computed"`
}

type RegionsNKSDataSourceModel struct {
	Autoscaling types.Bool `tfsdk:"autoscaling" json:"autoscaling,computed"`
	Clusters    types.Bool `tfsdk:"clusters" json:"clusters,computed"`
}

type RegionsStorageDataSourceModel struct {
	ABS       types.Bool `tfsdk:"abs" json:"abs,computed"`
	LocalNvme types.Bool `tfsdk:"local_nvme" json:"local_nvme,computed"`
}
