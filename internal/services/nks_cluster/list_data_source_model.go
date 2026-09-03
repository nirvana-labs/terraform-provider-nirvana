// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type NKSClustersItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[NKSClustersItemsDataSourceModel] `json:"items,computed"`
}

type NKSClustersDataSourceModel struct {
	ProjectID         types.String                                                  `tfsdk:"project_id" query:"project_id,required"`
	Autoscaling       types.Bool                                                    `tfsdk:"autoscaling" query:"autoscaling,optional"`
	KubernetesVersion types.String                                                  `tfsdk:"kubernetes_version" query:"kubernetes_version,optional"`
	Name              types.String                                                  `tfsdk:"name" query:"name,optional"`
	Region            types.String                                                  `tfsdk:"region" query:"region,optional"`
	Status            types.String                                                  `tfsdk:"status" query:"status,optional"`
	VPCID             types.String                                                  `tfsdk:"vpc_id" query:"vpc_id,optional"`
	Tags              *[]types.String                                               `tfsdk:"tags" query:"tags,optional"`
	Sort              types.String                                                  `tfsdk:"sort" query:"sort,computed_optional"`
	MaxItems          types.Int64                                                   `tfsdk:"max_items"`
	Items             customfield.NestedObjectList[NKSClustersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NKSClustersDataSourceModel) toListParams(_ context.Context) (params nks.ClusterListParams, diags diag.Diagnostics) {
	mTags := []string{}
	if m.Tags != nil {
		for _, item := range *m.Tags {
			mTags = append(mTags, item.ValueString())
		}
	}

	params = nks.ClusterListParams{
		ProjectID: m.ProjectID.ValueString(),
		Tags:      mTags,
	}

	if !m.Autoscaling.IsNull() {
		params.Autoscaling = param.NewOpt(m.Autoscaling.ValueBool())
	}
	if !m.KubernetesVersion.IsNull() {
		params.KubernetesVersion = param.NewOpt(m.KubernetesVersion.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = param.NewOpt(m.Name.ValueString())
	}
	if !m.Region.IsNull() {
		params.Region = param.NewOpt(m.Region.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = param.NewOpt(m.Sort.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = nks.ClusterListParamsStatus(m.Status.ValueString())
	}
	if !m.VPCID.IsNull() {
		params.VPCID = param.NewOpt(m.VPCID.ValueString())
	}

	return
}

type NKSClustersItemsDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Autoscaling       types.Bool                     `tfsdk:"autoscaling" json:"autoscaling,computed"`
	CreatedAt         timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	KubernetesVersion types.String                   `tfsdk:"kubernetes_version" json:"kubernetes_version,computed"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	PoolIDs           customfield.List[types.String] `tfsdk:"pool_ids" json:"pool_ids,computed"`
	PrivateIP         types.String                   `tfsdk:"private_ip" json:"private_ip,computed"`
	ProjectID         types.String                   `tfsdk:"project_id" json:"project_id,computed"`
	PublicIP          types.String                   `tfsdk:"public_ip" json:"public_ip,computed"`
	Region            types.String                   `tfsdk:"region" json:"region,computed"`
	Status            types.String                   `tfsdk:"status" json:"status,computed"`
	Tags              customfield.List[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt         timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID             types.String                   `tfsdk:"vpc_id" json:"vpc_id,computed"`
}
