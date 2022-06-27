package providers

import (
	"io"

	"gorm.io/gorm"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

type EmailConfiger interface {
	Decoder(string) ([]byte, error)
	MailConfig() string
}

type SlackConfiger interface {
	Decoder(string) ([]byte, error)
	SlackConfig() string
}

type SMSConfiger interface {
	Decoder(string) ([]byte, error)
	SMSConfig() string
}

type DBConfiger interface {
	Decoder(string) ([]byte, error)
	DBConfig() string
}

type BaseProvider interface {
	Send(from, to, message string) error
	Init(config interface{}) (interface{}, error)
	Name() string
	Process(options interface{}) error
}

type EmailProvider interface {
	BaseProvider
}

type SlackProvider interface {
	BaseProvider
}

type SMSProvider interface {
	BaseProvider
}

type RequestHelper interface {
	Validate(body io.ReadCloser, request BaseRequester) error
}

type BaseRequester interface {
	Validate() error
	Server() string
	Options() interface{}
}

type EmailRequester interface {
	BaseRequester
}

type SlackRequester interface {
	BaseRequester
}

type SMSRequester interface {
	BaseRequester
}

type BaseRegister interface {
	Get(key string) (BaseProvider, error)
}

type EmailRegister interface {
	BaseRegister
}

type SlackRegister interface {
	BaseRegister
}

type SMSRegister interface {
	BaseRegister
}

type DBCreate interface {
	Create(value interface{}) *gorm.DB
}

type DBGet interface {
	First(dest interface{}, conds ...interface{}) *gorm.DB
}

type DBDriver interface {
	DBCreate
	DBGet
}

type Validator interface {
	Struct(s interface{}) error
}

type Finalizer interface {
	Process(key string) error
}
