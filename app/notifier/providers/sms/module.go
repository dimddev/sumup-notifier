package sms

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(NewSMS, fx.As(new(providers.SMSProvider))),
		fx.Annotate(NewRegistry, fx.As(new(providers.SMSRegister))),
		fx.Annotate(NewRequest, fx.As(new(providers.SMSRequester))),
		fx.Annotate(NewHandler, fx.As(new(bootstrap.SMSHandler))),
	)
}
