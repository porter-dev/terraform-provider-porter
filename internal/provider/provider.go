package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure PorterProvider satisfies various provider interfaces.
var _ provider.Provider = &providerPorter{}

// providerPorter creates a TF provider for the Porter API
type providerPorter struct {
	// version is the Porter provider version
	version string
}

// NewProviderPorter is a helper function to simplify provider server and testing implementation.
func NewProviderPorter(version string) func() provider.Provider {
	return func() provider.Provider {
		return &providerPorter{
			version: version,
		}
	}
}

// providerPorterModel is the data for configuring the Porter provider
type providerPorterModel struct {
	Host     types.String `tfsdk:"host"`
	ApiToken types.String `tfsdk:"api_token"`
}

func (p *providerPorter) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "porter"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *providerPorter) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Porter.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for Porter API. May also be provided via PORTER_HOST environment variable.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Description: "API Token for Porter API. May also be provided via PORTER_TOKEN environment variable.",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *providerPorter) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Porter client")

	// Retrieve provider data from configuration
	var config providerPorterModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Porter API Host",
			"The provider cannot create the Porter API client as there is an unknown configuration value for the Porter API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PORTER_HOST environment variable.",
		)
	}

	if config.ApiToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown Porter API Token",
			"The provider cannot create the Porter API client as there is an unknown configuration value for the Porter API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PORTER_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("PORTER_HOST")
	apiToken := os.Getenv("PORTER_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.ApiToken.IsNull() {
		apiToken = config.ApiToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Porter API Host",
			"The provider cannot create the Porter API client as there is a missing or empty value for the Porter API host. ",
		)
	}

	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing Porter API Token",
			"Please set via PORTER_TOKEN environment variable or in the provider configuration.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "porter_host", host)
	ctx = tflog.SetField(ctx, "porter_token", apiToken)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "porter_token")

	tflog.Debug(ctx, "Creating Porter client")

	// TODO:  this part
	// Create a new Porter client using the configuration values
	// client, err := hashicups.NewClient(&host, &username, &password)
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Unable to Create Porter API Client",
	// 		"An unexpected error occurred when creating the Porter API client. "+
	// 			"If the error is not clear, please contact the provider developers.\n\n"+
	// 			"Porter Client Error: "+err.Error(),
	// 	)
	// 	return
	// }

	// Make the Porter client available during DataSource and Resource
	// type Configure methods.
	// resp.DataSourceData = client
	// resp.ResourceData = client

	tflog.Info(ctx, "Configured Porter client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *providerPorter) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// NewCoffeesDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *providerPorter) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewEksClusterResource,
		NewAksClusterResource,
	}
}
