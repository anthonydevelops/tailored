//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model ./model/server.proto

package grpc

import (
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/logging"
)

// RegisterFlags registers command line flags.
func RegisterFlags() {
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
