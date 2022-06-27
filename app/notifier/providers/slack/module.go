package slack

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewSlack, fx.As(new(providers.SlackProvider))),
		fx.Annotate(NewRegistry, fx.As(new(providers.SlackRegister))),
		fx.Annotate(NewRequest, fx.As(new(providers.SlackRequester))),
		fx.Annotate(NewHandler, fx.As(new(bootstrap.SlackHandler))),
	)
}
