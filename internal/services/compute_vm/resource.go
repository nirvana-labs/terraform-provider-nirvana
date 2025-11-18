// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/lib"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/importpath"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ComputeVMResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ComputeVMResource)(nil)
var _ resource.ResourceWithImportState = (*ComputeVMResource)(nil)

func NewResource() resource.Resource {
	return &ComputeVMResource{
		waiter: lib.NewOperationWaiter().WithLinearBackoff(1*time.Second, 500*time.Millisecond, 10*time.Second),
	}
}

// ComputeVMResource defines the resource implementation.
type ComputeVMResource struct {
	client *nirvana.Client
	waiter *lib.OperationWaiter
}

func (r *ComputeVMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compute_vm"
}

func (r *ComputeVMResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*nirvana.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *nirvana.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ComputeVMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ComputeVMModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	operation, err := r.client.Compute.VMs.New(
		ctx,
		compute.VMNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := r.waiter.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Compute.VMs.Get(
		ctx,
		operation.ResourceID,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeVMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ComputeVMModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ComputeVMModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	operation, err := r.client.Compute.VMs.Update(
		ctx,
		data.ID.ValueString(),
		compute.VMUpdateParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := r.waiter.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Compute.VMs.Get(
		ctx,
		operation.ResourceID,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeVMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ComputeVMModel
	var oldState *ComputeVMModel // Store original state

	// Get the current state
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &oldState)...) // Save copy of original state

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.Compute.VMs.Get(
		ctx,
		data.ID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Get all volumes for the VM using VMs.Volumes.List
	volumesRes := new(http.Response)
	_, err = r.client.Compute.VMs.Volumes.List(
		ctx,
		data.ID.ValueString(),
		compute.VMVolumeListParams{},
		option.WithResponseBodyInto(&volumesRes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get volumes", err.Error())
	} else {
		volumesBytes, _ := io.ReadAll(volumesRes.Body)

		var volumeList compute.VolumeList
		if jsonErr := apijson.Unmarshal(volumesBytes, &volumeList); jsonErr != nil {
			resp.Diagnostics.AddError("failed to deserialize volumes", jsonErr.Error())
		} else {
			var bootVolume *ComputeVMBootVolumeModel
			var bootVolumeID types.String

			// Only collect volumes that match the original configuration
			var configuredVolumeIDs []types.String
			var configuredDataVolumes []*ComputeVMDataVolumesModel

			for _, volume := range volumeList.Items {
				switch volume.Kind {
				case compute.VolumeKindBoot:
					bootVolumeID = types.StringValue(volume.ID)
					bootVolume = &ComputeVMBootVolumeModel{
						Size: types.Int64Value(volume.Size),
					}
				case compute.VolumeKindData:
					// Only include this volume if it matches something in the original configuration
					if oldState.DataVolumes != nil {
						for _, configuredVolume := range *oldState.DataVolumes {
							if volume.Name == configuredVolume.Name.ValueString() &&
								volume.Size == configuredVolume.Size.ValueInt64() {
								configuredVolumeIDs = append(configuredVolumeIDs, types.StringValue(volume.ID))
								configuredDataVolumes = append(configuredDataVolumes, &ComputeVMDataVolumesModel{
									Name: types.StringValue(volume.Name),
									Size: types.Int64Value(volume.Size),
								})
								break
							}
						}
					}
				default:
					continue
				}
			}

			if bootVolume != nil {
				data.BootVolume = bootVolume
				data.BootVolumeID = bootVolumeID
			} else {
				data.BootVolume = nil
				data.BootVolumeID = types.StringNull()
			}

			// Set data_volume_ids with only the IDs of volumes that match the configuration
			if len(configuredVolumeIDs) > 0 {
				attrValues := make([]attr.Value, len(configuredVolumeIDs))
				for i, id := range configuredVolumeIDs {
					attrValues[i] = id
				}
				data.DataVolumeIDs = customfield.NewListMust[types.String](ctx, attrValues)
			} else {
				data.DataVolumeIDs = customfield.NewListMust[types.String](ctx, []attr.Value{})
			}

			// Set data_volumes with only the volumes that match the configuration
			if oldState.DataVolumes != nil {
				// DataVolumes was configured in Terraform, so we should update it with matching volumes
				if len(configuredDataVolumes) > 0 {
					data.DataVolumes = &configuredDataVolumes
				} else {
					// Configuration had data_volumes but none are attached now
					// Keep the field as configured (empty slice) rather than nil
					emptyDataVolumes := []*ComputeVMDataVolumesModel{}
					data.DataVolumes = &emptyDataVolumes
				}
			} else {
				// DataVolumes was never configured in Terraform, so volumes are managed externally
				// Keep it nil to indicate this resource doesn't manage data volumes
				data.DataVolumes = nil
			}
		}
	}

	// Preserve values that aren't returned by the API
	data.OSImageName = oldState.OSImageName
	data.SSHKey.PublicKey = oldState.SSHKey.PublicKey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeVMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ComputeVMModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := r.client.Compute.VMs.Delete(
		ctx,
		data.ID.ValueString(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := r.waiter.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeVMResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ComputeVMModel)

	path := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<vm_id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(path)

	res := new(http.Response)
	_, err := r.client.Compute.VMs.Get(
		ctx,
		path,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Get all volumes for the VM during import to populate state correctly
	volumesRes := new(http.Response)
	_, err = r.client.Compute.VMs.Volumes.List(
		ctx,
		path,
		compute.VMVolumeListParams{},
		option.WithResponseBodyInto(&volumesRes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get volumes during import", err.Error())
	} else {
		volumesBytes, _ := io.ReadAll(volumesRes.Body)

		var volumeList compute.VolumeList
		if jsonErr := apijson.Unmarshal(volumesBytes, &volumeList); jsonErr != nil {
			resp.Diagnostics.AddError("failed to deserialize volumes during import", jsonErr.Error())
		} else {
			var bootVolume *ComputeVMBootVolumeModel
			var bootVolumeID types.String

			for _, volume := range volumeList.Items {
				switch volume.Kind {
				case compute.VolumeKindBoot:
					bootVolumeID = types.StringValue(volume.ID)
					bootVolume = &ComputeVMBootVolumeModel{
						Size: types.Int64Value(volume.Size),
					}
				case compute.VolumeKindData:
					// During import, we don't collect data volume IDs since we don't know
					// which ones were configured in Terraform vs added externally
					continue
				default:
					continue
				}
			}

			if bootVolume != nil {
				data.BootVolume = bootVolume
				data.BootVolumeID = bootVolumeID
			} else {
				data.BootVolume = nil
				data.BootVolumeID = types.StringNull()
			}

			// For import, don't populate data_volume_ids since we don't know which volumes
			// were configured in Terraform vs added externally. This will be populated
			// correctly on the first refresh after import when the user has their configuration.
			data.DataVolumeIDs = customfield.NewListMust[types.String](ctx, []attr.Value{})

			// For import, we don't populate data_volumes since we don't know if they
			// were originally configured in Terraform. The user will need to add them
			// to their configuration if they want to manage them through this resource.
			data.DataVolumes = nil
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeVMResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
