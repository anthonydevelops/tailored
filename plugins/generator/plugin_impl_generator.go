package generator

import (
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/logging"
	// todo: add any necessary imports for your plugin
)

// RegisterFlags registers command line flags.
func RegisterFlags() {
	// todo: add command line flags here if needed
}

func init() {
	RegisterFlags()
}

// Plugin holds the internal data structures of the Rest Plugin
type Plugin struct {
	Deps
}

// Deps groups the dependencies of the Rest Plugin.
type Deps struct {
	infra.PluginDeps
	// todo: add any additional dependencies here
}

// Init initializes the Plugin
func (p *Plugin) Init() error {
	p.Log.SetLevel(logging.DebugLevel)
	return nil
}

// AfterInit can be used to run plugin functionality.
func (p *Plugin) AfterInit() (err error) {
	return nil
}

// Close is NOOP.
func (p *Plugin) Close() error {
	return nil
}
