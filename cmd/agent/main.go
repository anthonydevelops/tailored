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

// tailor is a struct holding internal data for the tailored Agent
type tailor struct {
	GRPC         grpc.Server
	KVStore      keyval.KvProtoPlugin
	LogManager   *logmanager.Plugin
	ServiceLabel *servicelabel.Plugin
}

// New creates new tailor instance.
func New() *tailor {
	return &tailor{
		GRPC:         &grpc.DefaultPlugin,
		KVStore:      &etcd.DefaultPlugin,
		LogManager:   &logmanager.DefaultPlugin,
		ServiceLabel: &servicelabel.DefaultPlugin,
	}
}

// Init initializes main plugin.
func (pr *tailor) Init() error {
	return nil
}

func (pr *tailor) AfterInit() error {
	resync.DefaultPlugin.DoResync()
	return nil
}

// Close can be used to close used resources.
func (pr *tailor) Close() error {
	return nil
}

// String returns name of the plugin.
func (pr *tailor) String() string {
	return "tailor"
}

func main() {
	tailor := New()

	a := agent.NewAgent(agent.AllPlugins(tailor))

	if err := a.Run(); err != nil {
		log.DefaultLogger().Fatal(err)
	}
}

func init() {
	log.DefaultLogger().SetOutput(os.Stdout)
	log.DefaultLogger().SetLevel(logging.DebugLevel)
}
