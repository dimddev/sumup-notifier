package bootstrap

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Handle(path string, handler http.Handler) *mux.Route
}

type BaseHandler interface {
	Handle() (http.Handler, error)
}

type EmailHandler interface {
	BaseHandler
}

type SlackHandler interface {
	BaseHandler
}

type SMSHandler interface {
	BaseHandler
}

type Idempoter interface {
	Middleware(next http.Handler) http.Handler
}

type Migrator interface {
	AutoMigrate(dst ...interface{}) error
}
