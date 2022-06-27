package request

import (
	"encoding/json"
	"io"

	"sumup-notifier/app/notifier/providers"
)

type Helper struct {
	logger providers.Logger
}

func NewHelper(logger providers.Logger) *Helper {
	return &Helper{logger: logger}
}

func (h *Helper) Validate(body io.ReadCloser, request providers.BaseRequester) error {
	if errDecode := h.decode(body, request); errDecode != nil {
		return errDecode
	}

	errValidate := request.Validate()
	if errValidate != nil {
		h.logger.Info(errValidate.Error())
		return errValidate
	}

	return nil
}

func (h *Helper) decode(body io.ReadCloser, request providers.BaseRequester) error {
	decoder := json.NewDecoder(body)

	errDecode := decoder.Decode(request)
	if errDecode != nil {
		return errDecode
	}

	return nil
}
