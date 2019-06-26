package main

import (
	"os"

	"github.com/ligato/cn-infra/agent"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logmanager"
	log "github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/cn-infra/rpc/grpc"
	"github.com/ligato/cn-infra/servicelabel"
)

// Tailor is a struct holding internal data for the tailored Agent
type Tailor struct {
	GRPC         grpc.Server
	KVStore      keyval.KvProtoPlugin
	LogManager   *logmanager.Plugin
	ServiceLabel *servicelabel.Plugin
}

// New creates new tailor instance.
func New() *Tailor {
	return &Tailor{
		GRPC:         &grpc.DefaultPlugin,
		KVStore:      &etcd.DefaultPlugin,
		LogManager:   &logmanager.DefaultPlugin,
		ServiceLabel: &servicelabel.DefaultPlugin,
	}
}

// Init initializes main plugin.
func (pr *Tailor) Init() error {
	return nil
}

// AfterInit executes operations after the init function.
func (pr *Tailor) AfterInit() error {
	resync.DefaultPlugin.DoResync()
	return nil
}

// Close can be used to close used resources.
func (pr *Tailor) Close() error {
	return nil
}

// String returns name of the plugin.
func (pr *Tailor) String() string {
	return "tailor"
}

func main() {
	Tailor := New()

	a := agent.NewAgent(agent.AllPlugins(Tailor))

	if err := a.Run(); err != nil {
		log.DefaultLogger().Fatal(err)
	}
}

func init() {
	log.DefaultLogger().SetOutput(os.Stdout)
	log.DefaultLogger().SetLevel(logging.DebugLevel)
}
