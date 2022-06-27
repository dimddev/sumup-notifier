package validate

import (
	"go.uber.org/fx"

	"sumup-notifier/app/notifier/providers"
)

func Module() fx.Option {
	return fx.Provide(fx.Annotate(NewValidate, fx.As(new(providers.Validator))))
}
