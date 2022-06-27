package email

import (
	"sumup-notifier/app/notifier/providers"
)

type Options struct {
	From    string
	To      string
	Message string
}

type Request struct {
	From     string `json:"from" validate:"required,email"`
	To       string `json:"to" validate:"required,email"`
	Message  string `json:"message" validate:"required"`
	Provider string `json:"provider" validate:"required,alpha"`

	validate providers.Validator
}

var _ providers.EmailRequester = (*Request)(nil)

func NewRequest(validate providers.Validator) *Request {
	request := new(Request)
	request.validate = validate
	return request
}

func (r *Request) Validate() error {
	err := r.validate.Struct(r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) Server() string {
	return r.Provider
}

func (r *Request) Options() interface{} {
	return &Options{From: r.From, To: r.To, Message: r.Message}
}
