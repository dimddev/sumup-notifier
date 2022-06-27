package sms

import (
	"sumup-notifier/app/notifier/providers"
)

type Options struct {
	From    string
	To      string
	Message string
}

type Request struct {
	From     string `json:"from" validate:"required,ascii"`
	To       string `json:"to" validate:"required,ascii"`
	Message  string `json:"message" validate:"required,alpha"`
	Provider string `json:"provider" validate:"required,alpha"`

	validate providers.Validator
}

var _ providers.SMSRequester = (*Request)(nil)

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
