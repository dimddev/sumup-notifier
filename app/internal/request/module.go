package request

import (
	"go.uber.org/fx"

	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewHelper, fx.As(new(providers.RequestHelper))),
	)
}
