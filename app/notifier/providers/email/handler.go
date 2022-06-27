package email

import (
	"net/http"

	"sumup-notifier/app/notifier/providers"
)

type Handler struct {
	logger   providers.Logger
	request  providers.EmailRequester
	registry providers.EmailRegister
	db       providers.Finalizer
	helper   providers.RequestHelper
}

func NewHandler(
	logger providers.Logger,
	request providers.EmailRequester,
	registry providers.EmailRegister,
	db providers.Finalizer,
	helper providers.RequestHelper,
) *Handler {
	return &Handler{logger: logger, request: request, registry: registry, db: db, helper: helper}
}

func (h *Handler) Handle() (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if errDecode := h.helper.Validate(r.Body, h.request); errDecode != nil {
			http.Error(w, errDecode.Error(), http.StatusBadRequest)
		}

		provider, errProvider := h.registry.Get(h.request.Server())
		if errProvider != nil {
			http.Error(w, errProvider.Error(), http.StatusBadRequest)
			return
		}

		if errProcess := provider.Process(h.request.Options()); errProcess != nil {
			http.Error(w, errProcess.Error(), http.StatusBadRequest)
			return
		}

		val := r.Header.Get("Idempotency-Key")
		if val == "" {
			http.Error(w, "idempotency key is missing", http.StatusBadRequest)
			return
		}

		if errStore := h.db.Process(val); errStore != nil {
			http.Error(w, errStore.Error(), http.StatusBadRequest)
		}
	}), nil
}
