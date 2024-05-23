// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &PortainerProvider{}
var _ provider.ProviderWithFunctions = &PortainerProvider{}

type PortainerProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PortainerProviderModel describes the provider data model.
type PortainerProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ApiKey   types.String `tfsdk:"api_key"`
}

func (p *PortainerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "portainer"
	resp.Version = p.version
}

func (p *PortainerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "URL of the portainer instance",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API key",
				Required:            true,
			},
		},
	}
}

func (p *PortainerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *PortainerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *PortainerProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *PortainerProvider) Functions(ctx context.Context) []func() function.Function {
	return nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PortainerProvider{
			version: version,
		}
	}
}
