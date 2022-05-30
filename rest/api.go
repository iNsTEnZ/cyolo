package rest

import (
	"context"
	"cyolo-exercise/configuration"
	"cyolo-exercise/output"
	"cyolo-exercise/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type API struct {
	cfg         configuration.RestConfiguration
	Router      *httprouter.Router
	server      *http.Server
	service     service.Processor
	printer     output.Printer
	rateLimiter chan struct{}
}

func New(cfg configuration.RestConfiguration, srv service.Processor, printer output.Printer) *API {
	var api *API

	if cfg.Enabled {
		api = &API{
			cfg:         cfg,
			service:     srv,
			printer:     printer,
			rateLimiter: make(chan struct{}, cfg.MaxConcurrentRequests),
		}
	}

	return api
}

func (api *API) Start() {
	api.Initialize()
	api.startServe()
}

func (api *API) Initialize() {
	api.Router = httprouter.New()

	logMiddleware := []func(next httprouter.Handle, name string) httprouter.Handle{
		api.RequestLogger,
	}

	api.routeWithMiddleware("POST", "/api/v1/process", api.process, logMiddleware...)
	api.routeWithMiddleware("GET", "/api/v1/histogram", api.histogram, logMiddleware...)
	api.Router.GET("/health", api.Health)

	api.server = &http.Server{
		Addr:         api.cfg.Port,
		Handler:      api.Router,
		ReadTimeout:  api.cfg.ServerReadTimeoutSec,
		WriteTimeout: api.cfg.ServerWriteTimeoutSec,
		IdleTimeout:  api.cfg.ServerIdleTimeoutSec,
	}
}

func (api *API) routeWithMiddleware(method, path string, handler httprouter.Handle, mws ...func(next httprouter.Handle, name string) httprouter.Handle) {
	for _, mw := range mws {
		handler = mw(handler, path)
	}

	api.Router.Handle(method, path, handler)
}

func (api *API) startServe() {
	log.Printf("listening on port %s", api.cfg.Port)

	if err := api.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("can't launch the server on port %s: %s", api.cfg.Port, err.Error())
	}
}

func (api *API) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), api.cfg.GracefulShutdownSec)
	defer cancel()

	api.server.SetKeepAlivesEnabled(false)
	err := api.server.Shutdown(ctx)

	if err != nil {
		log.Printf("api shutdown error: %s", err.Error())
	}
}
