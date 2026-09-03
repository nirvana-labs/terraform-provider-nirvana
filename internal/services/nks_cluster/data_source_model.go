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

type NKSClusterDataSourceModel struct {
	ID                types.String                        `tfsdk:"id" path:"cluster_id,computed"`
	ClusterID         types.String                        `tfsdk:"cluster_id" path:"cluster_id,optional"`
	Autoscaling       types.Bool                          `tfsdk:"autoscaling" json:"autoscaling,computed"`
	CreatedAt         timetypes.RFC3339                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	KubernetesVersion types.String                        `tfsdk:"kubernetes_version" json:"kubernetes_version,computed"`
	Name              types.String                        `tfsdk:"name" json:"name,computed"`
	PrivateIP         types.String                        `tfsdk:"private_ip" json:"private_ip,computed"`
	ProjectID         types.String                        `tfsdk:"project_id" json:"project_id,computed"`
	PublicIP          types.String                        `tfsdk:"public_ip" json:"public_ip,computed"`
	Region            types.String                        `tfsdk:"region" json:"region,computed"`
	Status            types.String                        `tfsdk:"status" json:"status,computed"`
	UpdatedAt         timetypes.RFC3339                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	VPCID             types.String                        `tfsdk:"vpc_id" json:"vpc_id,computed"`
	PoolIDs           customfield.List[types.String]      `tfsdk:"pool_ids" json:"pool_ids,computed"`
	Tags              customfield.List[types.String]      `tfsdk:"tags" json:"tags,computed"`
	FindOneBy         *NKSClusterFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *NKSClusterDataSourceModel) toListParams(_ context.Context) (params nks.ClusterListParams, diags diag.Diagnostics) {
	mFindOneByTags := []string{}
	if m.FindOneBy.Tags != nil {
		for _, item := range *m.FindOneBy.Tags {
			mFindOneByTags = append(mFindOneByTags, item.ValueString())
		}
	}

	params = nks.ClusterListParams{
		ProjectID: m.FindOneBy.ProjectID.ValueString(),
		Tags:      mFindOneByTags,
	}

	if !m.FindOneBy.Autoscaling.IsNull() {
		params.Autoscaling = param.NewOpt(m.FindOneBy.Autoscaling.ValueBool())
	}
	if !m.FindOneBy.KubernetesVersion.IsNull() {
		params.KubernetesVersion = param.NewOpt(m.FindOneBy.KubernetesVersion.ValueString())
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = param.NewOpt(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Region.IsNull() {
		params.Region = param.NewOpt(m.FindOneBy.Region.ValueString())
	}
	if !m.FindOneBy.Sort.IsNull() {
		params.Sort = param.NewOpt(m.FindOneBy.Sort.ValueString())
	}
	if !m.FindOneBy.Status.IsNull() {
		params.Status = nks.ClusterListParamsStatus(m.FindOneBy.Status.ValueString())
	}
	if !m.FindOneBy.VPCID.IsNull() {
		params.VPCID = param.NewOpt(m.FindOneBy.VPCID.ValueString())
	}

	return
}

type NKSClusterFindOneByDataSourceModel struct {
	ProjectID         types.String    `tfsdk:"project_id" query:"project_id,required"`
	Autoscaling       types.Bool      `tfsdk:"autoscaling" query:"autoscaling,optional"`
	KubernetesVersion types.String    `tfsdk:"kubernetes_version" query:"kubernetes_version,optional"`
	Name              types.String    `tfsdk:"name" query:"name,optional"`
	Region            types.String    `tfsdk:"region" query:"region,optional"`
	Sort              types.String    `tfsdk:"sort" query:"sort,computed_optional"`
	Status            types.String    `tfsdk:"status" query:"status,optional"`
	Tags              *[]types.String `tfsdk:"tags" query:"tags,optional"`
	VPCID             types.String    `tfsdk:"vpc_id" query:"vpc_id,optional"`
}
