// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm_os_image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type ComputeVMOSImagesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ComputeVMOSImagesItemsDataSourceModel] `json:"items,computed"`
}

type ComputeVMOSImagesDataSourceModel struct {
	MaxItems types.Int64                                                         `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ComputeVMOSImagesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ComputeVMOSImagesDataSourceModel) toListParams(_ context.Context) (params compute.VMOSImageListParams, diags diag.Diagnostics) {
	params = compute.VMOSImageListParams{}

	return
}

type ComputeVMOSImagesItemsDataSourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DisplayName types.String      `tfsdk:"display_name" json:"display_name,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
}
