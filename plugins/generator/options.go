package generator

import (
	"log"

	"github.com/ligato/cn-infra/logging"
)

// DefaultPlugin is default instance of Plugin.
var DefaultPlugin = *NewPlugin()

// NewPlugin creates a new Plugin with the provides Options
func NewPlugin(opts ...Option) *Plugin {
	p := &Plugin{}

	p.PluginName = "GENERATOR"
	// todo: initialize any other pluign Deps here, if applicable

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
