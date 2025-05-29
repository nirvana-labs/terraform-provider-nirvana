// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/api_key"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_vm"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_volume"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/networking_firewall_rule"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/networking_vpc"
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
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey  types.String `tfsdk:"api_key" json:"api_key,optional"`
}

func (p *NirvanaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nirvana"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *NirvanaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *NirvanaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data NirvanaProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("NIRVANA_LABS_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("NIRVANA_LABS_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"NIRVANA_LABS_API_KEY\" environment variable.",
		)
		return
	}

	client := nirvana.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *NirvanaProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *NirvanaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		api_key.NewResource,
		compute_vm.NewResource,
		compute_volume.NewResource,
		networking_vpc.NewResource,
		networking_firewall_rule.NewResource,
	}
}

func (p *NirvanaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NirvanaProvider{
			version: version,
		}
	}
}
