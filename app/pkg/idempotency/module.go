package idempotency

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewIdempotency, fx.As(new(bootstrap.Idempoter))),
		fx.Annotate(NewIdempotency, fx.As(new(providers.Finalizer))),
	)
}
