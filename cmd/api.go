package cmd

import (
	"time"

	"gitlab.sicepat.tech/platform/golib/log"
	"gitlab.sicepat.tech/platform/golib/tracing"

	rest "github.com/arymaulanamalik/sicepat_sample/transport/api"
	"github.com/spf13/cobra"
)

type Trace struct {
	TraceConfig TraceConfig `json:"tracer_config"`
}

type TraceConfig struct {
	Tracer     string `json:"tracer"`
	JeagerURL  string `json:"jeager_url"`
	JeagerMode string `json:"jeager_mode"`
}

var serverCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		initModule()
	},

	Run: func(cmd *cobra.Command, args []string) {
		_, err := tracing.GetOtelProvider("jeager", "authorization-svc", Trace{
			TraceConfig: TraceConfig{
				Tracer: "jeager", JeagerURL: "http://localhost:14268/api/traces", JeagerMode: "collector",
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		rest := rest.NewRest(&rest.Options{
			Port:            ":3001",
			GracefulTimeout: time.Duration(10) * time.Second,
			ReadTimeout:     time.Duration(10) * time.Second,
			WriteTimeout:    time.Duration(10) * time.Second,
			APITimeout:      10,
			MongoDB:         mdb,
		})
		go rest.Serve()
		rest.SignalCheck()
	},
}
