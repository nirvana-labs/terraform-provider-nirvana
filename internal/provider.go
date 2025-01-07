// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/stainless-sdks/nirvana-terraform/internal/services/firewall_rule"
	"github.com/stainless-sdks/nirvana-terraform/internal/services/vm"
	"github.com/stainless-sdks/nirvana-terraform/internal/services/vm_operation"
	"github.com/stainless-sdks/nirvana-terraform/internal/services/vpc"
	"github.com/stainless-sdks/nirvana-terraform/internal/services/vpc_operation"
)

var _ provider.ProviderWithConfigValidators = (*NirvanaProvider)(nil)

// NirvanaProvider defines the provider implementation.
type NirvanaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NirvanaProviderModel describes the provider data model.
type NirvanaProviderModel struct {
	BaseURL   types.String `tfsdk:"base_url" json:"base_url,optional"`
	AuthToken types.String `tfsdk:"auth_token" json:"auth_token,required"`
}

func (p *NirvanaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nirvana"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"auth_token": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *NirvanaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *NirvanaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data NirvanaProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.AuthToken.IsNull() {
		opts = append(opts, option.WithAuthToken(data.AuthToken.ValueString()))
	}

	client := nirvana.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *NirvanaProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *NirvanaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		vm.NewResource,
		vpc.NewResource,
		firewall_rule.NewResource,
	}
}

func (p *NirvanaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		vm.NewVMDataSource,
		vm_operation.NewVMOperationDataSource,
		vpc.NewVPCDataSource,
		vpc_operation.NewVPCOperationDataSource,
		firewall_rule.NewFirewallRuleDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NirvanaProvider{
			version: version,
		}
	}
}
