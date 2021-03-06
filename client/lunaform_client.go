// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/getlunaform/lunaform/client/modules"
	"github.com/getlunaform/lunaform/client/providers"
	"github.com/getlunaform/lunaform/client/resources"
	"github.com/getlunaform/lunaform/client/stacks"
	"github.com/getlunaform/lunaform/client/state_backends"
	"github.com/getlunaform/lunaform/client/workspaces"
)

// Default lunaform HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new lunaform HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Lunaform {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new lunaform HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Lunaform {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new lunaform client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Lunaform {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Lunaform)
	cli.Transport = transport

	cli.Modules = modules.New(transport, formats)

	cli.Providers = providers.New(transport, formats)

	cli.Resources = resources.New(transport, formats)

	cli.Stacks = stacks.New(transport, formats)

	cli.StateBackends = state_backends.New(transport, formats)

	cli.Workspaces = workspaces.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Lunaform is a client for lunaform
type Lunaform struct {
	Modules *modules.Client

	Providers *providers.Client

	Resources *resources.Client

	Stacks *stacks.Client

	StateBackends *state_backends.Client

	Workspaces *workspaces.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Lunaform) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Modules.SetTransport(transport)

	c.Providers.SetTransport(transport)

	c.Resources.SetTransport(transport)

	c.Stacks.SetTransport(transport)

	c.StateBackends.SetTransport(transport)

	c.Workspaces.SetTransport(transport)

}
