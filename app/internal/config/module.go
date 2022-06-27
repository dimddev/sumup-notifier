package config

import (
	"go.uber.org/fx"

	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(
		fx.Annotate(
			NewConfig,
			fx.As(new(providers.SlackConfiger)),
			fx.As(new(providers.EmailConfiger)),
			fx.As(new(providers.SMSConfiger)),
			fx.As(new(providers.DBConfiger)),
		),
	)
}
