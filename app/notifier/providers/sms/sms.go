package sms

import (
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type SMS struct {
	Server    string `json:"server"`
	APIKey    string `json:"apikey"`
	Signature string `json:"name"`

	logging providers.Logger
}

var _ providers.SMSProvider = (*SMS)(nil)

func NewSMS(logging providers.Logger) *SMS {
	return &SMS{logging: logging}
}

func (e *SMS) Send(from, to, message string) error {
	e.logging.Infof("Sending a SMS with content: <%s> from: <%s> to: <%s> were send succesful", message, from, to)
	return nil
}

func (e *SMS) Name() string {
	return e.Signature
}

func (e *SMS) Init(smsConf interface{}) (interface{}, error) {
	sms, ok := smsConf.(config.SMSConfig)
	if !ok {
		return nil, errors.New("sms config doesn't exist")
	}

	smsCfg := &SMS{
		Server:    sms.Server,
		APIKey:    sms.APIKey,
		Signature: sms.Name,
		logging:   e.logging,
	}

	if smsCfg.Server == "" && smsCfg.APIKey == "" {
		return nil, errors.New("invalid configuration")
	}

	return smsCfg, nil
}

func (e *SMS) Connect() error {
	e.logging.Infof("Connect to SMS server %s", e.Server)

	return nil
}

func (e *SMS) Process(options interface{}) error {
	opt, ok := options.(*Options)
	if !ok {
		return errors.New("options doesn't exists")
	}

	if errConnect := e.Connect(); errConnect != nil {
		return errConnect
	}

	return e.Send(opt.From, opt.To, opt.Message)
}
