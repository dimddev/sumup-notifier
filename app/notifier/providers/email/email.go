package email

import (
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type Email struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	SSL       bool   `json:"ssl"`
	Signature string `json:"name"`

	logging providers.Logger
}

var _ providers.EmailProvider = (*Email)(nil)

func NewEmail(logging providers.Logger) *Email {
	return &Email{logging: logging}
}

func (e *Email) Send(from, to, message string) error {
	e.logging.Infof("Sending an email with content <%s> from: <%s> to: <%s>", message, from, to)
	return nil
}

func (e *Email) Name() string {
	return e.Signature
}

func (e *Email) Init(mailConf interface{}) (interface{}, error) {
	mail, ok := mailConf.(config.MailConfig)
	if !ok {
		return nil, errors.New("mail config doesn't exists")
	}

	mailCfg := &Email{
		Server:    mail.Server,
		Port:      mail.Port,
		SSL:       mail.SSL,
		Signature: mail.Name,
		logging:   e.logging,
	}

	if mailCfg.Server == "" && mailCfg.Port == 0 {
		return nil, errors.New("invalid configuration")
	}

	return mailCfg, nil
}

func (e *Email) Connect() error {
	e.logging.Infof("Connect to mail server %s:%d with SSL: %v", e.Server, e.Port, e.SSL)

	return nil
}

func (e *Email) Process(options interface{}) error {
	opt, ok := options.(*Options)
	if !ok {
		return errors.New("options does not exists")
	}

	if errConnect := e.Connect(); errConnect != nil {
		return errConnect
	}

	return e.Send(opt.From, opt.To, opt.Message)
}
