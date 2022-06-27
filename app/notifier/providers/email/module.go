package email

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewEmail, fx.As(new(providers.EmailProvider))),
		fx.Annotate(NewRegistry, fx.As(new(providers.EmailRegister))),
		fx.Annotate(NewRequest, fx.As(new(providers.EmailRequester))),
		fx.Annotate(NewHandler, fx.As(new(bootstrap.EmailHandler))),
	)
}
