package bootstrap

// Bootstrap - a handler and migration initializer.
func Bootstrap(
	router Router,
	emailHandle EmailHandler,
	slackHandle SlackHandler,
	smsHandle SMSHandler,
	middleware Idempoter,
	db Migrator,
) error {
	emailRequest, errEmail := emailHandle.Handle()
	if errEmail != nil {
		return errEmail
	}

	slackRequest, errSlack := slackHandle.Handle()
	if errSlack != nil {
		return errSlack
	}

	smsRequest, errSMS := smsHandle.Handle()
	if errSMS != nil {
		return errSMS
	}

	router.Handle("/email-notify", middleware.Middleware(emailRequest)).Methods("POST")
	router.Handle("/slack-notify", middleware.Middleware(slackRequest)).Methods("POST")
	router.Handle("/sms-notify", middleware.Middleware(smsRequest)).Methods("POST")

	err := db.AutoMigrate(&middleware)
	if err != nil {
		return err
	}

	return nil
}
