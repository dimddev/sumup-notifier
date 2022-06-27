package main

import (
	"go.uber.org/fx"

	"sumup-notifier/app/internal/bootstrap"
	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/internal/request"
	"sumup-notifier/app/internal/router"
	"sumup-notifier/app/notifier/providers/email"
	"sumup-notifier/app/notifier/providers/slack"
	"sumup-notifier/app/notifier/providers/sms"
	"sumup-notifier/app/pkg/db"
	"sumup-notifier/app/pkg/idempotency"
	"sumup-notifier/app/pkg/logger"
	"sumup-notifier/app/pkg/validate"
)

func main() {
	app := fx.New(
		idempotency.Module(),
		logger.Module(),
		config.Module(),
		validate.Module(),
		db.Module(),
		request.Module(),
		router.Module(),
		email.Module(),
		slack.Module(),
		sms.Module(),
		fx.Invoke(bootstrap.Bootstrap),
	)

	app.Run()
}
