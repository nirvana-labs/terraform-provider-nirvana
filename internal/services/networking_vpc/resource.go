// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/lib"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/importpath"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*NetworkingVPCResource)(nil)
var _ resource.ResourceWithModifyPlan = (*NetworkingVPCResource)(nil)
var _ resource.ResourceWithImportState = (*NetworkingVPCResource)(nil)

func NewResource() resource.Resource {
	return &NetworkingVPCResource{}
}

// NetworkingVPCResource defines the resource implementation.
type NetworkingVPCResource struct {
	client *nirvana.Client
}

func (r *NetworkingVPCResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_vpc"
}

func (r *NetworkingVPCResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NetworkingVPCResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *NetworkingVPCModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	operation, err := r.client.Networking.VPCs.New(
		ctx,
		networking.VPCNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := lib.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Networking.VPCs.Get(
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

func (r *NetworkingVPCResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *NetworkingVPCModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *NetworkingVPCModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	operation, err := r.client.Networking.VPCs.Update(
		ctx,
		data.ID.ValueString(),
		networking.VPCUpdateParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := lib.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Networking.VPCs.Get(
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

func (r *NetworkingVPCResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *NetworkingVPCModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.Networking.VPCs.Get(
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

	// Get subnet name from nested response
	subnet, diags := data.Subnet.Value(ctx)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || subnet == nil || subnet.Name.IsNull() || subnet.Name.ValueString() == "" {
		resp.Diagnostics.AddError("invalid subnet", "subnet name is required but was not provided")
		return
	}
	data.SubnetName = subnet.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkingVPCResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *NetworkingVPCModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := r.client.Networking.VPCs.Delete(
		ctx,
		data.ID.ValueString(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if errWaitForOperation := lib.Wait(ctx, r.client, operation.ID); errWaitForOperation != nil {
		resp.Diagnostics.AddError("failed to wait for operation", errWaitForOperation.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkingVPCResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *NetworkingVPCModel = new(NetworkingVPCModel)

	path := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<vpc_id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(path)

	res := new(http.Response)
	_, err := r.client.Networking.VPCs.Get(
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkingVPCResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
