package router

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
)

func Module() fx.Option {
	return fx.Provide(fx.Annotate(NewRouter, fx.As(new(bootstrap.Router))))
}
