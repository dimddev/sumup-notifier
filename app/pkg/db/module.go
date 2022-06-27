package db

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		NewDriver,
		fx.Annotate(NewDriver, fx.As(new(providers.DBCreate))),
		fx.Annotate(NewDriver, fx.As(new(providers.DBGet))),
		fx.Annotate(NewDriver, fx.As(new(providers.DBDriver))),
		fx.Annotate(NewDriver, fx.As(new(bootstrap.Migrator))),
	)
}
