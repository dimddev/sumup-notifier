package slack

import (
	"sumup-notifier/app/notifier/providers"
)

type Options struct {
	Channel string
	Message string
}

type Request struct {
	Channel  string `json:"channel" validate:"required,alpha"`
	Message  string `json:"message" validate:"required"`
	Provider string `json:"provider" validate:"required,alpha"`

	validate providers.Validator
}

var _ providers.SlackRequester = (*Request)(nil)

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
	return &Options{Channel: r.Channel, Message: r.Message}
}
