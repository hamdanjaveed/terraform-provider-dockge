// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &DockgeProvider{}
var _ provider.ProviderWithFunctions = &DockgeProvider{}

type DockgeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type DockgeProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

// Metadata implements provider.Provider.
func (d *DockgeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dockge"
	resp.Version = d.version
}

// Schema implements provider.Provider.
func (d *DockgeProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

// Configure implements provider.Provider.
func (d *DockgeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data DockgeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources implements provider.Provider.
func (d *DockgeProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources implements provider.Provider.
func (d *DockgeProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewStackResource,
	}
}

// Functions implements provider.ProviderWithFunctions.
func (d *DockgeProvider) Functions(context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DockgeProvider{
			version: version,
		}
	}
}
