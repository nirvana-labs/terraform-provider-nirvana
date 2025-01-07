// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc_operation

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VPCOperationDataSourceModel struct {
	OperationID types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Kind        types.String `tfsdk:"kind" json:"kind,computed"`
	ResourceID  types.String `tfsdk:"resource_id" json:"resource_id,computed"`
	Status      types.String `tfsdk:"status" json:"status,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
}
