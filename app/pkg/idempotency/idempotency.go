package idempotency

import (
	"net/http"

	"github.com/google/uuid"

	"sumup-notifier/app/notifier/providers"
)

type RequestHeader map[string][]string

type Logger interface {
	Infof(format string, args ...interface{})
}

type Idempotency struct {
	ID        uuid.UUID
	Processed bool

	logger Logger
	driver providers.DBDriver
}

func NewIdempotency(logger Logger, db providers.DBDriver) *Idempotency {
	return &Idempotency{logger: logger, driver: db}
}

func (i *Idempotency) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Header["Idempotency-Key"]
		if !ok {
			http.Error(w, "idempotency key is missing", http.StatusBadRequest)
			return
		}

		payload := &Idempotency{}

		i.driver.First(payload, "id = ?", val)
		if payload.ID != uuid.Nil {
			i.logger.Infof("request with an ID: %v is already being processed", payload.ID)
			http.Error(w, "message is already being processed", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (i *Idempotency) Process(key string) error {
	keyUUID, errUUID := uuid.Parse(key)
	if errUUID != nil {
		return errUUID
	}

	idemp := Idempotency{ID: keyUUID, Processed: true}

	result := i.driver.Create(idemp)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
