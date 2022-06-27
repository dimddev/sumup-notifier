package logger

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/router"
	"sumup-notifier/app/notifier/providers"
	"sumup-notifier/app/pkg/idempotency"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewLogRus, fx.As(new(Logger))),
		fx.Annotate(
			NewLogging,
			fx.As(new(router.Logger)),
			fx.As(new(providers.Logger)),
			fx.As(new(idempotency.Logger)),
		),
	)
}
