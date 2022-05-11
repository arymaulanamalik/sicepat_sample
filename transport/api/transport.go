package api

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/prometheus/client_golang/prometheus/promhttp" // http-swagger middleware
	//_ "github.com/arymaulanamalik/sicepat_sample/docs"
	"github.com/arymaulanamalik/sicepat_sample/pkg/database"
	"github.com/arymaulanamalik/sicepat_sample/pkg/logger"
	"github.com/arymaulanamalik/sicepat_sample/registry"
	"github.com/arymaulanamalik/sicepat_sample/transport/api/group"
	"github.com/arymaulanamalik/sicepat_sample/transport/api/handler"
	grace "gitlab.sicepat.tech/platform/golib/httputil"
	myrouter "gitlab.sicepat.tech/platform/golib/router"
)

type Options struct {
	Port            string
	GracefulTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	APITimeout      int
	MongoDB         *database.MongoDatabase
}

type Handler struct {
	options     *Options
	listenErrCh chan error
}

// NewRest
// @title Authorization Service API
// @version 1.0
// @description Authorization Service API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRest(o *Options) *Handler {

	reg := registry.NewRegistry(
		registry.NewMongoConn(o.MongoDB.DB),
	)

	appController := reg.NewAppController()
	middleware := reg.Middleware()

	route := myrouter.New(&myrouter.Options{Timeout: o.APITimeout})

	v2 := myrouter.New(&myrouter.Options{
		Prefix:  "/v2",
		Timeout: o.APITimeout,
	})

	handlerImpl := handler.HandlerImpl{
		Controller: appController,
	}

	// global
	route.Httprouter.Handler("GET", "/metrics", promhttp.Handler())
	route.Httprouter.HandlerFunc("GET", "/swagger/*any", httpSwagger.WrapHandler)
	route.GET("/health", handler.Health)

	group.NewUsers(v2, handlerImpl, middleware)

	return &Handler{
		options: o,
	}
}

func (h *Handler) Serve() {
	log := logger.GetLogger("transport", "serve")
	log.Infof("API Listening on %s", h.options.Port)
	h.listenErrCh <- grace.Serve(h.options.Port, myrouter.WrapperHandler(), h.options.GracefulTimeout, h.options.ReadTimeout, h.options.WriteTimeout)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}

func (h *Handler) SignalCheck() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		logger.Log.Infoln("Exiting gracefully...")
	case err := <-h.ListenError():
		logger.Log.Errorln("Error starting web server, exiting gracefully:", err)
	}
}
