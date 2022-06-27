package router

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

type Logger interface {
	Info(args ...interface{})
	Errorf(format string, args ...interface{})
}

func NewRouter(lc fx.Lifecycle, logger Logger) *mux.Router {
	// First, we construct the router and server. We don't want to start the server
	// until all handlers are registered.
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server.")

			g, _ := errgroup.WithContext(ctx)
			go g.Go(func() error {
				return server.ListenAndServe()
			})

			if err := g.Wait(); err != nil {
				logger.Errorf("fatal error: %v", err.Error())
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return router
}
