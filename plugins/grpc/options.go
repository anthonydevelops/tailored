package grpc

import (
	"log"

	"github.com/ligato/cn-infra/rpc/grpc"
	"github.com/ligato/cn-infra/rpc/rest"

	"github.com/ligato/cn-infra/logging"
)

// DefaultPlugin is default instance of Plugin.
var DefaultPlugin = *NewPlugin()

// NewPlugin creates a new Plugin with the provides Options
func NewPlugin(opts ...Option) *Plugin {
	p.PluginName = "GRPC"
	grpcConf := grpc.NewPlugin(
		grpc.UseHTTP(&rest.DefaultPlugin),
		grpc.UseConf(grpc.Config{
			Endpoint: "localhost:9191",
		}),
	)

	p := &Plugin{
		Log:  logging.PluginLogger,
		GRPC: grpcConf,
	}

	for _, o := range opts {
		o(p)
	}

	if p.Deps.Log == nil {
		log.Println(p.String())
		p.Deps.Log = logging.ForPlugin(p.String())
	}

	return p
}

// Option is a function that acts on a Plugin to inject Dependencies or configuration
type Option func(*Plugin)

// UseDeps returns Option that can inject custom dependencies.
func UseDeps(cb func(*Deps)) Option {
	return func(p *Plugin) {
		cb(&p.Deps)
	}
}
